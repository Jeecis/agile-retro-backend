package main

import (
	"log"

	"github.com/Jeecis/goapi/internal/api/routes"
	"github.com/Jeecis/goapi/internal/config"
	"github.com/Jeecis/goapi/internal/models"
	"github.com/Jeecis/goapi/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := config.InitializeDB(cfg)
	if err != nil {
		log.Fatal("Postgres error: " + err.Error())
	}

	// Run database migrations
	if err := db.AutoMigrate(&models.Board{}, &models.Column{}, &models.Record{}); err != nil {
		log.Fatal("Migration error: " + err.Error())
	}

	minio, err := config.InitializeMinio(cfg)
	if err != nil {
		log.Fatal("Minio error: " + err.Error())
	}

	boardRepo := repository.NewBoardRepository(db)
	columnRepo := repository.NewColumnRepository(db)
	recordRepo := repository.NewRecordRepository(db)

	// Setup router
	r := routes.SetupRouter(db, minio, boardRepo, columnRepo, recordRepo)

	// Start server
	r.Run(":8080")
}
