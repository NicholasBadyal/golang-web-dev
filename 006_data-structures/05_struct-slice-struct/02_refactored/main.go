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

type tvcar struct {
	Make  string
	Model string
	Year  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	nm := tvdoctor{"Neil Melendez", "The Good Doctor"}
	gh := tvdoctor{"Gregory House", "House"}
	pc := tvdoctor{"Perry Cox", "Scrubs"}

	am := tvcar{"Aston Martin", "DB5", 1963}
	fgt := tvcar{"Ford", "GT40", 1964}
	pa := tvcar{"Pontiac", "Aztek", 2004}

	docs := []tvdoctor{nm, gh, pc}
	cars := []tvcar{am, fgt, pa}

	data := struct {
		Docs []tvdoctor
		Cars []tvcar
	}{
		docs,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
