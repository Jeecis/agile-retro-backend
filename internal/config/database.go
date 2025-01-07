package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitializeDB(cfg *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.PGUser, cfg.DB.PGPassword, cfg.DB.PGHost, cfg.DB.PGPort, cfg.DB.PGName)

	log.Print(dsn)

	// Open a connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}

	log.Print("Connected to postgres DB!")
	return db, nil
}
