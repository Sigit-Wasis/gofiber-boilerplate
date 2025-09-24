package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


var conn *sql.DB


func Init(databaseURL string) error {
var err error
conn, err = sql.Open("postgres", databaseURL)
if err != nil {
return err
}
if err := conn.Ping(); err != nil {
return fmt.Errorf("db ping: %w", err)
}
return nil
}


func Get() *sql.DB { return conn }


func Close() error {
if conn == nil {
return nil
}
return conn.Close()
}