package main

import (
	"log"

	"github.com/Jeecis/goapi/internal/api/routes"
	"github.com/Jeecis/goapi/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := config.InitializeDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Setup router
	r := routes.SetupRouter(db)

	// Start server
	r.Run(":8080")
}
