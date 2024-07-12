package main

import (
	"database/sql"
	"log"

	"github.com/tonnarruda/products_go/config"
	"github.com/tonnarruda/products_go/db"
	"github.com/tonnarruda/products_go/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega variáveis de ambiente do arquivo .env
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
	// Inicia o servidor HTTP na porta 8080
	log.Fatal(router.Run(":8080"))
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Middleware de logging e recuperação de pânico
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Rotas da API
	router.GET("/info", handlers.GetInfo(db))
	router.GET("/empresas", handlers.GetEmpresa(db))
	router.GET("/empresa/:codigo", handlers.GetEmpresaByCodigo(db))
	router.POST("/empresa", handlers.PostEmpresa(db))

	return router
}
