# Go api for Agile retrospective board

This is go API for Agile Retrospective board


### Build

```
docker compose up -d
```

or change .env.example to .env and press `F5`

or cd int cmd/api and run `go run main.go`

### Used libraries
* github.com/gin-gonic/gin: A web framework for building APIs in Go. It provides routing, middleware support, and HTTP handling.

* github.com/gin-contrib/cors: A middleware for handling Cross-Origin Resource Sharing (CORS) in Gin applications.

* gorm.io/gorm: An ORM (Object-Relational Mapper) library for Go, used for interacting with databases.

* gorm.io/driver/postgres: A GORM driver for PostgreSQL, enabling database connectivity.

* github.com/joho/godotenv: A library for loading environment variables from a .env file.

* github.com/gorilla/websocket: A library for handling WebSocket connections in Go.

* github.com/google/uuid: A library for generating UUIDs.