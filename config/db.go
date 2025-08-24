package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// Load .env hanya untuk lokal
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, assuming environment variables are set")
	}

	// Ambil DATABASE_URL
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging DB: ", err)
	}

	fmt.Println("Database connected 🚀")
	DB = db
}
