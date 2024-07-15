package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/services"
)

type InfoHandler struct {
	infoService *services.InfoService
}

func NewInfoHandler(infoService *services.InfoService) *InfoHandler {
	return &InfoHandler{infoService: infoService}
}

func (h *InfoHandler) GetAllInfoHandler(c *gin.Context) {
	infos, err := h.infoService.GetAllInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch companies"})
		return
	}

	c.JSON(http.StatusOK, infos)
}
