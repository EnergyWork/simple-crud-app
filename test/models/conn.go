package models

import (
	"database/sql"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/pkg/database"
	"testing"

	_ "github.com/lib/pq"
)

func GetConn(t *testing.T, configFile string) *sql.DB {
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := database.NewPostgresConnection(cfg)
	if err != nil {
		t.Fatal(err)
	}

	return conn
}
