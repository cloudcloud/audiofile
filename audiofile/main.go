// Package main provides the binary for audiofile orchestrating and
// running the server.
package main

import (
	"log"

	"github.com/cloudcloud/audiofile/data"
	"github.com/gin-gonic/gin"
)

func main() {
	d, err := data.Open()
	if err != nil {
		log.Fatalf("Unable to setup database: %s", err)
	}

	r := addApi(addMiddleware(gin.Default(), d))
	r.GET("/health", health)
	r.GET("/", root)

	r.Run(":8008")
}

func addApi(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	api.GET("/albums", albums)
	api.GET("/artist/:artist", artist)
	api.GET("/artists", getArtists)
	api.GET("/settings/directories", getDirectories)

	api.POST("/settings/directory", storeDirectory)

	api.DELETE("/settings/directory", deleteDirectory)

	return r
}
