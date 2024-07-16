package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func SetupDatabase() (*sql.DB, error) {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the connection is active
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Get absolute path for migrations directory
	migrationsPath, err := filepath.Abs("db/migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for migrations: %v", err)
	}

	// Print migrations directory path
	fmt.Printf("Migrations directory path: %s\n", migrationsPath)

	// Migrate database
	m, err := migrate.New("file://"+migrationsPath, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("failed to apply migrations: %v", err)
	}

	return db, nil
}
