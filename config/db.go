package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load .env untuk lokal
	_ = godotenv.Load()

	// Ambil env vars
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database env vars not set")
	}

	// Buat DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	// Connect DB
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	// Ping DB
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging DB: ", err)
	}

	fmt.Println("Database connected 🚀")

	// Jalankan migration otomatis
	migrationsDir := "./database/migrations"
	if err := goose.Up(DB, migrationsDir); err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Migration complete ✅")
}
