package middleware

import (
	//"github.com/zenazn/goji"
	"fmt"
	"github.com/IronPark/gotham/module/session"
	"github.com/zenazn/goji/web"
	"net/http"
)

func SessionMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/signin" {
			h.ServeHTTP(w, r)
			return
		} else {
			fmt.Println(r.URL.Path)
		}
		cookie, err := r.Cookie("SESSION")
		cookie_id, err := r.Cookie("SESSION_ID")
		if err != nil {
			http.Redirect(w, r, "/signin", 302)
			return
		}
		token := cookie.Value
		id := cookie_id.Value

		if session.SessionCheck(id, token) {
			c.Env["User"] = id
			h.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/signin", 302)
			return
		}

	}
	return http.HandlerFunc(fn)
}
