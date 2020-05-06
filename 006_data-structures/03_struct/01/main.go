package main

import (
	"log"
	"os"
	"text/template"
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

	err := tpl.Execute(os.Stdout, davidTennant)
	if err != nil {
		log.Fatalln(err)
	}
}
