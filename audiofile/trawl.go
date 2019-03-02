package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cloudcloud/audiofile"
	"github.com/cloudcloud/audiofile/data"
	id3 "github.com/cloudcloud/go-id3"
)

// Walk provides some structure to enable a simplification of walking
// files recursively whilst processing metadata appropriately.
type Walk struct {
	Errors []string
	Files  []audiofile.File

	db *data.Data
	fc chan audiofile.File
	mc chan bool
}

// Trawl accepts a Directory and processes the location recursively,
// looking for appropriate files to ingest.
func Trawl(dir audiofile.Directory, w *Walk) *Walk {
	w.fc = make(chan audiofile.File, 30)
	w.mc = make(chan bool, 1)

	go w.Push(dir.Directory)

	err := filepath.Walk(dir.Directory, w.Read)
	if err != nil {
		w.Errors = append(w.Errors, err.Error())
	}
	close(w.fc)

	<-w.mc
	close(w.mc)

	// Run through the files now to properly store them

	return w
}

// Push will loop and wait for incoming File items on the internal
// Walk channel.
func (w *Walk) Push(d string) {
	for {
		f, more := <-w.fc
		if !more {
			w.mc <- true
			return
		} else {
			w.Files = append(w.Files, f)
		}
	}
}

// Read will pick up a found file from the walk and push any that
// are appropriate into the internal Walk channel.
func (w *Walk) Read(p string, i os.FileInfo, err error) error {
	if err != nil || i.IsDir() {
		return nil
	}

	if filepath.Ext(i.Name()) == ".mp3" {
		w.fc <- audiofile.File{Filename: p}
	}

	return nil
}

// RetrieveID3 will take the provided filename and use go-id3
// to pull found ID3 tag information.
func RetrieveID3(f string) (id3.File, error) {
	file, err := os.Open(f)
	if err != nil {
		return id3.File{}, err
	}

	i := id3.File{Filename: f}
	i.Process(file)

	return i, nil
}

// MakeURL will convert the provided string into something a little
// better for routing and matching from URLs.
func MakeURL(s string) string {
	r := regexp.MustCompile("[^a-z0-9-_]+")
	rd := regexp.MustCompile("-\\-+")

	return rd.ReplaceAllString(
		r.ReplaceAllString(
			strings.ToLower(s),
			"-",
		),
		"-",
	)
}
