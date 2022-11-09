package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Database отвечает за взаимодействие с бд
type DatabaseCRUD struct {
	conn *pgxpool.Pool
}

// NewDatabase конструктор Database
func NewDatabaseCRUD(cfg Config) *DatabaseCRUD {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetConnectionTimeout())
	defer cancel()

	connection, err := connect(ctx, cfg.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return &DatabaseCRUD{
		conn: connection,
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
func (d *DatabaseCRUD) Close() error {
	d.conn.Close()
	return nil
}

// GetBalance для просмотра баланса
func (d *DatabaseCRUD) GetBalance(id string) (float64, error) {
	tr, err := d.conn.Begin(context.Background())
	if err != nil {
		return 0, err
	}

	var amount float64
	row := tr.QueryRow(context.Background(),
		`SELECT balance FROM public.balances
			WHERE userid = $1;`, id,
	)
	row.Scan(&amount)

	err = tr.Commit(context.TODO())
	return amount, err
}

// ChangeBalance для изменения баланса
func (d *DatabaseCRUD) ChangeBalance(id string, new_balance float64) error {
	tr, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	rows, err := tr.Query(context.Background(),
		`
			UPDATE public.balances
			SET balance = $1
			WHERE userid = $2;
		`,
		new_balance, id,
	)
	if err != nil {
		tr.Rollback(context.Background())
		return err
	}
	rows.Close()

	err = tr.Commit(context.TODO())
	return err
}

// CreateBalance для создания баланса
func (d *DatabaseCRUD) CreateBalance(id string, new_balance float64) error {
	tr, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	rows, err := tr.Query(context.Background(),
		`
		INSERT INTO public.balances (userid, balance)
		SELECT $1, $2
		WHERE NOT EXISTS (SELECT * FROM public.balances WHERE userid = $3);
		`, id, new_balance, id,
	)
	if err != nil {
		tr.Rollback(context.Background())
		return err
	}
	rows.Close()

	err = tr.Commit(context.TODO())
	return err
}

// DeleteBalance для удаления баланса
func (d *DatabaseCRUD) DeleteBalance(id string) error {
	tr, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	rows, err := tr.Query(context.Background(),
		`
		DELETE FROM public.balances 
		WHERE userid = $1;
		`, id,
	)
	if err != nil {
		tr.Rollback(context.Background())
		return err
	}
	rows.Close()

	err = tr.Commit(context.TODO())
	return err
}
