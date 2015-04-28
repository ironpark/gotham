package models

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type Project struct {
	Name string
	Desc string
	Path string
	Meta string
}

func NewProject(user, name, desc string) {
	fmt.Println("NewProject")
	db.Update(func(tx *bolt.Tx) error {
		projectList := tx.Bucket([]byte("project_list"))
		project, err := projectList.CreateBucket([]byte(user + "/" + name))
		project.Put([]byte("name"), []byte(name))
		project.Put([]byte("desc"), []byte(desc))
		return err
	})
	fmt.Println("NewProject end")
}

func Projects() []Project {
	projects := make([]Project, 0)
	db.View(func(tx *bolt.Tx) error {
		projectList := tx.Bucket([]byte("project_list"))
		projectList.ForEach(func(k, v []byte) error {
			projectBK := projectList.Bucket(k)
			project := Project{}
			project.Name = string(projectBK.Get([]byte("name")))
			project.Desc = string(projectBK.Get([]byte("desc")))
			project.Path = string(k)
			projects = append(projects, project)
			return nil
		})
		return nil
	})
	return projects
}
