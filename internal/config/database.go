package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializeDB establishes a connection to a PostgreSQL database instance,
// typically running in a Docker container, using the provided configuration.
// It constructs a DSN (Data Source Name) from the configuration, opens a connection
// using GORM, and verifies the connection by pinging the database.
//
// Parameters:
//   - cfg: A pointer to a Config struct containing database connection details.
//
// Returns:
//   - *gorm.DB: A pointer to the GORM database instance if the connection is successful.
//   - error: An error if the connection fails or if there are issues during initialization.
func InitializeDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.PGHost, cfg.DB.PGUser, cfg.DB.PGPassword, cfg.DB.PGName, cfg.DB.PGPort)

	// Open a connection using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
