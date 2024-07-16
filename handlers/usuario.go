package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/services"
	"github.com/tonnarruda/ponto_api_go/structs"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var newUser structs.Usuario
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Codigo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The field 'codigo' is required. Please ensure it is provided."})
		return
	}

	if err := h.userService.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create User",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByCodigoHandler(c *gin.Context) {
	code := c.Param("codigo")
	user, err := h.userService.GetUserByCodigo(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
