package repositories

import (
	"database/sql"
	"errors"
	"log"

	"github.com/tonnarruda/ponto_api_go/structs"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(company *structs.Empresa) error {
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
		log.Printf("Failed to insert company into database: %v", err)
		return err
	}
	return nil
}

func (r *CompanyRepository) UpdateByCodigo(codigo string, company *structs.Empresa) error {
	if codigo == "" {
		return errors.New("codigo is required")
	}
	query := `
		UPDATE EMPRESA SET 
			Nome = $1, RazaoSocial = $2, CNPJBase = $3, USU_CODIGO = $4,
			CONVERTETIPOHE = $5, CPF = $6, DTENCERRAMENTO = $7, Ultima_Atualizacao_AC = $8,
			Falta_Ajustar_No_AC = $9, ADERIU_ESOCIAL = $10, DATA_ADESAO_ESOCIAL = $11,
			DATA_ADESAO_ESOCIAL_F2 = $12, TP_AMB_ESOCIAL = $13, STATUSENVIOAPP = $14,
			NMFANTASIA = $15, CNPJLICENCIADO = $16, Freemium_Last_Update = $17
		WHERE Codigo = $18
	`
	_, err := r.db.Exec(query, company.Nome, company.RazaoSocial, company.CNPJBase, company.USUCodigo, company.ConvertTipoHe,
		company.CPF, company.DataEncerramento, company.UltimaAtualizacaoAC, company.FaltaAjustarNoAC, company.AderiuESocial, company.DataAdesaoESocial,
		company.DataAdesaoESocialF2, company.TpAmbESocial, company.StatusEnvioApp, company.NomeFantasia, company.CNPJLicenciado, company.FreemiumLastUpdate, codigo)
	if err != nil {
		log.Printf("Failed to update company in database: %v", err)
		return err
	}
	return nil
}

func (r *CompanyRepository) GetAll() ([]structs.Empresa, error) {
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

	var companies []structs.Empresa
	for rows.Next() {
		var company structs.Empresa
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

func (r *CompanyRepository) GetByCodigo(codigo string) (*structs.Empresa, error) {
	query := `SELECT id, codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				   CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				   Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				   DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				   NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update 
			FROM EMPRESA 
			WHERE codigo = $1`
	row := r.db.QueryRow(query, codigo)

	var company structs.Empresa
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

func (r *CompanyRepository) DeleteByCodigo(codigo string, company *structs.Empresa) error {
	query := `DELETE FROM EMPRESA WHERE Codigo = $1`
	_, err := r.db.Exec(query, codigo)
	if err != nil {
		log.Printf("Failed to delete company by code: %v", err)
		return err
	}

	return nil
}

func (r *CompanyRepository) DeleteAll() error {
	query := `DELETE FROM EMPRESA`
	_, err := r.db.Exec(query)
	if err != nil {
		log.Println("Failed to delete company", err)
		return err
	}

	return nil
}

func (r *CompanyRepository) GetLastCompanyCode() (string, error) {
	var lastCode string
	query := `SELECT codigo FROM empresa ORDER BY codigo DESC LIMIT 1`
	err := r.db.QueryRow(query).Scan(&lastCode)
	if err != nil {
		return "", err
	}
	return lastCode, nil
}
