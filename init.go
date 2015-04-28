package main

import (
	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	DATABASE_NAME = "data.db"
)

//DatabaseInitialization
func init() {
	db, err := bolt.Open(DATABASE_NAME, 0600, nil)
	defer db.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		//Users
		users, _ := tx.CreateBucketIfNotExists([]byte("user"))
		users.Put([]byte("admin"), hash) //Default Admin User

		//Project
		tx.CreateBucketIfNotExists([]byte("project_list"))
		return nil
	})
}
