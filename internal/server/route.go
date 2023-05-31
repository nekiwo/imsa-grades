package server

import (
	"net/http"
	"regexp"
	"strings"

	handler "github.com/nekiwo/imsa-grades/api"
)

type Route struct {
	path    *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(path string, handler http.HandlerFunc) Route {
	return Route{regexp.MustCompile("^" + path + "($|.html$)"), handler}
}

var routeSwitch = []Route{
	newRoute("/", handler.IndexRoute),
	newRoute("/static/([^$]+)", handler.StaticRoute),
	newRoute("/index", handler.IndexRoute),
	newRoute("/class/([^/]+)", handler.ClassRoute),
	newRoute("/about", handler.AboutRoute),
}

func routeServer(w http.ResponseWriter, r *http.Request) {
	url := strings.Replace(r.URL.Path, ".html", "", 1)

	for _, route := range routeSwitch {
		matches := route.path.FindStringSubmatch(url)
		if len(matches) > 0 {
			route.handler(w, r)
			return
		}
	}

	handler.ErrorRoute(w, r, http.StatusNotFound)
}
