package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/products_go/models"
	"github.com/tonnarruda/products_go/services"
)

type CompanyHandler struct {
	companyService *services.CompanyService
}

func NewCompanyHandler(companyService *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) CreateCompanyHandler(c *gin.Context) {
	var newCompany models.Empresa
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.companyService.CreateCompany(&newCompany); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Company created successfully"})
}

func (h *CompanyHandler) GetAllCompaniesHandler(c *gin.Context) {
	companies, err := h.companyService.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) GetCompanyByCodigoHandler(c *gin.Context) {
	codigo := c.Param("codigo")
	company, err := h.companyService.GetCompanyByCodigo(codigo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch company"})
		return
	}
	if company == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) DeleteCompanyByCodigoHandler(c *gin.Context) {

	codigo := c.Param("codigo")

	err := h.companyService.DeleteCompanyByCodigo(codigo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
