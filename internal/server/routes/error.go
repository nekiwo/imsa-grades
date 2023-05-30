package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorData struct {
	Status  int
	Message string
}

func ErrorRoute(w http.ResponseWriter, r *http.Request, status int) {
	fmt.Println("error!!")
	errorTemp, _ := template.New("error.html").ParseFiles("web/pages/error.html")
	errorPage := ErrorData{status, "Page not found"}
	w.WriteHeader(status)
	errorTemp.Execute(w, errorPage)
}
