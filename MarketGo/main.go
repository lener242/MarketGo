package main

import (
	"MarketGo/database"
	routes "MarketGo/routes.go"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.Connect()
	r := routes.SetupRouter()
	r.SetTrustedProxies(nil)
	r.Run(":8080")
}
