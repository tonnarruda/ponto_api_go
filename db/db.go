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

// SetupDatabase configura a conexão com o banco de dados PostgreSQL e executa migrações.
func SetupDatabase() (*sql.DB, error) {
	// Obter a URL do banco de dados
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL not set in environment variables")
	}

	// Abrir conexão com o banco de dados
	db, err := openDatabase(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Executar migrações
	if err := runMigrations(databaseURL); err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %v", err)
	}

	return db, nil
}

// openDatabase abre uma conexão com o banco de dados PostgreSQL.
func openDatabase(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// runMigrations executa as migrações do banco de dados.
func runMigrations(databaseURL string) error {
	// Obter caminho absoluto para o diretório de migrações
	migrationsPath, err := filepath.Abs("db/migrations")
	if err != nil {
		return err
	}

	// Criar instância do migrator apontando para o diretório de migrações
	m, err := migrate.New("file://"+migrationsPath, databaseURL)
	if err != nil {
		return err
	}

	// Executar migrações pendentes
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
