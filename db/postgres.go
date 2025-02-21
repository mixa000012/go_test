package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitPostgres(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}
	return db
}
