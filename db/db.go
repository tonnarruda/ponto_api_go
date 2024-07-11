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
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS INFO (
			VERSAOBD INTEGER,
			STATUSBD VARCHAR(10),
			SISTEMA VARCHAR(20),
			VERSAOBDBETA INTEGER DEFAULT 0 NOT NULL,
			ATUALIZANDO INTEGER DEFAULT 0 NOT NULL CHECK (ATUALIZANDO IN (0, 1)),
			FORTES INTEGER DEFAULT 0 NOT NULL CHECK (FORTES IN (0, 1)),
			CONVERTEPONTO3 INTEGER DEFAULT 1 CHECK (CONVERTEPONTO3 IN (0, 1))
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS USU (
			Codigo VARCHAR(20) DEFAULT '' NOT NULL,
			Senha INT NOT NULL,
			UltimoAcesso TIMESTAMP,
			Bloqueado INTEGER DEFAULT 0 NOT NULL CHECK (Bloqueado IN (0, 1)),
			UserRegistrationDate TIMESTAMP DEFAULT '0001-01-01 00:00:00' NOT NULL,
			LimiteEpgData TIMESTAMP DEFAULT '0001-01-01 00:00:00' NOT NULL,
			PRIMARY KEY (Codigo)
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS MUSU (
			Id INTEGER NOT NULL,
			USU_Codigo VARCHAR(20) DEFAULT '' NOT NULL,
			PRIMARY KEY (Id, USU_Codigo)
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS EMP (
    Codigo VARCHAR(4) NOT NULL,
    Nome VARCHAR(15) NOT NULL,
    RazaoSocial VARCHAR(60),
    CNPJBase VARCHAR(8),
    USU_CODIGO VARCHAR(20),
    CONVERTETIPOHE INTEGER DEFAULT 1 CHECK (CONVERTETIPOHE IN (0, 1)),
    CPF VARCHAR(11),
    DTENCERRAMENTO TIMESTAMP,
    Ultima_Atualizacao_AC TIMESTAMP,
    Falta_Ajustar_No_AC INTEGER DEFAULT 0 NOT NULL CHECK (Falta_Ajustar_No_AC IN (0, 1)),
    ADERIU_ESOCIAL INTEGER DEFAULT 0 NOT NULL CHECK (ADERIU_ESOCIAL IN (0, 1)),
    DATA_ADESAO_ESOCIAL TIMESTAMP,
    DATA_ADESAO_ESOCIAL_F2 TIMESTAMP,
    TP_AMB_ESOCIAL INTEGER,
    STATUSENVIOAPP INT DEFAULT 0 NOT NULL,
    NMFANTASIA VARCHAR(40),
    CNPJLICENCIADO VARCHAR(14),
    Freemium_Last_Update VARCHAR(15) DEFAULT '0' NOT NULL,
    PRIMARY KEY (Codigo),
    UNIQUE (Nome),
    CONSTRAINT FK_EMP_USU FOREIGN KEY (USU_CODIGO) REFERENCES USU (Codigo)
);

	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS CFE (
			EMP_Codigo VARCHAR(4) NOT NULL,
			Codigo VARCHAR(50) NOT NULL,
			Valor VARCHAR(255) NOT NULL,
			PRIMARY KEY (EMP_Codigo, Codigo),
			CONSTRAINT FK_CFE_EMP FOREIGN KEY (EMP_Codigo) REFERENCES EMP (Codigo)
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
