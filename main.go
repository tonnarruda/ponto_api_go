package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/config"
	"github.com/tonnarruda/ponto_api_go/db"
	"github.com/tonnarruda/ponto_api_go/handlers"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/services"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}
	defer database.Close()

	router := setupRouter(database)

	log.Fatal(router.Run(":8080"))
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	companyRepo := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	companyHandler := handlers.NewCompanyHandler(companyService)

	// Rotas da API
	router.POST("/empresa", companyHandler.CreateCompanyHandler)
	router.GET("/empresas", companyHandler.GetAllCompaniesHandler)
	router.GET("/empresa/:codigo", companyHandler.GetCompanyByCodigoHandler)
	router.DELETE("/empresa/:codigo", companyHandler.DeleteCompanyByCodigoHandler)

	return router
}
