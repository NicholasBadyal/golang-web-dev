package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func home(res http.ResponseWriter, req *http.Request) {
	data := "World"
	err := tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func me(res http.ResponseWriter, req *http.Request) {
	data := "Nicholas"
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
