package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Init() *sql.DB {
	dsn := "postgres://user:password@localhost:5432/myapp?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("DB open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	return db
}
