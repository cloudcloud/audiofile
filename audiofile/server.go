package main

import (
	"net/http"

	"github.com/cloudcloud/audiofile/data"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func addMiddleware(r *gin.Engine, d *data.Data) *gin.Engine {
	r.StaticFS("/js",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "js/",
		},
	)

	r.StaticFS("/css",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "css/",
		},
	)

	r.Use(
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:8008", "http://localhost:8080"},
			AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEAD", "DELETE"},
			AllowHeaders: []string{"Origin", "X-Client", "Content-Type"},
		}),
		pushData(d),
	)

	// TODO: setup the logger with gelf style

	return r
}

func pushData(d *data.Data) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", d)
		c.Next()
	}
}

func health(c *gin.Context) {
	d := c.MustGet("db").(*data.Data)
	if d != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "no-db",
		})
	}
}
