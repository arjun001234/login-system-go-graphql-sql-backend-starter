package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(connString string) *sql.DB {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Panicf("Database Error: %v", err.Error())
	}
	err = conn.Ping()
	if err != nil {
		log.Panicf("Database Error: %v", err.Error())
	}
	return conn
}
