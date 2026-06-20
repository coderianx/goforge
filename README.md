# GoForge

Quickly bootstrap Go API projects with your preferred web framework.

## Features

- **6 framework templates**: Gin, Fiber, Chi, Echo, Gorilla/Mux, Standard Library
- **Interactive TUI** framework selector via `charmbracelet/huh`
- **Docker support**: Multi-stage Dockerfile + docker-compose.yml included
- **Graceful shutdown**: Proper signal handling in all templates
- **Database scaffolding**: PostgreSQL or SQLite support (optional)
- **Professional structure**: Clean project layout with separated handlers
- **Single binary**: Install via `go install`

## Installation

```bash
go install github.com/coderianx/goforge/cmd/goforge@latest
```

Make sure `$GOPATH/bin` is in your PATH.

## Usage

```bash
# Create a new project (interactive framework selection)
goforge new my-app

# List supported frameworks
goforge list

# Show version
goforge --version

# Create with database support
goforge new my-app --db postgres
goforge new my-app --db sqlite
```

After creation:

```bash
cd my-app
go mod tidy
go run .
```

## Docker

```bash
# Build and run with Docker Compose
docker compose up --build
```

## Supported Frameworks

| Framework | Port | Database Support |
|-----------|------|-----------------|
| Gin | 8080 | postgres, sqlite |
| Fiber | 3000 | postgres, sqlite |
| Chi | 8080 | postgres, sqlite |
| Echo | 8080 | postgres, sqlite |
| Gorilla/Mux | 8080 | postgres, sqlite |
| Standard Library | 8080 | postgres, sqlite |

## Development

```bash
# Build
make build

# Test
make test

# Lint (requires golangci-lint)
make lint

# Run
make run
```

## Project Structure (scaffolded)

```
my-app/
├── main.go              # Entry point with graceful shutdown
├── go.mod
├── hello.go             # Hello handler
├── ping.go              # Ping/health handler
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Docker Compose config
├── .env.example         # Environment variables
└── database/            # (optional) Database support
    └── postgres.go / sqlite.go
```

## License

[MIT](./LICENSE)
