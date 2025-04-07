package config

import "os"

type Config struct {
	Host      string
	DB        PGConfig
	JWTSecret string
	Minio     MinioConfig
}

type PGConfig struct {
	PGHost     string
	PGPort     string
	PGUser     string
	PGPassword string
	PGName     string
}

type MinioConfig struct {
	MinioUser   string
	MinioPW     string
	MinioBucket string
	MinioHost   string
	SSL         bool
}

func Load() *Config {

	// set up db config
	db := PGConfig{
		PGHost:     os.Getenv("POSTGRES_HOST"),
		PGPort:     os.Getenv("POSTGRES_PORT"),
		PGUser:     os.Getenv("POSTGRES_USER"),
		PGPassword: os.Getenv("POSTGRES_PASSWORD"),
		PGName:     os.Getenv("POSTGRES_DB"),
	}
	sslStr := os.Getenv("MINIO_SSL")
	var ssl bool
	if sslStr == "true" {
		ssl = true
	} else {
		ssl = false
	}

	minioCfg := MinioConfig{
		MinioUser:   os.Getenv("MINIO_ROOT_USER"),
		MinioPW:     os.Getenv("MINIO_PASSWORD"),
		MinioBucket: os.Getenv("MINIO_BUCKET"),
		MinioHost:   os.Getenv("MINIO_HOST"),
		SSL:         ssl,
	}

	// Load configuration from environment variables or file
	return &Config{
		Host:      os.Getenv("HOST"),
		DB:        db,
		JWTSecret: os.Getenv("JWT_SECRET"),
		Minio:     minioCfg,
	}
}
