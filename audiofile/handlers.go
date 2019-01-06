package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudcloud/audiofile"
	"github.com/cloudcloud/audiofile/data"
	"github.com/gin-gonic/gin"
)

type MetaFunc func(*gin.Context, *data.Data) (gin.H, []string)

func albums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items":  []gin.H{},
		"errors": []string{},
		"meta":   map[string]interface{}{},
	})
}

func artist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func artists(c *gin.Context) {
	// pull the data from the context
	// retrieve artists based on the query

	c.JSON(http.StatusOK, gin.H{
		"items": []gin.H{
			{"text": "Ministry", "href": "/artist/ministry", "albums": []gin.H{}, "status": "Active"},
			{"text": "T.O.O.H", "href": "/artist/t.o.o.h", "albums": []gin.H{}, "status": "Active"},
		},
		"errors": []string{},
		"meta":   map[string]interface{}{},
	})
}

func getDirectories(c *gin.Context) {
	b := time.Now()
	db := c.MustGet("db")

	d, err := db.(*data.Data).GetDirectories()
	e := time.Now().Sub(b)
	errs := []string{}
	if err != nil {
		errs = append(errs, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"items":  d,
		"errors": errs,
		"meta": map[string]interface{}{
			"time_taken": fmt.Sprintf("%v", e),
			"errors":     len(errs),
		},
	})
}

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func storeDirectory(c *gin.Context) {
	b := time.Now()
	db := c.MustGet("db")
	dir := audiofile.Directory{}
	errs := []string{}
	var err error

	if err := c.BindJSON(&dir); err != nil {
		errs = append(errs, err.Error())
	}
	if dir.ID == "" {
		dir, err = dir.GenerateID()
		if err != nil {
			errs = append(errs, err.Error())
		}

		dir.DateAdded = time.Now()
		dir.DateUpdated = time.Now()
	}

	d, err := db.(*data.Data).StoreDirectory(dir)
	e := time.Now().Sub(b)
	if err != nil {
		errs = append(errs, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"items":  []audiofile.Directory{d},
		"errors": errs,
		"meta": map[string]interface{}{
			"time_taken": fmt.Sprintf("%v", e),
			"errors":     len(errs),
		},
	})
}

func deleteDirectory(c *gin.Context) {
	withMeta(c, func(c *gin.Context, db *data.Data) (gin.H, []string) {
		errs := []string{}
		dir := audiofile.Directory{}

		if err := c.BindJSON(&dir); err != nil {
			errs = append(errs, err.Error())
		}

		err := db.DeleteDirectory(dir)
		if err != nil {
			errs = append(errs, err.Error())

			return gin.H{"status": "failure"}, errs
		}

		return gin.H{"status": "success"}, errs
	})
}

func withMeta(c *gin.Context, f MetaFunc) *gin.Context {
	b := time.Now()
	db := c.MustGet("db").(*data.Data)

	d, errs := f(c, db)
	e := time.Now().Sub(b)

	c.JSON(http.StatusOK, gin.H{
		"items":  []gin.H{d},
		"errors": errs,
		"meta": map[string]interface{}{
			"time_taken": fmt.Sprintf("%v", e),
			"errors":     len(errs),
		},
	})

	return c
}
