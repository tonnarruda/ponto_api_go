package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/ponto_api_go/repositories"
)

func TestGetByCodigo_Success(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected query and result
	code := "ADMIN"
	rows := sqlmock.NewRows([]string{"id", "codigo"}).
		AddRow("1", code)

	mock.ExpectQuery(`SELECT id, codigo FROM USUARIO WHERE codigo = \$1`).
		WithArgs(code).
		WillReturnRows(rows)

	// Create a new UserRepository
	repo := repositories.NewUserRepository(db)

	// Call the method
	user, err := repo.GetByCodigo(code)

	// Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, code, user.Codigo)
}

func TestGetByCodigo_NoRows(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected query and result
	codigo := "12345"
	mock.ExpectQuery(`SELECT id, codigo FROM USUARIO WHERE codigo = \$1`).
		// WithArgs(codigo).
		WillReturnError(sql.ErrNoRows)

	// Create a new UserRepository
	repo := repositories.NewUserRepository(db)

	// Call the method
	user, err := repo.GetByCodigo(codigo)

	// Assert the results
	assert.NoError(t, err)
	assert.Nil(t, user)
}
