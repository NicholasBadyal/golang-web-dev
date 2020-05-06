package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
)

type data struct{
	Date string
	Open string
	High string
	Low string
	Close string
	Volume string
	AdjClose string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	csvFile, err := os.Open("scratch.csv")
	if err != nil{
		log.Fatalln(err)
	}
	defer csvFile.Close()

	datapts := []data{}

	csvr := csv.NewReader(csvFile)
	for {
		row, err := csvr.Read()
		if err != nil{
			if err == io.EOF{
				break
			} else {
				log.Fatalln(err)
			}
		}

		fmt.Println(row)
		datapts = append(datapts, data{
			Date:     row[0],
			Open:     row[1],
			High:     row[2],
			Low:      row[3],
			Close:    row[4],
			Volume:   row[5],
			AdjClose: row[6],
		})
	}

	err = tpl.Execute(os.Stdout, datapts)
	if err != nil{
		log.Fatalln(err)
	}
}