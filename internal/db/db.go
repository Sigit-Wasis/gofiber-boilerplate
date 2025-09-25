package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(databaseURL string) error {
	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(20)
    DB.SetMaxIdleConns(5)
    DB.SetConnMaxLifetime(time.Hour)
	
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
