package util

import (
	"crypto/sha512"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
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

func HashSha512(str string) []byte {
	h512 := sha512.New()
	io.WriteString(h512, str)
	return h512.Sum([]byte("gotham!"))
}

func NowTimeStamp() string {
	return time.Now().Format(time.RFC850)
}
