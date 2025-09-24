package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(databaseURL string) error {
	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func GetDB() *sql.DB {
	return DB
}
