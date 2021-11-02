package main

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/api"
)

func main() {
	router := gin.Default()

	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumById)
	router.POST("/albums", api.PostAlbums)

	router.GET("/courses", api.GetCourses)

	router.Run("localhost:7070")
}
