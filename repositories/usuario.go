package repositories

import (
	"database/sql"
	"log"

	"github.com/tonnarruda/ponto_api_go/structs"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *structs.Usuario) error {
	query := `
		INSERT INTO USUARIO (
			Codigo, Senha, UltimoAcesso, Bloqueado, UserRegistrationDate, LimiteEpgData
		) VALUES ($1, $2, $3, $4, $5, $6
		)
		RETURNING id
	`
	_, err := r.db.Exec(query, user.Codigo, user.Senha, user.UltimoAcesso,
		user.Bloqueado, user.UserRegistrationDate, user.LimiteEpgData)
	if err != nil {
		log.Printf("Failed to insert user into database: %v", err)
		return err
	}
	return nil
}
