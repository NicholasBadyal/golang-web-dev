package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// tvdoctors := []string{"Gregory House", "Perry Cox", "Leonard McCoy", "Neil Melendez", "David Tennant"}

	tvdoctors := map[string]string{
		"House":           "Gregory House",
		"Scrubs":          "Perry Cox",
		"Star Trek":       "Leonard McCoy",
		"The Good Doctor": "Neil Melendez",
		"Doctor Who":      "David Tennant",
	}

	err := tpl.Execute(os.Stdout, tvdoctors)
	if err != nil {
		log.Fatalln(err)
	}
}
