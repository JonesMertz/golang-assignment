# Golang assignment 💻
A simple Golang application that serves a RESTful API for board games
### Run the following commands
 ```
go mod tidy
go mod download
```

### Configuration    
#### .env example
```
    DB_PATH="./database/database.db"
    PORT="9090"
```

### Do a test run
```
go run main.go
```

## Docker
```
docker build -t golang-assignment .  
docker run -p 9090:9090 golang-assignment
```


_Endpoints:_

```sh
# Get all board games
GET /boardGames

# Get a specific board game
GET /boardGames/:id

# Create a board game
POST /boardGames

# Update a board game
PUT /boardGames/:id

# Delete a board game
DELETE /boardGames/:id
```

_Model:_

```json
{
    "id": 1,
    "name": "Catan",
    "minPlayers": 3,
    "maxPlayers": 4,
    "playTime": 60,
    "age": 10,
    "description": "A fun game",
    "price": 50
}
```

## Libraries used

The service uses the following libraries:
-  Database - https://github.com/mattn/go-sqlite3
-  HTTP server - https://github.com/labstack/echo
-  API testing - https://echo.labstack.com/guide/testing/