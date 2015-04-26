package main

import (
	"fmt"
	"github.com/IronPark/gotham/module/db"
	"github.com/IronPark/gotham/module/template"
	"github.com/IronPark/gotham/module/util"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func init() {
	goji.Get("/signup", signupHandler)
	goji.Post("/signup", signupFormHandler)

	goji.Get("/signin", signinHandler)
}

func signupHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	obj := struct {
		Nav int
	}{
		Nav: PAGE_SIGNUP,
	}
	err := template.Render(w, "sign-up.html", obj)
	if err != nil {
		fmt.Println(err)
	}
}
func signupFormHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	inputEmail := r.FormValue("inputEmail")
	inputPassword := r.FormValue("inputPassword")
	inputPasswordCheck := r.FormValue("inputPasswordCheck")
	inputName := r.FormValue("inputName")
	fmt.Println(inputEmail, inputName, inputPassword, inputPasswordCheck)

	if inputPassword == inputPasswordCheck &&
		inputEmail != "" &&
		inputName != "" &&
		inputPassword != "" {
		db.InsertUser(inputEmail, inputName, inputPassword)
		fmt.Println("Insert User")
	}

	obj := struct {
		Nav int
	}{
		Nav: PAGE_SIGNUP,
	}

	err := template.Render(w, "sign-up.html", obj)
	if err != nil {
		fmt.Println(err)
	}
}
func signinHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	err := template.LoadTemplates(util.WorkingDir()+"/view", ".html")
	err = template.Render(w, "sign-up.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
