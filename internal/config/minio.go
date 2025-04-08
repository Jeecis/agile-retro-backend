package config

import (
	"github.com/minio/minio-go"
)

// InitializeMinio initializes and returns a MinIO client instance.
// MinIO is a high-performance, S3-compatible object storage system.
// This function takes a configuration object and uses it to create a new MinIO client.
//
// Parameters:
//   - cfg: A pointer to a Config struct containing MinIO connection details such as
//     host, user credentials, and SSL settings.
//
// Returns:
//   - *minio.Client: A pointer to the initialized MinIO client.
//   - error: An error if the client initialization fails.
func InitializeMinio(cfg *Config) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Minio.MinioHost, cfg.Minio.MinioUser, cfg.Minio.MinioPW, cfg.Minio.SSL)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
