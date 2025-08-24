package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// InitDB connects to the database
func InitDB() {
	// Load .env lokal
	_ = godotenv.Load()

	var connStr string
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		connStr = dbURL
		// Railway biasanya butuh sslmode=require
		if !strings.Contains(connStr, "sslmode") {
			connStr += "?sslmode=require"
		}
	} else {
		port := os.Getenv("PGPORT")
		if port == "" {
			port = "5432"
		}

		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("PGHOST"),
			port,
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGDATABASE"),
		)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging DB: ", err)
	}

	fmt.Println("Database connected 🚀")
	DB = db
}

func RunMigrations() {
	files, err := filepath.Glob("database/migrations/*.sql")
	if err != nil {
		log.Fatal("Error reading migrations folder: ", err)
	}

	for _, file := range files {
		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal("Error reading file: ", err)
		}

		queries := strings.Split(string(sqlBytes), ";")
		for _, query := range queries {
			query = strings.TrimSpace(query)
			if query == "" {
				continue
			}
			_, err := DB.Exec(query)
			if err != nil {
				log.Fatalf("Error executing migration %s: %v", file, err)
			}
		}
		fmt.Println("✅ Migrated:", file)
	}
}
