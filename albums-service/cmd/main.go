package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stevenstr/web-service-gin/albums-service/internal/handler"
)

func main() {
	log.Println("Starting the albums service")
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbum)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.NoRoute(handler.NotFoundRoute)

	log.Println("Starting service at localhost:3000..")
	if err := router.Run("localhost:3000"); err != nil {
		log.Fatal(err)
	}
}
