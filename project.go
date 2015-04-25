package main

import (
	"fmt"
	//	"github.com/IronPark/gotham/module/db"
	"github.com/IronPark/gotham/module/git"
	"github.com/IronPark/gotham/module/template"
	"github.com/IronPark/gotham/module/util"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func init() {
	goji.Get("/project", projectListHandler)
	goji.Get("/project/:user/:name", projectInfoHandler)
}

func projectListHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	err := template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	if err != nil {
		fmt.Println(err)
	}
	err = template.Render(w, "project-list.html", templateObj(PROJECT))
	if err != nil {
		fmt.Println(err)
	}
}

func projectInfoHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	user := c.URLParams["user"]
	name := c.URLParams["name"]
	template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	Tree, _ := git.NewRepo(user, name).FileList()
	obj := struct {
		Nav         int
		User        string
		ProjectName string
		FileTree    []string
	}{
		Nav:         PROJECT,
		User:        user,
		ProjectName: name,
		FileTree:    Tree,
	}
	err := template.Render(w, "project-info.html", obj)
	if err != nil {
		fmt.Println(err)
	}
}
