package api

import (
	"conexionMysql/conexion"
	"conexionMysql/modelo"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB = conexion.CrearConexion()

func GetCanciones(c *gin.Context) {
	ctx := context.Background()
	c.IndentedJSON(http.StatusOK, conexion.QueryMusic(ctx, db, 5))
}

func GetCancionByID(c *gin.Context) {

}

func PostCancion(c *gin.Context) {
	var cancion modelo.Cancion
	err := json.NewDecoder(c.Request.Body).Decode(&cancion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el cuerpo de la solicitud"})
		return
	}

	ctx := context.Background()
	err = conexion.AddMusica(ctx, db, cancion.ID, cancion.Name, cancion.Album, cancion.Artist, cancion.Genre, cancion.Year, cancion.Url_image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar la canción"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Canción agregada exitosamente"})
}

func PutCancion(c *gin.Context) {

}

func DeleteCancion(c *gin.Context) {

}
