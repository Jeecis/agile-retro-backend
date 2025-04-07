package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitializeDB(cfg *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.PGUser, cfg.DB.PGPassword, cfg.DB.PGHost, cfg.DB.PGPort, cfg.DB.PGName)

	// Open a connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
