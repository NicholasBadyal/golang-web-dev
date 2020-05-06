package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	tvdoctors := []string{"Gregory House", "Perry Cox", "Leonard McCoy", "Neil Melendez", "David Tennant"}

	err := tpl.Execute(os.Stdout, tvdoctors)
	if err != nil {
		log.Fatalln(err)
	}
}
