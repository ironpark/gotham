package db

import (
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

	if err != nil {
		log.Fatal(err)
	}
}

func DatabaseInitialization() {
	db.Update(func(tx *bolt.Tx) error {
		//Users
		users, _ := tx.CreateBucketIfNotExists([]byte("user"))
		users.Put([]byte("admin"), hash("admin")) //Default Admin User

		//Project
		tx.CreateBucketIfNotExists([]byte("project_list"))
		//
		return nil
	})
}

func NewProject(user, name, desc string) {
	db.Update(func(tx *bolt.Tx) error {
		projectList := tx.Bucket([]byte("project_list"))
		project, _ := projectList.CreateBucket([]byte(user + "/" + name))
		project.Put([]byte("name"), []byte(name))
		project.Put([]byte("desc"), []byte(desc))
		project.CreateBucket([]byte("commit"))
		project.CreateBucket([]byte("comments"))
		project.CreateBucket([]byte("files"))
		project.CreateBucket([]byte("task"))
		return nil
	})
}

func Project() []ProjectST {
	fmt.Println("Projects")
	projects := make([]ProjectST, 0)
	db.View(func(tx *bolt.Tx) error {
		projectList := tx.Bucket([]byte("project_list"))
		projectList.ForEach(func(k, v []byte) error {
			projectBK := projectList.Bucket(k)
			project := ProjectST{}
			project.Name = string(projectBK.Get([]byte("name")))
			project.Desc = string(projectBK.Get([]byte("desc")))
			project.Path = string(k)
			projects = append(projects, project)

			// string(projectList.Bucket(k).Get([]byte("name"))))
			return nil
		})
		return nil
	})
	return projects
}
func Close() {
	db.Close()
}
