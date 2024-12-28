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

	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
	 CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(100) NOT NULL,
		name NVARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	 );
	 
	  CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INT NOT NULL,
		title VARCHAR(255) NOT NULL,
    description TEXT,
    is_completed BOOLEAN DEFAULT FALSE,
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	 );
	 `

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}
}

func Close() {
	DB.Close()
}
