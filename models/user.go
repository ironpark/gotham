package models

import (
	"bytes"
	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Username string
	Password []byte
}

func (user *User) HashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//		glog.Fatalf("Couldn't hash password: %v", err)
		panic(err)
	}
	user.Password = hash
}

func (user *User) EqPassword(password string) bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//		glog.Fatalf("Couldn't hash password: %v", err)
		panic(err)
	}
	return bytes.Equal(user.Password, hash)
}

func GetUserByEmail(email string) (user *User) {
	user = &User{}
	db.View(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("user"))
		userdata := users.Get([]byte(email))
		decode(userdata, user)
		return nil
	})
	return
}
