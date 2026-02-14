# Backend - Go Application

This is the backend service for the Xenova application, built with Go using the Echo framework.

## Project Structure

```
backend/
├── config/           # Configuration files
├── initialize/       # Initialization logic (router, database, redis, viper)
├── internal/         # Internal application code
│   ├── api/         # API handlers
│   ├── global/      # Global variables and types
│   ├── models/      # Database models
│   └── service/     # Business logic services
├── pkg/             # Public packages
│   ├── email/       # Email functionality
│   ├── middleware/  # Middleware (auth, logger)
│   ├── request/     # Request DTOs
│   ├── response/    # Response DTOs
│   └── validator/   # Custom validators
├── router/          # Route definitions
├── utils/           # Utility functions
└── main.go          # Application entry point
```

## Tech Stack

- **Framework**: Echo v5
- **ORM**: GORM
- **Database**: MySQL
- **Cache**: Redis
- **Authentication**: JWT (golang-jwt/jwt)
- **Config Management**: Viper
- **Validation**: go-playground/validator
- **Logging**: Logrus

## Getting Started

### Prerequisites

- Go 1.25.6 or higher
- MySQL database
- Redis server

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   cd backend
   go mod download
   ```

3. Configure your environment (see `config/` directory)

4. Run the application:
   ```bash
   go run main.go
   ```

## Testing

The backend includes comprehensive test coverage for all modules.

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test ./... -v
```

Run tests with coverage:
```bash
go test ./... -cover
```

Generate detailed coverage report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Run tests for a specific package:
```bash
go test ./utils -v
go test ./internal/models -v
go test ./pkg/middleware -v
```

### Test Organization

Tests are organized by module with the following coverage:

#### Utils Package (`utils/`)
- `jwt_test.go` - JWT token operations
- `paginate_test.go` - Pagination utilities
- `request_test.go` - Request parsing and validation
- `response_test.go` - Response formatting
- `user_test.go` - Password hashing and verification

#### Middleware Package (`pkg/middleware/`)
- `auth_test.go` - Authentication middleware

#### Service Layer (`internal/service/`)
- `user_test.go` - User service
- `software_test.go` - Software service
- `product_test.go` - Product service
- `journalism_test.go` - Journalism service

#### API Layer (`internal/api/`)
- `user_test.go` - User API handlers
- `software_test.go` - Software API handlers
- `product_test.go` - Product API handlers
- `journalism_test.go` - Journalism API handlers

#### Models (`internal/models/`)
- `user_test.go` - User model
- `types_test.go` - Custom types (StringList)

#### Other Packages
- `pkg/email/enter_test.go` - Email sender
- `pkg/validator/enter_test.go` - Custom validator
- `pkg/request/user_test.go` - Request DTOs
- `pkg/response/user_test.go` - Response DTOs

### Testing Best Practices

1. **Unit Tests**: Focus on testing individual functions and methods
2. **Table-Driven Tests**: Use table-driven test patterns for multiple test cases
3. **Mock Dependencies**: For tests requiring database or external services, consider using mocks
4. **Test Naming**: Follow the convention `Test<FunctionName>` for test functions
5. **Coverage**: Aim for meaningful test coverage, not just high percentages

### Continuous Integration

Tests are automatically run on:
- Pull requests
- Commits to main branches
- Before deployment

## Development

### Code Quality

Run linting:
```bash
go vet ./...
go fmt ./...
```

### Adding New Tests

When adding new functionality:
1. Create a corresponding `*_test.go` file in the same package
2. Write tests that cover:
   - Happy path scenarios
   - Edge cases
   - Error conditions
3. Run tests locally before committing
4. Ensure all tests pass in CI

## API Modules

The backend provides the following main modules:

- **User Management**: Authentication, registration, user operations
- **Software**: Software resource management with categories
- **Product**: Product catalog with series organization
- **Journalism**: Content management for articles/posts

## License

[Add your license information here]
