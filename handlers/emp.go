package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Driver PostgreSQL
)

type Emp struct {
	Codigo              string       `json:"codigo"`
	Nome                string       `json:"nome"`
	RazaoSocial         string       `json:"razao_social"`
	CNPJBase            string       `json:"cnpj_base"`
	USUCodigo           *string      `json:"usu_codigo"` // Usando *string para suportar NULL
	ConvertTipoHe       int          `json:"convert_tipo_he"`
	CPF                 string       `json:"cpf"`
	DataEncerramento    sql.NullTime `json:"dt_encerramento"`       // Usando sql.NullTime para suportar NULL
	UltimaAtualizacaoAC sql.NullTime `json:"ultima_atualizacao_ac"` // Usando sql.NullTime para suportar NULL
	FaltaAjustarNoAC    int          `json:"falta_ajustar_no_ac"`
	AderiuESocial       int          `json:"aderiu_esocial"`
	DataAdesaoESocial   sql.NullTime `json:"data_adesao_esocial"` // Usando sql.NullTime para suportar NULL
	DataAdesaoESocialF2 sql.NullTime `json:"data_adesao_esocial_f2"`
	TpAmbESocial        int          `json:"tp_amb_esocial"`
	StatusEnvioApp      int          `json:"status_envio_app"`
	NomeFantasia        string       `json:"nmfantasia"`
	CNPJLicenciado      string       `json:"cnpj_licenciado"`
	FreemiumLastUpdate  string       `json:"freemium_last_update"`
}

func GetEmp(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				   CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				   Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				   DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				   NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update
			FROM EMP
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database", "details": err.Error()})
			return
		}
		defer rows.Close()

		var emps []Emp
		for rows.Next() {
			var emp Emp
			err := rows.Scan(
				&emp.Codigo, &emp.Nome, &emp.RazaoSocial, &emp.CNPJBase,
				&emp.USUCodigo, &emp.ConvertTipoHe, &emp.CPF, &emp.DataEncerramento,
				&emp.UltimaAtualizacaoAC, &emp.FaltaAjustarNoAC, &emp.AderiuESocial,
				&emp.DataAdesaoESocial, &emp.DataAdesaoESocialF2, &emp.TpAmbESocial,
				&emp.StatusEnvioApp, &emp.NomeFantasia, &emp.CNPJLicenciado, &emp.FreemiumLastUpdate,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan EMP", "details": err.Error()})
				return
			}

			// Verificar se DataAdesaoESocial é nulo
			if emp.DataAdesaoESocialF2.Valid {
				emp.DataAdesaoESocialF2.Time = emp.DataAdesaoESocialF2.Time.UTC() // Ajuste opcional de timezone
			}

			// Verificar se DataAdesaoESocial é nulo
			if emp.DataAdesaoESocial.Valid {
				emp.DataAdesaoESocial.Time = emp.DataAdesaoESocial.Time.UTC() // Ajuste opcional de timezone
			}

			// Verificar se UltimaAtualizacaoAC é nulo
			if emp.UltimaAtualizacaoAC.Valid {
				emp.UltimaAtualizacaoAC.Time = emp.UltimaAtualizacaoAC.Time.UTC() // Ajuste opcional de timezone
			}

			// Verificar se DataEncerramento é nulo
			if emp.DataEncerramento.Valid {
				emp.DataEncerramento.Time = emp.DataEncerramento.Time.UTC() // Ajuste opcional de timezone
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

func CreateEmp(db *sql.DB) gin.HandlerFunc {
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
			INSERT INTO EMP (
				Codigo, Nome, RazaoSocial, CNPJBase, USU_CODIGO,
				CONVERTETIPOHE, CPF, DTENCERRAMENTO, Ultima_Atualizacao_AC,
				Falta_Ajustar_No_AC, ADERIU_ESOCIAL, DATA_ADESAO_ESOCIAL,
				DATA_ADESAO_ESOCIAL_F2, TP_AMB_ESOCIAL, STATUSENVIOAPP,
				NMFANTASIA, CNPJLICENCIADO, Freemium_Last_Update
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				$11, $12, $13, $14, $15, $16, $17, $18, $19
			)
		`)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
			return
		}
		defer stmt.Close()

		// Executando a inserção
		_, err = stmt.Exec(
			emp.Codigo, emp.Nome, emp.RazaoSocial, emp.CNPJBase,
			emp.USUCodigo, emp.ConvertTipoHe, emp.CPF, emp.DataEncerramento,
			emp.UltimaAtualizacaoAC, emp.FaltaAjustarNoAC, emp.AderiuESocial,
			emp.DataAdesaoESocial, emp.DataAdesaoESocialF2, emp.TpAmbESocial,
			emp.StatusEnvioApp, emp.NomeFantasia, emp.CNPJLicenciado, emp.FreemiumLastUpdate,
		)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute statement"})
			return
		}

		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Empresa criada com sucesso", "empresa": emp})
	}
}
