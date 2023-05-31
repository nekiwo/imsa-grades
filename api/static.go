package handler

import (
	"net/http"
	"strings"
)

func StaticRoute(w http.ResponseWriter, r *http.Request) {
	p := strings.Replace(r.URL.Path, "/static/", "", 1)
	p = "web/static/" + p
	http.ServeFile(w, r, p)
}
