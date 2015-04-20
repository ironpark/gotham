package util

import (
	"log"
	"os"
	"path/filepath"
)

var (
	wkdir string
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	} else {
		wkdir = dir
		wkdir = filepath.ToSlash(wkdir)
	}
}

func WorkingDir() string {
	return wkdir
}
