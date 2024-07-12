package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Emp struct {
	ID                  string     `json:"id"`
	Codigo              string     `json:"codigo"`
	Nome                string     `json:"nome"`
	RazaoSocial         string     `json:"razao_social"`
	CNPJBase            string     `json:"cnpj_base"`
	USUCodigo           *string    `json:"usu_codigo,omitempty"`
	ConvertTipoHe       int        `json:"convert_tipo_he"`
	CPF                 string     `json:"cpf"`
	DataEncerramento    *time.Time `json:"dt_encerramento,omitempty"`
	UltimaAtualizacaoAC *time.Time `json:"ultima_atualizacao_ac,omitempty"`
	FaltaAjustarNoAC    int        `json:"falta_ajustar_no_ac"`
	AderiuESocial       int        `json:"aderiu_esocial"`
	DataAdesaoESocial   *time.Time `json:"data_adesao_esocial,omitempty"`
	DataAdesaoESocialF2 *time.Time `json:"data_adesao_esocial_f2,omitempty"`
	TpAmbESocial        int        `json:"tp_amb_esocial"`
	StatusEnvioApp      int        `json:"status_envio_app"`
	NomeFantasia        string     `json:"nmfantasia"`
	CNPJLicenciado      string     `json:"cnpj_licenciado"`
	FreemiumLastUpdate  string     `json:"freemium_last_update"`
}

func GetEmpresa(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT id, Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				   CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				   Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				   DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				   NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update
			FROM EMPRESA
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database", "details": err.Error()})
			return
		}
		defer rows.Close()

		var emps []Emp
		for rows.Next() {
			var emp Emp
			var dtEncerramento, ultimaAtualizacaoAC, dataAdesaoESocial, dataAdesaoESocialF2 sql.NullTime
			err := rows.Scan(
				&emp.ID, &emp.Codigo, &emp.Nome, &emp.RazaoSocial, &emp.CNPJBase,
				&emp.USUCodigo, &emp.ConvertTipoHe, &emp.CPF, &dtEncerramento,
				&ultimaAtualizacaoAC, &emp.FaltaAjustarNoAC, &emp.AderiuESocial,
				&dataAdesaoESocial, &dataAdesaoESocialF2, &emp.TpAmbESocial,
				&emp.StatusEnvioApp, &emp.NomeFantasia, &emp.CNPJLicenciado, &emp.FreemiumLastUpdate,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan Table EMPRESA", "details": err.Error()})
				return
			}

			// Verificar se DataEncerramento é nulo
			if dtEncerramento.Valid {
				emp.DataEncerramento = &dtEncerramento.Time
			} else {
				emp.DataEncerramento = nil
			}

			// Verificar se UltimaAtualizacaoAC é nulo
			if ultimaAtualizacaoAC.Valid {
				emp.UltimaAtualizacaoAC = &ultimaAtualizacaoAC.Time
			} else {
				emp.UltimaAtualizacaoAC = nil
			}

			// Verificar se DataAdesaoESocial é nulo
			if dataAdesaoESocial.Valid {
				emp.DataAdesaoESocial = &dataAdesaoESocial.Time
			} else {
				emp.DataAdesaoESocial = nil
			}

			// Verificar se DataAdesaoESocialF2 é nulo
			if dataAdesaoESocialF2.Valid {
				emp.DataAdesaoESocialF2 = &dataAdesaoESocialF2.Time
			} else {
				emp.DataAdesaoESocialF2 = nil
			}

			emps = append(emps, emp)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during iteration", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, emps)
	}
}

func GetEmpresaByCodigo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		codigo := c.Param("codigo")

		var emp Emp
		query := `
			SELECT id, codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				   CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				   Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				   DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				   NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update 
			FROM EMPRESA 
			WHERE codigo = $1
		`
		fmt.Printf("Executing query: %s with codigo: %s\n", query, codigo) // Imprimir consulta para depuração

		err := db.QueryRow(query, codigo).Scan(
			&emp.ID, &emp.Codigo, &emp.Nome, &emp.RazaoSocial, &emp.CNPJBase, &emp.USUCodigo,
			&emp.ConvertTipoHe, &emp.CPF, &emp.DataEncerramento, &emp.UltimaAtualizacaoAC,
			&emp.FaltaAjustarNoAC, &emp.AderiuESocial, &emp.DataAdesaoESocial,
			&emp.DataAdesaoESocialF2, &emp.TpAmbESocial, &emp.StatusEnvioApp,
			&emp.NomeFantasia, &emp.CNPJLicenciado, &emp.FreemiumLastUpdate,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database", "details": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, emp)
	}
}

func PostEmpresa(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp Emp
		if err := c.ShouldBindJSON(&emp); err != nil {
			fmt.Println("Erro ao fazer bind do JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}

		stmt, err := tx.Prepare(`
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
	`)
		if err != nil {
			tx.Rollback()
			fmt.Println("Failed to prepare statement:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
			return
		}
		defer stmt.Close()

		var id string
		err = stmt.QueryRow(
			emp.Codigo, emp.Nome, emp.RazaoSocial, emp.CNPJBase,
			emp.USUCodigo, emp.ConvertTipoHe, emp.CPF, emp.DataEncerramento,
			emp.UltimaAtualizacaoAC, emp.FaltaAjustarNoAC, emp.AderiuESocial,
			emp.DataAdesaoESocial, emp.DataAdesaoESocialF2, emp.TpAmbESocial,
			emp.StatusEnvioApp, emp.NomeFantasia, emp.CNPJLicenciado, emp.FreemiumLastUpdate,
		).Scan(&id)
		if err != nil {
			tx.Rollback()
			fmt.Println("Failed to execute statement:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute statement", "details": err.Error()})
			return
		}

		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
			return
		}

		emp.ID = id
		c.JSON(http.StatusCreated, gin.H{"message": "Empresa criada com sucesso", "empresa": emp})

	}
}
