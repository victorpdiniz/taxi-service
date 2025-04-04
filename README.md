# Go API Base Project

A clean and well-structured Go API base project that provides a solid foundation for building RESTful APIs. This project uses Fiber as the web framework and GORM for database operations with PostgreSQL.

## 🚀 Features

- 🏗️ Clean architecture with separation of concerns
- 🔐 Environment variable configuration
- 🗄️ PostgreSQL database integration with GORM
- 🛣️ Organized routing system
- 🛡️ Middleware support
- 📦 Dependency management with Go modules
- ✅ End-to-end testing with SQLite

## 📋 Prerequisites

- Go 1.23.5 or higher
- PostgreSQL database
- Basic understanding of Go programming

## 🛠️ Project Structure

```
.
├── app.go              # Main application entry point
├── controllers/        # Request handlers
├── database/          # Database configuration and models
├── middlewares/       # Custom middleware functions
├── models/            # Data models
├── routes/            # Route definitions
├── services/          # Business logic
├── test/              # Test utilities and e2e tests
│   ├── e2e/          # End-to-end tests
│   └── test_helper.go # Test utilities
└── .env               # Environment variables
```

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone git@github.com:Mlcarvalho1/golang-fiber-base.git
   cd golang-api-base
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory with the following variables:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database
   ```

4. Run the application:
   ```bash
   go run app.go
   ```

The server will start on `http://localhost:3000`

## 🧪 Testing

The project includes end-to-end tests using SQLite for the test database. This allows for fast and isolated testing without affecting your production database.

### Running Tests

To run all tests:
```bash
go test ./test/e2e/... -v
```

To run a specific test:
```bash
go test ./test/e2e/... -v -run TestDummyCRUD
```

### Writing Tests

The test infrastructure provides helper functions to make writing e2e tests easier:

1. `test.SetupTestApp(t)`: Creates a new Fiber app with SQLite database
2. `test.MakeRequest(t, app, method, path, body)`: Makes HTTP requests to the test app
3. `test.ParseResponseBody(t, resp, v)`: Parses response body into a struct

Example test structure:
```go
func TestYourFeature(t *testing.T) {
    app := test.SetupTestApp(t)
    defer test.CleanupTestApp(t)

    // Your test code here
}
```

### Test Database

Tests use SQLite instead of PostgreSQL to:
- Speed up test execution
- Provide isolated test environments
- Avoid affecting production data
- Enable parallel test execution

## 🛠️ Dependencies

- [Fiber](https://github.com/gofiber/fiber) - Fast and efficient web framework
- [GORM](https://gorm.io/) - ORM for database operations
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader
- [testify](https://github.com/stretchr/testify) - Testing utilities

## 📚 Project Architecture

- **Controllers**: Handle HTTP requests and responses
- **Models**: Define data structures and database schemas
- **Services**: Implement business logic
- **Routes**: Define API endpoints and middleware
- **Database**: Configure database connections and migrations
- **Middlewares**: Implement cross-cutting concerns
- **Tests**: End-to-end tests with SQLite

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🙏 Acknowledgments

- Thanks to the Go community for their amazing tools and libraries
- Special thanks to the Fiber and GORM teams for their excellent work
