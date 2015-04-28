package controllers

import (
	"github.com/IronPark/gotham/models"
	"github.com/IronPark/gotham/module/controller"
	"github.com/IronPark/gotham/module/git"
	"github.com/zenazn/goji/web"
	"net/http"
)

type ApiController struct {
	controller.Controller
}

func (ctr *ApiController) Get(c web.C, r *http.Request) (string, int) {
	c.Env["Title"] = "Default Project - free Go website project template"
	c.Env["Nav"] = PAGE_DASHBOARD

	return ctr.RenderTemplate("index.html", c.Env), http.StatusOK
}

func (ctr *ApiController) Post(c web.C, r *http.Request) (string, int) {

	repoName := r.FormValue("name")
	desc := r.FormValue("description")
	user := "test"
	if repoName != "" {
		git.NewRepo(user, repoName).CreateRepo()
		models.NewProject(user, repoName, desc)
	}
	return ctr.RenderTemplate("index.html", c.Env), http.StatusOK
}
