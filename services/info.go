package services

import (
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

type InfoService struct {
	infoRepository repositories.InfoRepositoryInterface
}

func NewInfoService(infoRepo repositories.InfoRepositoryInterface) *InfoService {
	return &InfoService{infoRepository: infoRepo}
}

func (s *InfoService) GetAllInfo() ([]structs.Info, error) {
	return s.infoRepository.GetAll()
}
