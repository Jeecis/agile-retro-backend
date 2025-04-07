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
		log.Fatal(err)
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := config.InitializeDB(cfg)
	if err != nil {
		log.Fatal("Postgres error: " + err.Error())
	}

	minio, err := config.InitializeMinio(cfg)
	if err != nil {
		log.Fatal("Minio error: " + err.Error())
	}

	// Setup router
	r := routes.SetupRouter(db, minio)

	// Start server
	r.Run(":8080")
}
