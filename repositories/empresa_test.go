package repositories_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

func TestCreateCompany_Success(t *testing.T) {
	dataEncerramento, _ := time.Parse("2006-01-02", "2023-01-01")
	ultimaAtualizacaoAC, _ := time.Parse("2006-01-02", "2023-01-01")
	dataAdesaoESocial, _ := time.Parse("2006-01-02", "2023-01-01")
	dataAdesaoESocialF2, _ := time.Parse("2006-01-02", "2023-01-01")
	convertTipoHe := 1
	usuCodigo := ""

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCompanyRepository(db)

	query := `INSERT INTO EMPRESA \( Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO, CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC, Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL, DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP, NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update \) VALUES \( \$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10, \$11, \$12, \$13, \$14, \$15, \$16, \$17, \$18 \) RETURNING id`

	company := &structs.Empresa{
		Codigo:              "123",
		Nome:                "Test Company",
		RazaoSocial:         "Test Company Ltd",
		CNPJBase:            "123456789",
		USUCodigo:           &usuCodigo,
		ConvertTipoHe:       convertTipoHe,
		CPF:                 "12345678901",
		DataEncerramento:    &dataEncerramento,
		UltimaAtualizacaoAC: &ultimaAtualizacaoAC,
		FaltaAjustarNoAC:    0,
		AderiuESocial:       1,
		DataAdesaoESocial:   &dataAdesaoESocial,
		DataAdesaoESocialF2: &dataAdesaoESocialF2,
		TpAmbESocial:        2,
		StatusEnvioApp:      1,
		NomeFantasia:        "Test",
		CNPJLicenciado:      "12345678901234",
		FreemiumLastUpdate:  "2023-01-01",
	}

	mock.ExpectExec(query).
		WithArgs(
			company.Codigo, company.Nome, company.RazaoSocial, company.CNPJBase, company.USUCodigo,
			company.ConvertTipoHe, company.CPF, company.DataEncerramento, company.UltimaAtualizacaoAC,
			company.FaltaAjustarNoAC, company.AderiuESocial, company.DataAdesaoESocial, company.DataAdesaoESocialF2,
			company.TpAmbESocial, company.StatusEnvioApp, company.NomeFantasia, company.CNPJLicenciado,
			company.FreemiumLastUpdate,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(company)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateCompany_Failure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewCompanyRepository(db)

	query := `INSERT INTO EMPRESA \( Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO, CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC, Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL, DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP, NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update \) VALUES \( \$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10, \$11, \$12, \$13, \$14, \$15, \$16, \$17, \$18 \) RETURNING id`

	company := &structs.Empresa{
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

	mock.ExpectExec(query).
		WithArgs(
			company.Codigo, company.Nome, company.RazaoSocial, company.CNPJBase, company.USUCodigo,
			company.ConvertTipoHe, company.CPF, company.DataEncerramento, company.UltimaAtualizacaoAC,
			company.FaltaAjustarNoAC, company.AderiuESocial, company.DataAdesaoESocial, company.DataAdesaoESocialF2,
			company.TpAmbESocial, company.StatusEnvioApp, company.NomeFantasia, company.CNPJLicenciado,
			company.FreemiumLastUpdate,
		).
		WillReturnError(fmt.Errorf("Failed to insert company into database"))

	err = repo.Create(company)
	assert.Error(t, err)
	assert.Equal(t, "Failed to insert company into database", err.Error())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
