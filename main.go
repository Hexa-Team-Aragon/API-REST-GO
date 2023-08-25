package main

import (
	"conexionMysql/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/canciones", api.GetCanciones)
}
