package repositories_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

const queryString = `INSERT INTO EMPRESA \( Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO, CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC, Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL, DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP, NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update \) VALUES \( \$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10, \$11, \$12, \$13, \$14, \$15, \$16, \$17, \$18 \) RETURNING id`

type TestCase struct {
	name        string
	setupMock   func(sqlmock.Sqlmock, *structs.Empresa)
	expectedErr error
}

// Helper function to create a sample company
func getSampleCompany() *structs.Empresa {
	return &structs.Empresa{
		Codigo:              "123",
		Nome:                "Test Company",
		RazaoSocial:         "Test Company Ltd",
		CNPJBase:            "123456789",
		USUCodigo:           nil,
		ConvertTipoHe:       1,
		CPF:                 "12345678901",
		DataEncerramento:    nil,
		UltimaAtualizacaoAC: nil,
		FaltaAjustarNoAC:    0,
		AderiuESocial:       1,
		DataAdesaoESocial:   nil,
		DataAdesaoESocialF2: nil,
		TpAmbESocial:        2,
		StatusEnvioApp:      1,
		NomeFantasia:        "Test",
		CNPJLicenciado:      "12345678901234",
		FreemiumLastUpdate:  "2023-01-01",
	}
}

// RunTestCase executes a test case
func RunTestCase(t *testing.T, tc TestCase) {
	t.Run(tc.name, func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := repositories.NewCompanyRepository(db)
		company := getSampleCompany()

		tc.setupMock(mock, company)

		err = repo.Create(company)
		if tc.expectedErr != nil {
			assert.Error(t, err)
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.NoError(t, err)
		}

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestCreateCompany(t *testing.T) {
	testCases := []TestCase{
		{
			name: "Success",
			setupMock: func(mock sqlmock.Sqlmock, company *structs.Empresa) {
				query := queryString
				mock.ExpectExec(query).
					WithArgs(
						company.Codigo, company.Nome, company.RazaoSocial, company.CNPJBase, company.USUCodigo,
						company.ConvertTipoHe, company.CPF, company.DataEncerramento, company.UltimaAtualizacaoAC,
						company.FaltaAjustarNoAC, company.AderiuESocial, company.DataAdesaoESocial, company.DataAdesaoESocialF2,
						company.TpAmbESocial, company.StatusEnvioApp, company.NomeFantasia, company.CNPJLicenciado,
						company.FreemiumLastUpdate,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name: "Failure",
			setupMock: func(mock sqlmock.Sqlmock, company *structs.Empresa) {
				query := queryString
				mock.ExpectExec(query).
					WithArgs(
						company.Codigo, company.Nome, company.RazaoSocial, company.CNPJBase, company.USUCodigo,
						company.ConvertTipoHe, company.CPF, company.DataEncerramento, company.UltimaAtualizacaoAC,
						company.FaltaAjustarNoAC, company.AderiuESocial, company.DataAdesaoESocial, company.DataAdesaoESocialF2,
						company.TpAmbESocial, company.StatusEnvioApp, company.NomeFantasia, company.CNPJLicenciado,
						company.FreemiumLastUpdate,
					).
					WillReturnError(fmt.Errorf("Failed to insert company into database"))
			},
			expectedErr: fmt.Errorf("Failed to insert company into database"),
		},
	}

	for _, tc := range testCases {
		RunTestCase(t, tc)
	}
}
