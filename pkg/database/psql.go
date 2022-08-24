package database

import (
	"database/sql"
)

type IConfig interface {
	GetDBConnection() string
}

func NewPostgresConnection(cfg IConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDBConnection())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
