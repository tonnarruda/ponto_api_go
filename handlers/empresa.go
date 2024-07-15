package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/models"
	"github.com/tonnarruda/ponto_api_go/services"
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

	if newCompany.Codigo == "" || newCompany.Nome == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The fields 'codigo' and 'nome' are required. Please ensure both are provided."})
		return
	}

	if err := h.companyService.CreateCompany(&newCompany); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create company",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Company created successfully"})
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	code := c.Query("codigo")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	var updatedCompany models.Empresa
	if err := c.ShouldBindJSON(&updatedCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.companyService.UpdateCompany(code, &updatedCompany); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update company",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully."})
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
	code := c.Param("codigo")
	company, err := h.companyService.GetCompanyByCodigo(code)
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
	code := c.Query("codigo")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The 'codigo' parameter is required"})
		return
	}

	var deletedCompany models.Empresa
	err := h.companyService.DeleteCompanyByCodigo(code, &deletedCompany)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
