package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const (
	DATABASE_NAME = "data.db"
)

var (
	db *bolt.DB
)

func init() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open(DATABASE_NAME, 0600, nil)
	fmt.Println("init model")
	if err != nil {
		log.Fatal(err)
	}
}

func encode(st interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, st)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buf.Bytes()
}

func decode(b []byte, st interface{}) {
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, st)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
