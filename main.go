package main

import (
	"conexionMysql/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/canciones", api.GetCanciones)
	router.POST("/canciones/add", api.PostCancion)
	router.DELETE("/canciones/delete/:id", api.DeleteCancion)
	router.Run("localhost:8888")
}
