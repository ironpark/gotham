// github.com/IronPark/gotham project main.go
package main

import (
	"fmt"
	"github.com/AaronO/go-git-http"
	"github.com/AaronO/go-git-http/auth"
	"github.com/IronPark/gotham/module/db"
	"github.com/IronPark/gotham/module/middleware"
	"github.com/IronPark/gotham/module/template"
	"github.com/IronPark/gotham/module/util"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

const (
	PAGE_DASHBOARD = 1
	PAGE_PROJECT   = 2
	PAGE_SIGNUP    = 3
)

func main() {
	db.DatabaseInitialization()
	defer db.Close()
	templateErr := template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	if templateErr != nil {
		fmt.Println(templateErr)

	}
	//git smart http
	workingDir := util.WorkingDir()
	gitHandler := githttp.New(workingDir)
	gitHandler.GitBinPath = "C:/Program Files (x86)/Git/bin/git"

	authenticator := auth.Authenticator(func(info auth.AuthInfo) (bool, error) {
		if db.UserVerification(info.Username, info.Password) {
			return true, nil
		}
		return false, nil
	})

	goji.Handle("/repo/*", authenticator(gitHandler))
	goji.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(workingDir+"/static"))))
	goji.Get("/", mainHandler)
	goji.Use(middleware.SessionMiddleware)

	goji.Serve()
}

type data struct {
	Nav      int
	Projects []db.ProjectST
}

func templateObj(where int) data {
	obj := data{}
	obj.Nav = where
	switch where {
	case PAGE_PROJECT:
		obj.Projects = db.Project()
	}
	return obj
}
func mainHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(w, "index.html", templateObj(PAGE_DASHBOARD))
}
