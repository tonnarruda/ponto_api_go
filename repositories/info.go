package repositories

import (
	"database/sql"
	"log"

	"github.com/tonnarruda/ponto_api_go/models"
)

type InfoRepository struct {
	db *sql.DB
}

func NewInfoRepository(db *sql.DB) *InfoRepository {
	return &InfoRepository{db: db}
}

func (r *InfoRepository) GetAll() ([]models.Info, error) {
	query := `SELECT versaobd, statusbd, sistema, versaobdbeta, atualizando, fortes, converteponto3
			FROM INFO`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to fetch companies from database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var infos []models.Info
	for rows.Next() {
		var info models.Info
		err := rows.Scan(&info.VersaoBD, &info.StatusBD, &info.Sistema,
			&info.VersaoBDBeta, &info.Atualizando, &info.Fortes, &info.ConvertePonto3)
		if err != nil {
			log.Printf("Failed to scan company: %v", err)
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}
