package api

import (
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

func (s *Cancion) getCanciones(c *gin.Context) {

}

func (s *Cancion) getCancionByID(c *gin.Context) {

}

func (s *Cancion) postCancion(c *gin.Context) {

}

func (s *Cancion) putCancion(c *gin.Context) {

}

func (s *Cancion) deleteCancion(c *gin.Context) {

}
