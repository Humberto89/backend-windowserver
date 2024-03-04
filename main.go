package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un enrutador Gin por defecto
	r := gin.Default()

	// Define una ruta para la raíz que responde con "Hola, Mundo!"
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "¡Hola, Mundo! y esto lo escribio frederick el senior",
		})
	})

	// Inicia el servidor en el puerto 8080
	r.Run(":8080")
}
