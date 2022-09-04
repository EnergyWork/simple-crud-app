package models

import (
	"database/sql"
	"simple-crud-app/internal/lib/config"
	"simple-crud-app/pkg/database"
	"testing"

	_ "github.com/lib/pq"
)

const (
	PRIVATE_KEY = "MIICXQIBAAKBgQDfbBzsQGdwdMJfUNXwzRujXZ05ddDIb9mj7odeZNcMzn+lmsz6XXsvWviIb4/IRqOJDGRB5PRgt4m+R46Lca3vef2z+PbyNvjzusv3YVIXnuBgbSMkmFL9OryfE7vkCOO2nLvCLeVd8bcIuKVFXN6Q9PREI9ej7aRKxX+I5V7xGQIDAQABAoGABg+20SoGJGTmiRN2WmwWHd6CT3bEzUtLikkEXyk5NF291M5YVUqH9wbuyzTLn9FaynMNnUQK5TzVfdYPJfVVlKat8WJ5Un1xJtoZsTp9IFpLSLWvCoom2ablcLEq7Fl3crHAk4e5Zr1HzgbNMqlEYmRFBG0fDit6GHRZ/FsRDSkCQQDz3JlyLj+htsWZThyEV9GchCN+6UYwv1UJONNrzM26VbcOmh5EzGbS6YfRIm1OU8Jrqdwlszn1ZeLTCfVa2B7nAkEA6osQrPs9j0NKK8ernGwvoagdsOzzQRPUSdWi9y3uzLZKQMYryUmQgci1OfsAl8AhFnplYZCApGh1asPM+35v/wJAQtUddK56H+7AXtCKfja3KqcIN1rlMqztODbLsoqRg1TEc4sHaqF+OKVp5IYD4OiRqwIFZIunAbsnm+DpzjjW1wJBAJXv4PE0i94R/lCOjL6qyqhleNWqJLftnUC2OkAaNRbZUg6moUdEqATP8krmkzJvuLdN95Gvdw2jWayvD1OXOLECQQDnAgwo3yzKbVLvFDjKJS2khEafO20LANkuoCeZ94Zm1EnyTxHWt3Jq8UGOpOsiNIq5GrhyJ7JhidOQ4cva/7Y1"
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
