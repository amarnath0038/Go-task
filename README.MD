# User CRUD API (Go)

This is a simple backend service written in Go to manage users with their name and DOB.
The API calculates a user’s age dynamically instead of storing it in the DB.

## What this project does

- Create a user with name and DOB
- Fetch all users (with pagination)
- Fetch a single user by ID
- Update a user
- Delete a user
- Calculate age dynamically using Go’s time package

Age is NOT stored in the database. It’s computed at runtime.

## Tech used

- Go
- Fiber
- PostgreSQL
- SQLC
- go-playground/validator
- Uber Zap
- Docker

## Project structure

cmd/server/main.go → app entry point
config/ → DB config
db/  
 ├── migrations/  
 └── sqlc/  
internal/  
 ├── handler/ → HTTP handlers  
 ├── service/ → business logic  
 ├── repository/ → database access  
 ├── routes/ → route registration  
 ├── middleware/ → requestId & timing middleware  
 ├── models/ → request/response models  
 └── logger/ → zap logger setup

## API endpoints

### Create user

POST /users

Request body:
{
"name": "Alice",
"dob": "1990-05-10"
}

Response:
201 Created

### List users (with pagination)

GET /users?page=1&limit=10

Response:
[
{
"id": 1,
"name": "Alice",
"dob": "1990-05-10",
"age": 34
}
]

### Get user by ID

GET /users/{id}

---

### Update user

PUT /users/{id}

Request body:
{
"name": "Alice Updated",
"dob": "1991-03-15"
}

---

### Delete user

DELETE /users/{id}

Response:
204 No Content

---

## Running the project (Docker Compose)

This is the recommended way to run the project locally.

### Requirements

- Docker
- Docker Compose (v2)

### Start the application

docker compose up --build

This will:

- start PostgreSQL
- apply database migrations
- start the API on port 8080

### Stop the application

docker compose down

To remove database data:

docker compose down -v

## Environment variables

The application uses environment variables for configuration.

Example .env file (not committed):

DB_USER=appuser  
DB_PASSWORD=password  
DB_NAME=userdb  
DB_HOST=db  
DB_PORT=5432

## Logging and middleware

- Each request gets a unique request ID
- Request duration is logged
- Logs are structured using Zap

## Tests

There is a small unit test for the age calculation logic.

Run all tests with:

go test ./...
