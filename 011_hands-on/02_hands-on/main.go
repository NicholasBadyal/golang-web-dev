package main

import (
	"log"
	"os"
	"text/template"
)

type region int

const (
	Southern region = iota
	Central
	Northern
)

type Hotel struct {
	Name   string
	Addr   string
	City   string
	Zip    string
	Region region
}

var tpl *template.Template
var fm = template.FuncMap{
	"reg2str": regionToString,
}

func regionToString(r region) string {
	if r > 2 || r < 0 {
		return "None"
	}
	return []string{
		"Southern",
		"Central",
		"Northern",
	}[r]
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []Hotel{
		{
			Name:   "The Ritz",
			Addr:   "1234 Fancy Pants Ave",
			City:   "Venice",
			Zip:    "90002",
			Region: 1,
		},
		{
			Name:   "Green Leaf Resort",
			Addr:   "803 PCH Rd",
			City:   "Palos Verdes",
			Zip:    "90009",
			Region: 2,
		},
		{
			Name:   "Hilton",
			Addr:   "8383 1st St.",
			City:   "Marina Del Rey",
			Zip:    "90012",
			Region: 0,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
