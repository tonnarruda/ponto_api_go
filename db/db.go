package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// SetupDatabase configura a conexão com o banco de dados
func SetupDatabase() (*sql.DB, error) {
	// Conecta ao banco de dados PostgreSQL
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão está ativa
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Cria as tabelas no banco de dados (se necessário)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produtos (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(255) NOT NULL
    )`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
