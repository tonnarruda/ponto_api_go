package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/handlers"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/services"
)

func SetupCompanyRoutes(router *gin.Engine, db *sql.DB) {
	companyRepo := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	companyHandler := handlers.NewCompanyHandler(companyService)

	companyRouter := router.Group("/empresa")
	{
		companyRouter.POST("", companyHandler.CreateCompanyHandler)
		companyRouter.PUT("", companyHandler.UpdateCompany)
		companyRouter.GET("", companyHandler.GetAllCompaniesHandler)
		companyRouter.GET("/:codigo", companyHandler.GetCompanyByCodigoHandler)
		companyRouter.DELETE("", companyHandler.DeleteCompanyByCodigoHandler)
		companyRouter.DELETE("/all", companyHandler.DeleteAllHandler)
	}
}
