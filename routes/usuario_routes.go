package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/handlers"
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/services"
)

func SetupUserRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRouter := router.Group("/usuario")
	{
		userRouter.POST("", userHandler.CreateUserHandler)
		userRouter.GET("", userHandler.GetAllUsersHandler)
		userRouter.GET("/:codigo", userHandler.GetUserByCodigoHandler)
	}
}
