package services

import (
	"github.com/tonnarruda/products_go/models"
	"github.com/tonnarruda/products_go/repositories"
)

type CompanyService struct {
	companyRepository *repositories.CompanyRepository
}

func NewCompanyService(companyRepo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepo}
}

func (s *CompanyService) CreateCompany(company *models.Empresa) error {
	err := s.companyRepository.Create(company)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) GetAllCompanies() ([]models.Empresa, error) {
	return s.companyRepository.GetAll()
}

func (s *CompanyService) GetCompanyByCodigo(codigo string) (*models.Empresa, error) {
	return s.companyRepository.GetByCodigo(codigo)
}

func (s *CompanyService) DeleteCompanyByCodigo(codigo string) error {
	return s.companyRepository.DeleteByCodigo(codigo)
}
