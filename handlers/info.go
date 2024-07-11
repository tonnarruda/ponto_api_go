package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Info struct {
	VersaoBD       int    `json:"versao_bd"`
	StatusBD       string `json:"status_bd"`
	Sistema        string `json:"sistema"`
	VersaoBDBeta   int    `json:"versao_bd_beta"`
	Atualizando    int    `json:"atualizando"`
	Fortes         int    `json:"fortes"`
	ConvertePonto3 int    `json:"converte_ponto3"`
}

func GetInfo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT VERSAOBD, STATUSBD, SISTEMA, VERSAOBDBETA, ATUALIZANDO, FORTES, CONVERTEPONTO3
			FROM INFO
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}
		defer rows.Close()

		var info Info
		for rows.Next() {
			if err := rows.Scan(&info.VersaoBD, &info.StatusBD, &info.Sistema, &info.VersaoBDBeta, &info.Atualizando, &info.Fortes, &info.ConvertePonto3); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan INFO"})
				return
			}
		}

		if err = rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during iteration"})
			return
		}

		c.JSON(http.StatusOK, info)
	}
}
