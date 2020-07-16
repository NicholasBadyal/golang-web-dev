package main

import (
	"fmt"
	"net/http"
)

type hero int

func (h hero) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mugs-Key", "This is from Mugs")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func<h1>")
}

func main() {
	var h hero
	http.ListenAndServe(":8080", h)
}
