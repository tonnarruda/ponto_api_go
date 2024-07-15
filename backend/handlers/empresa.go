package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/helper"
	"github.com/tonnarruda/ponto_api_go/services"
	"github.com/tonnarruda/ponto_api_go/structs"
)

type CompanyHandler struct {
	companyService *services.CompanyService
}

func NewCompanyHandler(companyService *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) CreateCompanyHandler(c *gin.Context) {
	var newCompany structs.Empresa
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newCompany.Nome == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The field 'nome' is required. Please ensure it is provided."})
		return
	}

	// Verifica se o código foi fornecido; se não, gera automaticamente
	if newCompany.Codigo == "" {
		lastCode, _ := h.companyService.GetLastCompanyCode()

		if lastCode == "" || strings.Contains(lastCode, "sql: no rows in result set") {
			newCompany.Codigo = "0001"
		} else {
			newCompany.Codigo = helper.GenerateNextCode(lastCode)
		}
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

	var updatedCompany structs.Empresa
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

	var deletedCompany structs.Empresa
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

func (h *CompanyHandler) DeleteAllHandler(c *gin.Context) {

	err := h.companyService.DeleteAllCompanies()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All Companies were deleted successfully"})
}
