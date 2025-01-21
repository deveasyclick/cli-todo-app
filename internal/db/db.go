package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/yusufniyi/cli-todo-app/internal/config"
)

type DB struct {
	Conn *pgx.Conn
}

var (
	DBInstance *DB
)

func ConnectDatabase() {
	dbConn, err := pgx.Connect(context.Background(), config.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unale to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected")
	DBInstance = &DB{Conn: dbConn}

	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	// Create the todo table
	createTodoTable := `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			status VARCHAR(50) DEFAULT 'in-progress',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`
	_, err = dbConn.Exec(context.Background(), createUserTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %v\n", err)
	}

	_, err = dbConn.Exec(context.Background(), createTodoTable)
	if err != nil {
		log.Fatalf("Failed to create todos table: %v\n", err)
	}

	fmt.Println("Tables created successfully!")
}

func (db *DB) Close() {
	DBInstance.Conn.Close(context.Background())
}
