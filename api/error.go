package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type errorData struct {
	Status  int
	Message string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	p := "web/pages/error.html" //filepath.Join(exPath, "web/pages/error.html")
	fmt.Println(p)

	status := 404
	fmt.Println("error!!")
	errorTemp, err := template.New("error.html").ParseFiles(p)
	fmt.Println(err)
	errorPage := errorData{status, "Page not found"}
	w.WriteHeader(status)
	err = errorTemp.Execute(w, errorPage)
	fmt.Println(err)
}
