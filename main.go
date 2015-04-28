// github.com/IronPark/gotham project main.go
package main

import (
	"fmt"
	"github.com/AaronO/go-git-http"
	"github.com/AaronO/go-git-http/auth"
	//	"github.com/IronPark/gotham/module/middleware"
	"github.com/IronPark/gotham/controllers"
	"github.com/IronPark/gotham/models"
	"github.com/IronPark/gotham/module/template"
	"github.com/IronPark/gotham/module/util"
	"github.com/gorilla/sessions"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"io"
	"net/http"
	"reflect"
)

func main() {
	workdir := util.WorkingDir()
	templateErr := template.LoadTemplates(workdir+"/view", ".html")
	if templateErr != nil {
		fmt.Println(templateErr)

	}

	//git smart http
	gitHandler := githttp.New(util.WorkingDir())
	gitHandler.GitBinPath = "C:/Program Files (x86)/Git/bin/git"

	authenticator := auth.Authenticator(func(info auth.AuthInfo) (bool, error) {
		if models.GetUserByEmail(info.Username).EqPassword(info.Password) {
			return true, nil
		}
		return false, nil
	})

	goji.Handle("/repo/*", authenticator(gitHandler))
	goji.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(workdir+"/static"))))

	controller := &controllers.MainController{}
	apiController := &controllers.ApiController{}
	goji.Get("/", Route(controller, "Index"))
	goji.Get("/project", Route(controller, "ProjectList"))
	goji.Get("/project/:user/:name", Route(controller, "Project"))

	goji.Post("/api/repo", Route(apiController, "Post"))
	goji.Serve()
}

func Route(controller interface{}, route string) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		c.Env["Content-Type"] = "text/html"

		methodValue := reflect.ValueOf(controller).MethodByName(route)
		methodInterface := methodValue.Interface()
		method := methodInterface.(func(c web.C, r *http.Request) (string, int))

		body, code := method(c, r)

		if session, exists := c.Env["Session"]; exists {
			err := session.(*sessions.Session).Save(r, w)
			if err != nil {
				//glog.Errorf("Can't save session: %v", err)
			}
		}

		switch code {
		case http.StatusOK:
			if _, exists := c.Env["Content-Type"]; exists {
				w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
			}
			io.WriteString(w, body)
		case http.StatusSeeOther, http.StatusFound:
			http.Redirect(w, r, body, code)
		}
	}
	return fn
}

//package main

//import (
//	"fmt"
//	"github.com/IronPark/gotham/module/db"
//	"github.com/IronPark/gotham/module/template"
//	"github.com/IronPark/gotham/module/util"
//	"github.com/zenazn/goji"
//	"github.com/zenazn/goji/web"
//	"net/http"
//)

//func init() {
//	goji.Get("/signup", signupHandler)
//	goji.Post("/signup", signupFormHandler)

//	goji.Get("/signin", signinHandler)
//}

//func signupHandler(c web.C, w http.ResponseWriter, r *http.Request) {
//	template.LoadTemplates(util.WorkingDir()+"/view", ".html")
//	obj := struct {
//		Nav int
//	}{
//		Nav: PAGE_SIGNUP,
//	}
//	err := template.Render(w, "sign-up.html", obj)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//func signupFormHandler(c web.C, w http.ResponseWriter, r *http.Request) {
//	template.LoadTemplates(util.WorkingDir()+"/view", ".html")
//	inputEmail := r.FormValue("inputEmail")
//	inputPassword := r.FormValue("inputPassword")
//	inputPasswordCheck := r.FormValue("inputPasswordCheck")
//	inputName := r.FormValue("inputName")
//	fmt.Println(inputEmail, inputName, inputPassword, inputPasswordCheck)

//	if inputPassword == inputPasswordCheck &&
//		inputEmail != "" &&
//		inputName != "" &&
//		inputPassword != "" {
//		db.InsertUser(inputEmail, inputName, inputPassword)
//		fmt.Println("Insert User")
//	}

//	obj := struct {
//		Nav int
//	}{
//		Nav: PAGE_SIGNUP,
//	}

//	err := template.Render(w, "sign-up.html", obj)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//func signinHandler(c web.C, w http.ResponseWriter, r *http.Request) {
//	err := template.LoadTemplates(util.WorkingDir()+"/view", ".html")
//	obj := struct {
//		Nav int
//	}{
//		Nav: PAGE_SIGNUP,
//	}
//	err = template.Render(w, "sign-in.html", obj)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
