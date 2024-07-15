package services

import (
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

type InfoService struct {
	infoRepository *repositories.InfoRepository
}

func NewInfoService(infoRepo *repositories.InfoRepository) *InfoService {
	return &InfoService{infoRepository: infoRepo}
}

func (s *InfoService) GetAllInfo() ([]structs.Info, error) {
	return s.infoRepository.GetAll()
}
