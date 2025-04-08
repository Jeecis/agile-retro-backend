# Go api for Agile retrospective board

This is go API for Agile Retrospective board


### Build

```
docker compose up -d
```

### Routes

###### Health Check
- `GET /api/v0/health` - Check if the API is running

###### Board Routes
- `POST /api/v0/board` - Create a new board
- `GET /api/v0/board` - Get all boards
- `GET /api/v0/board/{id}` - Get a specific board by ID
- `PUT /api/v0/board` - Update an existing board
- `DELETE /api/v0/board` - Delete a board

###### Column Routes
- `POST /api/v0/column` - Create a new column
- `GET /api/v0/column` - Get all columns
- `GET /api/v0/column/{id}` - Get a specific column by ID
- `PUT /api/v0/column` - Update an existing column
- `DELETE /api/v0/column` - Delete a column

###### Record Routes
- `POST /api/v0/record` - Create a new record
- `GET /api/v0/record` - Get all records
- `GET /api/v0/record/{id}` - Get a specific record by ID
- `PUT /api/v0/record` - Update an existing record
- `DELETE /api/v0/record` - Delete a record

### Future considorations

This API also has partial implementation for File storage and user management and authetification with JWT, while this is not implemented yet, it will be in future.
