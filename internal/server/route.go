package server

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/nekiwo/imsa-grades/api"
)

type Route struct {
	path    *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(path string, handler http.HandlerFunc) Route {
	return Route{regexp.MustCompile("^" + path + "($|.html$)"), handler}
}

var routeSwitch = []Route{
	newRoute("/", api.IndexRoute),
	newRoute("/static/([^$]+)", api.StaticRoute),
	newRoute("/index", api.IndexRoute),
	newRoute("/class/([^/]+)", api.ClassRoute),
	newRoute("/about", api.AboutRoute),
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

	api.ErrorRoute(w, r, http.StatusNotFound)
}
