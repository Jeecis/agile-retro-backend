package config

import "os"

type Config struct {
	Host      string
	DB        PGconfig
	JWTSecret string
}

type PGconfig struct {
	PGHost     string
	PGPort     string
	PGUser     string
	PGPassword string
	PGName     string
}

func Load() *Config {

	// set up db config
	db := PGconfig{
		PGHost:     os.Getenv("POSTGRES_HOST"),
		PGPort:     os.Getenv("POSTGRES_PORT"),
		PGUser:     os.Getenv("POSTGRES_USER"),
		PGPassword: os.Getenv("POSTGRES_PASSWORD"),
		PGName:     os.Getenv("POSTGRES_DB"),
	}

	// Load configuration from environment variables or file
	return &Config{
		Host:      os.Getenv("HOST"),
		DB:        db,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
