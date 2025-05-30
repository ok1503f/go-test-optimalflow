# Go Test OptimalFlow API

A simple RESTful API for user management with authentication, built with Go and Fiber.

## Prerequisites

- Go 1.16 or higher
- PostgreSQL database
- Git

## Setup

1. Clone the repository:

   ```bash
   git clone git@github.com:ok1503f/go-test-optimalflow.git
   cd go-test-optimalflow
   go mod download
   ```

2. Set up the database:
   ```bash
   docker-compose up -d
   ```

CREATE TABLE IF NOT EXISTS users (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
email VARCHAR(255) UNIQUE NOT NULL,
password VARCHAR(255) NOT NULL,
balance DECIMAL(10,2) DEFAULT 100.00,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

3. Run the application:
   ```bash
   go run main.go
   ```

## API Test Use Postman

FIRST LOGIN

```bash
POST http://localhost:3000/api/users/login

{
    "email": "test@example.com",
    "password": "password"
}

```

COPY TOKEN FROM RESPONSE

COPY TOKEN TO Authorization

CHOOSE BEARER TOKEN AND PASTE TOKEN

YOU WILL DO THIS FOR ALL REQUESTS

FOR CREATE USER

```bash
POST http://localhost:3000/api/users

{
    "name": "test",
    "email": "test@example.com",
    "password": "password",
    "balance": 100
}

```

FOR GET ALL USERS

```bash
GET http://localhost:3000/api/users
```

FOR GET USER BY ID

```bash
GET http://localhost:3000/api/users/1
```

FOR TRANSFER BALANCE

```bash
POST http://localhost:3000/api/users/transfer

{
    "from_id": 1,
    "to_id": 2,
    "amount": 10
}
```
