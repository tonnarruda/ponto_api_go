package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/config"
	"github.com/tonnarruda/ponto_api_go/db"
	"github.com/tonnarruda/ponto_api_go/routes"
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

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.SetupCompanyRoutes(router, database)

	log.Fatal(router.Run(":8080"))
}
