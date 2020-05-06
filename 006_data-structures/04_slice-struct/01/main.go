package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type tvdoctor struct {
	Name string
	Show string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	davidTennant := tvdoctor{
		Name: "David Tennant",
		Show: "Doctor Who",
	}

	perryCox := tvdoctor{
		Name: "Perry Cox",
		Show: "Scrubs",
	}

	neilMelendez := tvdoctor{
		Name: "Neil Melendez",
		Show: "The Good Doctor",
	}

	docs := []tvdoctor{davidTennant, perryCox, neilMelendez}

	err := tpl.Execute(os.Stdout, docs)
	if err != nil {
		log.Fatalln(err)
	}
}
