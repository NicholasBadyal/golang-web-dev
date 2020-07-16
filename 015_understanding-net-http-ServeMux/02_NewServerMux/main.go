package main

import (
	"io"
	"net/http"
)

type hero int

func (h hero) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type villain int

func (v villain) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hero
	var c villain

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
