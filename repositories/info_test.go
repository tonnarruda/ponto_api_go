package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

func TestGetAllInfo(t *testing.T) {
	query := `SELECT versaobd, statusbd, sistema, versaobdbeta, atualizando, fortes, converteponto3
			FROM INFO`
	tests := []struct {
		name          string
		mockQuery     func(mock sqlmock.Sqlmock)
		expectedError error
		expectedInfo  []structs.Info
	}{
		{
			name: "Get info successfuly",
			mockQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"versaobd", "statusdb", "sistema", "versaobdbeta", "atualizando", "fortes", "converteponto3"}).
					AddRow(307, "OK", "PONTO", 0, 0, 0, 1)
				mock.ExpectQuery(query).
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedInfo:  []structs.Info{{VersaoBD: 307, StatusBD: "OK", Sistema: "PONTO", VersaoBDBeta: 0, Atualizando: 0, Fortes: 0, ConvertePonto3: 1}},
		},
		{
			name: "NoRows",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
			expectedInfo:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			tt.mockQuery(mock)
			repo := repositories.NewInfoRepository(db)
			users, err := repo.GetAll()

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedInfo, users)
		})
	}
}
