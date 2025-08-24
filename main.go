// @title Books API
// @version 1.0
// @description API for managing books and categories
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
