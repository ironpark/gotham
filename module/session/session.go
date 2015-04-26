package session

import (
	//"fmt"
	"github.com/IronPark/gotham/module/util"
	"github.com/boltdb/bolt"
	"log"
	//"time"
)

const (
	DATABASE_NAME = "session.db"
)

var (
	db *bolt.DB
)

func init() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = bolt.Open(DATABASE_NAME, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("session"))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewSession(id, userAgent string) string {
	now := util.NowTimeStamp()
	token := util.HashSha512(now + id)
	//TO-DO User Agent CK
	db.Update(func(tx *bolt.Tx) error {
		session := tx.Bucket([]byte("session"))
		time := tx.Bucket([]byte("time"))

		session.Put([]byte(id), token)
		time.Put([]byte(id), []byte(now))
		return nil
	})
	return string(token)
}

func SessionCheck(id, token string) bool {
	result := false
	db.View(func(tx *bolt.Tx) error {
		session_bk := tx.Bucket([]byte("session"))
		time_bk := tx.Bucket([]byte("time"))

		//if token != string(session.Get([]byte(id)){

		//	}
		serverToken := string(session_bk.Get([]byte(id)))
		serverTokenTime := string(time_bk.Get([]byte(id)))
		if serverToken != token {
			return nil
		}
		//TO-DO Time Limit
		if serverTokenTime == "" {
			return nil
		}
		result = true
		return nil
	})
	return result
}
