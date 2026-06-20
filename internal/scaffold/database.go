package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AddDatabaseSupport(projectDir, dbType string) error {
	switch strings.ToLower(dbType) {
	case "postgres":
		return addPostgresSupport(projectDir)
	case "sqlite":
		return addSQLiteSupport(projectDir)
	default:
		return fmt.Errorf("unsupported database: %s", dbType)
	}
}

func addPostgresSupport(dir string) error {
	content := `package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5432"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "postgres"),
			getEnv("DB_NAME", "postgres"),
		)
	}

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
`
	databaseDir := filepath.Join(dir, "database")
	if err := os.MkdirAll(databaseDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(databaseDir, "postgres.go"), []byte(content), 0644)
}

func addSQLiteSupport(dir string) error {
	content := `package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "./data.db"
	}

	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Connected to SQLite")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
`
	databaseDir := filepath.Join(dir, "database")
	if err := os.MkdirAll(databaseDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(databaseDir, "sqlite.go"), []byte(content), 0644)
}
