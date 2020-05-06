package main

import (
	"log"
	"os"
	"strings"
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

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
