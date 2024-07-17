package services

import (
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

type CompanyService struct {
	companyRepository *repositories.CompanyRepository
}

func NewCompanyService(companyRepo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepo}
}

func (s *CompanyService) CreateCompany(company *structs.Empresa) error {
	err := s.companyRepository.Create(company)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) GetAllCompanies() ([]structs.Empresa, error) {
	return s.companyRepository.GetAll()
}

func (s *CompanyService) GetCompanyByCodigo(codigo string) (*structs.Empresa, error) {
	return s.companyRepository.GetByCodigo(codigo)
}

func (s *CompanyService) DeleteCompanyByCodigo(codigo string) error {
	return s.companyRepository.DeleteByCodigo(codigo)
}

func (s *CompanyService) DeleteAllCompanies() error {
	return s.companyRepository.DeleteAll()
}

func (s *CompanyService) UpdateCompany(codigo string, company *structs.Empresa) error {
	return s.companyRepository.UpdateByCodigo(codigo, company)
}

func (s *CompanyService) GetLastCompanyCode() (string, error) {
	return s.companyRepository.GetLastCompanyCode()
}
