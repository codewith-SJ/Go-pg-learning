package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

// InitDB initializes database connection
func InitDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=mypassword dbname=testdb sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	createTable()
}

// CreateTable creates the users table if it doesn't exist
func createTable() {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("✅ Table ready")
}

// GetAllUsers fetches all users from database
func GetAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
