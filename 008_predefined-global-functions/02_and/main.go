package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Age   int
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user{"Joe", 45, false}
	u2 := user{"Shaun", 31, true}
	u3 := user{"Jesse", 23, true}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
