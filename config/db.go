package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "PGHOST=postgres.railway.internal PGPORT=5432 PGUSER=postgres PGPASSWORD=UJscaMybCULqDLisvrgODIFVnmhgZARg PGDATABASE=railway sslmode=disable"
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
