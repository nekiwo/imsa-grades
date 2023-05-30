package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	// Routing
	http.HandleFunc("/", routeServer)

	// Static hosting
	http.HandleFunc("/static/", staticServer)

	fmt.Println("Listening to: ")
	http.ListenAndServe(":3000", nil)
}
