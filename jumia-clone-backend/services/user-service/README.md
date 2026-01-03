# User Service

A gRPC microservice for user management in the Jumia clone application.

## Features

- User registration with email and password
- User authentication with JWT tokens
- User profile management (get, update, delete)
- Token verification for authentication
- Password hashing with bcrypt
- PostgreSQL database with GORM

## Setup

1. Install dependencies:

```bash
go mod tidy
```

2. Generate protobuf code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/user.proto
```

3. Configure environment:

```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. Run the service:

```bash
go run cmd/main.go
```

## API Endpoints (gRPC)

### Register

- **Request**: `RegisterRequest`
- **Response**: `RegisterResponse`
- Creates a new user account

### Login

- **Request**: `LoginRequest`
- **Response**: `LoginResponse`
- Authenticates user and returns JWT tokens

### GetUser

- **Request**: `GetUserRequest`
- **Response**: `GetUserResponse`
- Retrieves user information by ID

### UpdateUser

- **Request**: `UpdateUserRequest`
- **Response**: `UpdateUserResponse`
- Updates user profile information

### DeleteUser

- **Request**: `DeleteUserRequest`
- **Response**: `DeleteUserResponse`
- Soft deletes a user account

### VerifyToken

- **Request**: `VerifyTokenRequest`
- **Response**: `VerifyTokenResponse`
- Verifies JWT token validity

## Database Schema

The service uses PostgreSQL with the following schema:

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

## Environment Variables

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: postgres)
- `DB_NAME`: Database name (default: jumia_users)
- `GRPC_PORT`: gRPC server port (default: 50051)

## Architecture

```
cmd/
  └── main.go              # Application entry point
internal/
  ├── handler/             # gRPC handlers
  ├── models/              # Database models
  ├── repository/          # Data access layer
  └── service/             # Business logic layer
proto/
  └── user.proto           # Protocol buffer definitions
```
