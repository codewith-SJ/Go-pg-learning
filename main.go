package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// Connection string
	connStr := "host=localhost port=5432 user=postgres password=mypassword dbname=testdb sslmode=disable"

	// Open connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")
}
