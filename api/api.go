package api

import (
	"conexionMysql/conexion"
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB = conexion.CrearConexion()

func GetCanciones(c *gin.Context) {
	ctx := context.Background()
	c.IndentedJSON(http.StatusOK, conexion.QueryMusic(ctx, db, 4))
}

func GetCancionByID(c *gin.Context) {

}

func PostCancion(c *gin.Context) {

}

func PutCancion(c *gin.Context) {

}

func DeleteCancion(c *gin.Context) {

}
