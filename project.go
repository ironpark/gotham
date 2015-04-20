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
	goji.Get("/project", projectHandler)
}

func projectHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	err := template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	if err != nil {
		fmt.Println(err)
	}
	err = template.Render(w, "project.html", templateObj(PROJECT))
	if err != nil {
		fmt.Println(err)
	}
}
