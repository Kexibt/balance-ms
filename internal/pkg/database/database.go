package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Database отвечает за взаимодействие с бд
type Database struct {
	conn *pgxpool.Pool
	oni  chan map[string]interface{}
}

// NewDatabase конструктор Database
func NewDatabase(cfg Config) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetConnectionTimeout())
	defer cancel()

	connection, err := connect(ctx, cfg.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		conn: connection,
		oni:  make(chan map[string]interface{}, 100),
	}
}

func connect(ctx context.Context, connectionStr string) (conn *pgxpool.Pool, err error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err = pgxpool.Connect(context.Background(), connectionStr)
			if err == nil {
				return
			}
		}
	}
}

// Close закрывает пул подключений
func (d *Database) Close() error {
	close(d.oni)
	d.conn.Close()
	return nil
}

// ListenAndServe слушает и прислуживается
func (d *Database) ListenAndServe() {
	for updates := range d.oni {
		go insert(d.conn, updates, d.oni) // можно будет сделать батчами, если обработок много,
		// но так как не задано количество запросов, то я сделал без батчей
	}
}

func insert(conn *pgxpool.Pool, updates map[string]interface{}, oni chan map[string]interface{}) {
	tr, err := conn.Begin(context.Background())

	if err != nil {
		oni <- updates
		log.Println(err)
		if tr != nil {
			tr.Rollback(context.Background())
		}
		return
	}

	_, err = tr.Exec(context.Background(), fmt.Sprintf(
		`INSERT INTO public.balances ("userid", "balance") VALUES ('%s', %v);`,
		updates["userID"].(string), updates["balance"].(float64),
	),
	)
	if err != nil {
		tr.Rollback(context.Background()) // todo: написать триггер
		// update todo: в силу внешних обстоятельств многое не успеваю поэтому костыль
		// P.S. на костылях даже быстрее(не всегда, перед применением проконсультируйтесь с экспертом)
		tr, err = conn.Begin(context.Background())
		if err != nil {
			oni <- updates
			log.Println(err)
			if tr != nil {
				tr.Rollback(context.Background())
			}
			return
		}

		_, err = tr.Exec(context.Background(), fmt.Sprintf(
			`UPDATE balances SET balance = %v WHERE userid = '%s';`,
			updates["balance"].(float64), updates["userID"].(string),
		),
		)
		if err != nil {
			oni <- updates
			log.Println(err)
			tr.Rollback(context.Background())
			return
		}
	}

	err = tr.Commit(context.Background())
	if err != nil {
		oni <- updates
		log.Println(err)
		tr.Rollback(context.Background())
	}
}

// Add добавляет запрос на update в очередь запросов
func (d *Database) Add(update map[string]interface{}) {
	d.oni <- update
}
