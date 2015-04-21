package db

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"github.com/boltdb/bolt"
	"io"
)

func hash(str string) []byte {
	h512 := sha512.New()
	io.WriteString(h512, str)
	return h512.Sum([]byte("gotham!"))
}

func InsertUser(id, password string) {
	db.Update(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("user"))
		users.Put([]byte(id), hash(password))
		return nil
	})
}

func UserVerification(id, password string) bool {
	ok := false
	db.View(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("user"))
		if bytes.Equal(users.Get([]byte(id)), hash(password)) {
			fmt.Println("UserVerification")
			ok = true
		}
		return nil
	})
	return ok
}
