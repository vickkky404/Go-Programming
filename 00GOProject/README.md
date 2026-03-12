# Go Backend

A clean, idiomatic Go backend project...

## Project Structure

```
.
├── main.go                   # Entry point
├── go.mod                    # Module definition
├── .env.example              # Example environment variables
├── cmd/
│   └── api/
│       └── server.go         # HTTP server setup
├── internal/
│   ├── handlers/             # HTTP request handlers
│   │   └── home.go
│   ├── routes/               # Route registration
│   │   └── routes.go
│   ├── middleware/            # HTTP middleware (logging, auth, etc.)
│   │   └── logger.go
│   ├── models/               # Data models / structs
│   │   └── model.go
│   └── config/               # App configuration loader
│       └── config.go
├── pkg/
│   └── utils/                # Shared utility helpers
│       └── response.go
└── config/                   # Config files (YAML, JSON, etc.)
```

## Getting Started

```bash
# Copy env file
cp .env.example .env

# Run the server
go run main.go
```

## Endpoints

| Method | Path      | Description         |
|--------|-----------|---------------------|
| GET    | /         | Welcome message     |
| GET    | /health   | Health check        |
