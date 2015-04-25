package main

import (
	"github.com/IronPark/gotham/module/db"
	"github.com/IronPark/gotham/module/git"
	//"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func init() {
	goji.Get("/api/:what", apiGetHandler)
	goji.Post("/api/:what", apiPostHandler)
	goji.Put("/api/:what", apiPutHandler)
	goji.Delete("/api/:what", apiDeleteHandler)
}

func apiGetHandler(c web.C, w http.ResponseWriter, r *http.Request) {

	switch c.URLParams["what"] {
	case "repo":

	}
	http.Redirect(w, r, r.Referer(), 302)
}

func apiPostHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	switch c.URLParams["what"] {
	case "repo":
		repoName := r.FormValue("name")
		desc := r.FormValue("description")
		user := "test"
		if repoName != "" {
			git.NewRepo(user, repoName).CreateRepo()
			db.NewProject(user, repoName, desc)
		}
		db.Project()
	}
	http.Redirect(w, r, r.Referer(), 302)
}

func apiPutHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	switch c.URLParams["what"] {
	case "repo":

	}
	http.Redirect(w, r, r.Referer(), 302)
}

func apiDeleteHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	switch c.URLParams["what"] {
	case "repo":

	}
	http.Redirect(w, r, r.Referer(), 302)
}
