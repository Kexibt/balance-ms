package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
}

func NewDatabase(cfg Config) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetConnectionTimeout())
	defer cancel()

	connection, err := connect(ctx, cfg.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		conn: connection,
	}
}

func connect(ctx context.Context, connectionStr string) (conn *pgx.Conn, err error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err = pgx.Connect(ctx, os.Getenv(connectionStr))
			if err == nil {
				return
			}
		}
	}
}
