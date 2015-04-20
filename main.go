// github.com/IronPark/gotham project main.go
package main

import (
	"fmt"
	"github.com/AaronO/go-git-http"
	"github.com/IronPark/gotham/module/db"
	"github.com/IronPark/gotham/module/template"
	"github.com/IronPark/gotham/module/util"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
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
	//	gitHandler.ProjectRoot = "/"
	goji.Handle("/repo/*", gitHandler)
	goji.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(workingDir+"/static"))))
	goji.Get("/", mainHandler)
	goji.Serve()
}

func mainHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(w, "index.html", nil)
}
