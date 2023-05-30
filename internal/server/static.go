package server

import (
	"net/http"
	"strings"
)

func staticServer(w http.ResponseWriter, r *http.Request) {
	p := strings.Replace(r.URL.Path, "/static/", "", 1)
	p = "web/static/" + p
	http.ServeFile(w, r, p)
}
