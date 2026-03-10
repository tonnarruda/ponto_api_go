package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tonnarruda/ponto_api_go/services"
)

func main() {
	router := gin.Default()

	router.GET("/cnh/pode-tirar", func(c *gin.Context) {
		idadeStr := c.Query("idade")
		if idadeStr == "" {
			c.JSON(400, gin.H{"error": "idade é obrigatória"})
			return
		}
		idade, err := strconv.Atoi(idadeStr)
		if err != nil || idade < 0 {
			c.JSON(400, gin.H{"error": "idade inválida"})
			return
		}
		c.JSON(200, gin.H{"pode_tirar": services.PodeTirarCNH(idade)})
	})

	log.Fatal(router.Run(":8080"))
}
