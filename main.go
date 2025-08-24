package main

import (
	"books-api/config"
	"books-api/routes"
	"os"
)

func main() {
	config.InitDB()
	config.RunMigrations()
	r := routes.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
