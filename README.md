# LooksReady: Go Microservices Starter

This repository contains a Go-based microservices application with PostgreSQL, Kafka, and HTTP/3 support. It is designed for rapid development and deployment using Docker Compose.

## Features
- **User Service**: Handles user registration, login, and authentication.
- **Appointment Service**: Manages appointments and integrates with Kafka for messaging.
- **PostgreSQL Database**: Centralized data storage for services.
- **Kafka**: Event streaming for microservices communication.
- **HTTP/3**: Secure, modern protocol support for the main app.
- **Environment-based configuration**: All secrets and settings via `.env` files.

## Project Structure
```
├── main.go                # Main entrypoint for the app (HTTP/3 server)
├── Dockerfile             # Build and run the main app
├── docker-compose.yml     # Multi-service orchestration
├── .env                   # Environment variables for DB and app
├── config/                # Database connection logic
├── controllers/           # User controller logic
├── models/                # Data models (User)
├── routes/                # Gin route definitions
├── utils/                 # JWT and middleware utilities
├── Appointment_Service/   # Appointment microservice
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
├── user_service/          # User microservice
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
├── init/full_dump.sql     # PostgreSQL schema and data initialization
├── cert.pem, key.pem      # TLS certificates for HTTP/3
```

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Go 1.23+

### Setup & Run
1. **Clone the repository**
   ```sh
   git clone <your-repo-url>
   cd firstApp
   ```
2. **Configure environment variables**
   - Edit `.env` for database and app settings.
3. **Build and start all services**
   ```sh
   docker-compose up --build
   ```
4. **Access services**
   - Main app: [https://localhost:8000](https://localhost:8000)
   - Appointment service: [http://localhost:8001](http://localhost:8001)
   - User service: [http://localhost:8002](http://localhost:8002)

### Database Initialization
- The `init/full_dump.sql` file sets up roles, databases, schemas, and tables for PostgreSQL.
- On first run, the database is initialized automatically.

### API Endpoints
#### Main App
- `POST /users/register` — Register a new user
- `POST /users/login` — Login and get JWT
- `POST /users/protected` — Protected route (JWT required)

#### Appointment Service
- `GET /health` — Health check
- `GET /api/appointments` — List appointments
- `GET /api/appointments/:id` — Get appointment and push to Kafka

#### User Service
- `GET /health` — Health check

### Development
- All Go modules are managed via `go.mod` in each service.
- TLS certificates (`cert.pem`, `key.pem`) are required for HTTP/3.
- Environment variables are loaded using `github.com/joho/godotenv`.

### Testing
- Use `curl` or Postman to interact with endpoints.
- Run unit tests with `go test ./...` (add tests as needed).

### Security
- JWT authentication for protected routes.
- Passwords and secrets managed via `.env` (do not commit secrets to public repos).

### Contributing
1. Fork the repo
2. Create a feature branch
3. Commit your changes
4. Open a pull request

### License
MIT

---

**Contact:** For questions or support, open an issue or contact the maintainer.
