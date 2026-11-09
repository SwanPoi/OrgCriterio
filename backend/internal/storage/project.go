package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sqlx.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database via pgx: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Storage{db: db}, nil
}
