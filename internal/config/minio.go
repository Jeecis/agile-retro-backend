package config

import (
	"github.com/minio/minio-go"
)

func InitializeMinio(cfg *Config) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Minio.MinioHost, cfg.Minio.MinioUser, cfg.Minio.MinioPW, cfg.Minio.SSL)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
