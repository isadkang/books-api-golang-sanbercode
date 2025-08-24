package main

import (
	"books-api/config"
	"books-api/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
