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
	// Carrega variáveis de ambiente do arquivo .env
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup do banco de dados
	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}
	defer database.Close()

	// Inicializa o roteador com Gin
	router := setupRouter(database)

	// Inicia o servidor HTTP na porta 8080
	log.Fatal(router.Run(":8080"))
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Middleware de logging e recuperação de pânico
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Inicializar repositório e serviço do usuário
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
