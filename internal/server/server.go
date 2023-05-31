package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	// Routing
	http.HandleFunc("/", routeServer)

	fmt.Println("Listening to: ")
	http.ListenAndServe(":3000", nil)
}
