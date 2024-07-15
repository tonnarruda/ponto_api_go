package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/handlers"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/services"
)

func SetupInfoRoutes(router *gin.Engine, db *sql.DB) {
	infoRepo := repositories.NewInfoRepository(db)
	infoService := services.NewInfoService(infoRepo)
	infoHandler := handlers.NewInfoHandler(infoService)

	infoRouter := router.Group("/info")
	{
		infoRouter.GET("", infoHandler.GetAllInfoHandler)
	}
}
