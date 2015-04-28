package controllers

import (
	"github.com/IronPark/gotham/models"
	"github.com/IronPark/gotham/module/controller"
	"github.com/IronPark/gotham/module/git"
	"github.com/zenazn/goji/web"
	"net/http"
)

const (
	PAGE_DASHBOARD = 1
	PAGE_PROJECT   = 2
	PAGE_SIGNUP    = 3
)

type MainController struct {
	controller.Controller
}

// Home page route
func (ctr *MainController) Index(c web.C, r *http.Request) (string, int) {
	c.Env["Title"] = "Default Project - free Go website project template"
	c.Env["Nav"] = PAGE_DASHBOARD

	return ctr.RenderTemplate("index.html", c.Env), http.StatusOK
}

// Project route
func (ctr *MainController) ProjectList(c web.C, r *http.Request) (string, int) {
	c.Env["Title"] = "Default Project - free Go website project template"
	c.Env["Nav"] = PAGE_PROJECT
	c.Env["Projects"] = models.Projects()

	return ctr.RenderTemplate("project-list.html", c.Env), http.StatusOK
}

func (ctr *MainController) Project(c web.C, r *http.Request) (string, int) {
	c.Env["Title"] = "Default Project - free Go website project template"
	c.Env["Nav"] = PAGE_PROJECT
	c.Env["Projects"] = models.Projects()

	Tree, _ := git.NewRepo(c.URLParams["user"], c.URLParams["name"]).FileList()
	c.Env["User"] = c.URLParams["user"]
	c.Env["ProjectName"] = c.URLParams["name"]
	c.Env["FileTree"] = Tree

	return ctr.RenderTemplate("project-info.html", c.Env), http.StatusOK
}
