package database

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

func Open() {
	// connect
	var err error
	DB, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Panic(err)
	}
}

func Close() {
	DB.Close()
}
