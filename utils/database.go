package utils

import (
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"fmt"
	"os"
)

// buildConnectionString builds a connection string from environment variables
func buildConnectionString() (string, error) {
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	PG_USER := os.Getenv("PG_USER")
	PG_DBNAME := os.Getenv("PG_DBNAME")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	PG_HOST := os.Getenv("PG_HOST")
	PG_PORT := os.Getenv("PG_PORT")
	PG_SSLMODE := os.Getenv("PG_SSLMODE")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", PG_USER, PG_DBNAME, PG_PASSWORD, PG_HOST, PG_PORT, PG_SSLMODE)

	return connStr, nil
}

// Database opens a new database connection
func Database() (*sql.DB, error) {
	// Build the connection string
	connStr, err := buildConnectionString()
	if err != nil {
		return nil, err
	}

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		// Close the database connection if Ping fails
		db.Close()
		return nil, err
	}

	return db, nil
}
