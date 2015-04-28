package controller

import (
	"bytes"
	"github.com/IronPark/gotham/module/template"
	"github.com/gorilla/sessions"
	"github.com/zenazn/goji/web"
)

type Controller struct {
}

func (controller *Controller) GetSession(c web.C) *sessions.Session {
	return c.Env["Session"].(*sessions.Session)
}

func (controller *Controller) RenderTemplate(filename string, obj interface{}) string {
	var doc bytes.Buffer
	template.Render(&doc, filename, obj)
	return doc.String()
}

func (controller *Controller) IsXhr(c web.C) bool {
	return c.Env["IsXhr"].(bool)
}
