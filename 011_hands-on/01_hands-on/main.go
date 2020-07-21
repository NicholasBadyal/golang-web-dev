package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcademicYear         string
	Fall, Spring, Summer semester
}

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	years := []year{
		{
			AcademicYear: "2018",
			Fall: semester{
				Term: "Fall18",
				Courses: []course{
					{
						Number: "CS180000",
						Name:   "Object-Oriented Programming",
						Units:  "4",
					},
					{
						Number: "CS240000",
						Name:   "Basics of C",
						Units:  "4",
					},
					{
						Number: "CS315000",
						Name:   "Information Search and Management",
						Units:  "5",
					},
				},
			},
			Spring: semester{
				Term: "Spring18",
				Courses: []course{
					{
						Number: "CS192000",
						Name:   "CS Student Career Seminar",
						Units:  "1",
					},
					{
						Number: "MA107001",
						Name:   "Calculus 1",
						Units:  "3",
					},
				},
			},
		},
		{
			AcademicYear: "2019",
			Summer: semester{
				Term: "Summer19",
				Courses: []course{
					{
						Number: "CS250000",
						Name:   "Systems Programming in C++",
						Units:  "5",
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}
}
