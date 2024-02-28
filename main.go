package main

import (
	"fmt"
	"net/http"
	"text/template"

	"go.dev/femm/data"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// handle template
func handleTemplate(w http.ResponseWriter, r *http.Request) {
	// parse template
	html, err := template.ParseFiles("template/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error: " + err.Error()))
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handler)

	//  file server
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	// handle template
	server.HandleFunc("/template", handleTemplate)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
