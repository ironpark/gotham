package main

import (
	"fmt"
	//	"github.com/IronPark/gotham/module/db"
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
	obj := struct {
		Nav         int
		User        string
		ProjectName string
	}{
		Nav:         PROJECT,
		User:        c.URLParams["user"],
		ProjectName: c.URLParams["name"],
	}
	err := template.Render(w, "project-info.html", obj)
	if err != nil {
		fmt.Println(err)
	}
}
