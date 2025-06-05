package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenstr/web-service-gin/albums-service/internal/repository/memory"
	model "github.com/stevenstr/web-service-gin/albums-service/pkg"
)

var Albums = memory.New()

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum *model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Json binding error"})
		return
	}

	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, alb := range Albums {
		if alb.ID == id {
			c.IndentedJSON(http.StatusOK, alb)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("album with id=%s is not found.", id)})
}

func NotFoundRoute(c *gin.Context) {
	c.String(http.StatusBadRequest, "не поддерживаемый маршрут")
}
