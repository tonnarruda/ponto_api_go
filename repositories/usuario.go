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

func (r *UserRepository) GetAll() ([]structs.UserResponse, error) {
	query := `SELECT id, codigo
			FROM USUARIO`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to fetch users from database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []structs.UserResponse
	for rows.Next() {
		var user structs.UserResponse
		err := rows.Scan(&user.ID, &user.Codigo)
		if err != nil {
			log.Printf("Failed to scan company: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByCodigo(codigo string) (*structs.UserResponse, error) {
	query := `SELECT id, codigo 
			FROM USUARIO 
			WHERE codigo = $1`
	row := r.db.QueryRow(query, codigo)

	var user structs.UserResponse
	err := row.Scan(&user.ID, &user.Codigo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Empresa não encontrada
		}
		log.Printf("Failed to fetch user by code: %v", err)
		return nil, err
	}

	return &user, nil
}
