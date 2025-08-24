package main

import (
	"books-api/config"
	"books-api/routes"
)

func main() {
	config.InitDB()
	config.RunMigrations()
	r := routes.SetupRouter()
	r.Run(":8080")
}
