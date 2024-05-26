package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	var err error
    connStr := "user=postgres password=admin12345 dbname=postgres sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }


    log.Println("Database connected!")
    return db, nil
}






