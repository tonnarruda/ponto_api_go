package repositories

import (
	"database/sql"
	"log"

	"github.com/tonnarruda/products_go/models"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company *models.Empresa) error {
	query := `
		INSERT INTO EMPRESA (
			Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
			CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
			Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
			DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
			NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18
		)
		RETURNING id
	`
	_, err := r.db.Exec(query, company.Codigo, company.Nome, company.RazaoSocial, company.CNPJBase, company.USUCodigo, company.ConvertTipoHe,
		company.CPF, company.DataEncerramento, company.UltimaAtualizacaoAC, company.FaltaAjustarNoAC, company.AderiuESocial, company.DataAdesaoESocial,
		company.DataAdesaoESocialF2, company.TpAmbESocial, company.StatusEnvioApp, company.NomeFantasia, company.CNPJLicenciado, company.FreemiumLastUpdate)
	if err != nil {
		log.Printf("Failed to insert user into database: %v", err)
		return err
	}
	return nil
}

func (r *CompanyRepository) GetAll() ([]models.Empresa, error) {
	query := `SELECT id, Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				   CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				   Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				   DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				   NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update
			FROM EMPRESA`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to fetch companies from database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var companies []models.Empresa
	for rows.Next() {
		var company models.Empresa
		err := rows.Scan(&company.ID, &company.Codigo, &company.Nome, &company.RazaoSocial, &company.CNPJBase, &company.USUCodigo, &company.ConvertTipoHe,
			&company.CPF, &company.DataEncerramento, &company.UltimaAtualizacaoAC, &company.FaltaAjustarNoAC, &company.AderiuESocial, &company.DataAdesaoESocial,
			&company.DataAdesaoESocialF2, &company.TpAmbESocial, &company.StatusEnvioApp, &company.NomeFantasia, &company.CNPJLicenciado, &company.FreemiumLastUpdate)
		if err != nil {
			log.Printf("Failed to scan company: %v", err)
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (r *CompanyRepository) GetByCodigo(codigo string) (*models.Empresa, error) {
	query := `SELECT id, codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				   CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				   Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				   DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				   NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update 
			FROM EMPRESA 
			WHERE codigo = $1`
	row := r.db.QueryRow(query, codigo)

	var company models.Empresa
	err := row.Scan(&company.ID, &company.Codigo, &company.Nome, &company.RazaoSocial, &company.CNPJBase, &company.USUCodigo, &company.ConvertTipoHe,
		&company.CPF, &company.DataEncerramento, &company.UltimaAtualizacaoAC, &company.FaltaAjustarNoAC, &company.AderiuESocial, &company.DataAdesaoESocial,
		&company.DataAdesaoESocialF2, &company.TpAmbESocial, &company.StatusEnvioApp, &company.NomeFantasia, &company.CNPJLicenciado, &company.FreemiumLastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Empresa não encontrada
		}
		log.Printf("Failed to fetch company by code: %v", err)
		return nil, err
	}

	return &company, nil
}

func (r *CompanyRepository) DeleteByCodigo(codigo string) error {
	query := `DELETE FROM EMPRESA WHERE Codigo = $1`
	_, err := r.db.Exec(query, codigo)
	if err != nil {
		log.Printf("Failed to delete company by code: %v", err)
		return err
	}

	return nil
}
