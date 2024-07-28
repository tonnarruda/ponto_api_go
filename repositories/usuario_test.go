package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

func TestGetByCodigo(t *testing.T) {
	tests := []struct {
		name          string
		code          string
		mockQuery     func(mock sqlmock.Sqlmock)
		expectedError error
		expectedUser  *structs.UserResponse
	}{
		{
			name: "Get Users by code successfuly",
			code: "ADMIN",
			mockQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "codigo"}).
					AddRow("1", "ADMIN")
				mock.ExpectQuery(`SELECT id, codigo FROM USUARIO WHERE codigo = \$1`).
					WithArgs("ADMIN").
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedUser: &structs.UserResponse{
				ID:     "1",
				Codigo: "ADMIN",
			},
		},
		{
			name: "NoRows",
			code: "12345",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id, codigo FROM USUARIO WHERE codigo = \$1`).
					WithArgs("12345").
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: nil,
			expectedUser:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new mock database
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockQuery(mock)
			repo := repositories.NewUserRepository(db)
			user, err := repo.GetByCodigo(tt.code)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedUser, user)
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name          string
		mockQuery     func(mock sqlmock.Sqlmock)
		expectedError error
		expectedUsers []structs.UserResponse
	}{
		{
			name: "Get all users successfully",
			mockQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "codigo"}).
					AddRow("1", "ADMIN").
					AddRow("2", "USER")
				mock.ExpectQuery(`SELECT id, codigo FROM USUARIO`).
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedUsers: []structs.UserResponse{
				{ID: "1", Codigo: "ADMIN"},
				{ID: "2", Codigo: "USER"},
			},
		},
		{
			name: "NoRows",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id, codigo FROM USUARIO`).
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
			expectedUsers: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockQuery(mock)
			repo := repositories.NewUserRepository(db)
			users, err := repo.GetAll()

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedUsers, users)
		})
	}
}

func TestCreateUsers(t *testing.T) {
	tests := []struct {
		name          string
		mockQuery     func(mock sqlmock.Sqlmock)
		expectedError error
		expectedUsers []structs.UserResponse
	}{
		{
			name: "Create users successfully",
			mockQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "codigo"}).
					AddRow("1", "ADMIN").
					AddRow("2", "USER")
				mock.ExpectQuery(`SELECT id, codigo FROM USUARIO`).
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedUsers: []structs.UserResponse{
				{ID: "1", Codigo: "ADMIN"},
				{ID: "2", Codigo: "USER"},
			},
		},
		{
			name: "NoRows",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id, codigo FROM USUARIO`).
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
			expectedUsers: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new mock database
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockQuery(mock)
			repo := repositories.NewUserRepository(db)
			users, err := repo.GetAll()

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedUsers, users)
		})
	}
}
