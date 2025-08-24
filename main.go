package main

import (
	"books-api/config"
	"books-api/routes"
	"os"
)

func main() {
	config.InitDB() // connect + migrate otomatis

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
