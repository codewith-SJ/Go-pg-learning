package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=mypassword dbname=testdb sslmode=disable"
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
	fmt.Println("✅ Connected to PostgreSQL!")

	// Create table
	createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Table 'users' created!")

	// Insert data
	insertSQL := `INSERT INTO users (name, email) VALUES ($1, $2)`

	_, err = db.Exec(insertSQL, "John Doe", "john@example.com")
	if err != nil {
		fmt.Println("⚠️  Error inserting John:", err)
	} else {
		fmt.Println("✅ User 'John Doe' inserted!")
	}

	_, err = db.Exec(insertSQL, "Alice Smith", "alice@example.com")
	if err != nil {
		fmt.Println("⚠️  Error inserting Alice:", err)
	} else {
		fmt.Println("✅ User 'Alice Smith' inserted!")
	}

	_, err = db.Exec(insertSQL, "Bob Johnson", "bob@example.com")
	if err != nil {
		fmt.Println("⚠️  Error inserting Bob:", err)
	} else {
		fmt.Println("✅ User 'Bob Johnson' inserted!")
	}

	// Read all users
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("\n📋 Users in database:")
	fmt.Println("-----------------------------------")
	for rows.Next() {
		var id int
		var name, email string
		var createdAt string
		rows.Scan(&id, &name, &email, &createdAt)
		fmt.Printf("ID: %d | Name: %s | Email: %s\n", id, name, email)
	}
	fmt.Println("-----------------------------------")
}
