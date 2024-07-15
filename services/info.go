package services

import (
	"github.com/tonnarruda/ponto_api_go/models"
	"github.com/tonnarruda/ponto_api_go/repositories"
)

type InfoService struct {
	infoRepository *repositories.InfoRepository
}

func NewInfoService(infoRepo *repositories.InfoRepository) *InfoService {
	return &InfoService{infoRepository: infoRepo}
}

func (s *InfoService) GetAllInfo() ([]models.Info, error) {
	return s.infoRepository.GetAll()
}
