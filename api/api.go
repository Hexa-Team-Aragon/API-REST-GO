package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cancion struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Album     string `json:"album"`
	Artist    string `json:"Artist"`
	Genre     string `json:"Genre"`
	Year      int64  `json:"Year"`
	Url_image string `json:"url_image"`
}

func GetCanciones(c *gin.Context) {
	c.IndentedJSON(http.StatusOK)
}

func GetCancionByID(c *gin.Context) {

}

func PostCancion(c *gin.Context) {

}

func PutCancion(c *gin.Context) {

}

func DeleteCancion(c *gin.Context) {

}
