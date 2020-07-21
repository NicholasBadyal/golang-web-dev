package main

import (
	"io"
	"net/http"
)

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func main() {
	http.Handle("/cats", http.HandlerFunc(c))
	http.Handle("/dogs", http.HandlerFunc(d))

	http.ListenAndServe(":8080", nil)
}
