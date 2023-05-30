package server

import (
	"github.com/nekiwo/imsa-grades/internal/server/routes"
	"net/http"
	"regexp"
	"strings"
)

type Route struct {
	path    *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(path string, handler http.HandlerFunc) Route {
	return Route{regexp.MustCompile("^" + path + "($|.html$)"), handler}
}

var routeSwitch = []Route{
	newRoute("/", routes.IndexRoute),
	newRoute("/index", routes.IndexRoute),
	newRoute("/class/([^/]+)", routes.ClassRoute),
	newRoute("/about", routes.AboutRoute),
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

	routes.ErrorRoute(w, r, http.StatusNotFound)
}
