package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "¡Hola, Mundo! y esto lo escribieron Frederick el Senior",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
