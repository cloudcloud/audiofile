package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudcloud/audiofile"
	"github.com/cloudcloud/audiofile/data"
	"github.com/gin-gonic/gin"
)

// MetaFunc is short-hand method of encapsulating internal API
// handlers that are wrapped within a response structure and timed.
//
// This function is used by individual handlers as a means to execute
// their required actions before generating a response. This response
// is then returned with the additional metadata.
type MetaFunc func(*gin.Context, *data.Data) (interface{}, []string)

func albums(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (interface{}, []string) {
		return db.GetAlbums(), []string{}
	})
}

func artist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getArtists(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (interface{}, []string) {
		return db.GetArtists(), []string{}
	})
}

func getDirectories(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (interface{}, []string) {
		errs := []string{}
		d, err := db.GetDirectories()
		if err != nil {
			errs = append(errs, err.Error())
		}

		return d, errs
	})
}

func root(c *gin.Context) {
	r := bytes.NewReader(MustAsset("index.html"))
	c.DataFromReader(
		http.StatusOK,
		int64(len(MustAssetString("index.html"))),
		"text/html",
		r,
		map[string]string{},
	)
}

func storeDirectory(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (interface{}, []string) {
		errs := []string{}
		dir := audiofile.Directory{}

		if err := c.BindJSON(&dir); err != nil {
			errs = append(errs, err.Error())
		}

		d, err := dir.MaybeFirstTime()
		if err != nil {
			errs = append(errs, err.Error())
		}

		res, err := db.StoreDirectory(d)
		if err != nil {
			errs = append(errs, err.Error())
		}

		return []gin.H{gin.H{"directory": res}}, errs
	})
}

func triggerTrawl(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (interface{}, []string) {
		dirs, _ := db.GetDirectories()
		for _, x := range dirs {
			w := &Walk{db: db, Errors: []string{}, Files: []audiofile.File{}, mc: make(chan bool)}
			go Trawl(x, w)
		}

		return []gin.H{gin.H{"trawl": "triggered"}}, []string{}
	})
}

func deleteDirectory(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (interface{}, []string) {
		errs := []string{}
		dir := audiofile.Directory{}

		if err := c.BindJSON(&dir); err != nil {
			errs = append(errs, err.Error())
		}

		err := db.DeleteDirectory(dir)
		if err != nil {
			errs = append(errs, err.Error())

			return []gin.H{gin.H{"status": "failure"}}, errs
		}

		return []gin.H{gin.H{"status": "success"}}, errs
	})
}

func withMeta(c *gin.Context, f MetaFunc) *gin.Context {
	b := time.Now()
	db := c.MustGet("db").(*data.Data)

	d, errs := f(c, db)
	e := time.Now().Sub(b)

	c.JSON(http.StatusOK, gin.H{
		"items":  d,
		"errors": errs,
		"meta": map[string]interface{}{
			"time_taken": fmt.Sprintf("%v", e),
			"errors":     len(errs),
		},
	})

	return c
}
