package main

import (
	"io"
	"net/http"
)

type hero int

func (h hero) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "puppy, pup, doggy, dog, dogo")
	case "/cat":
		io.WriteString(w, "kitten, kit, cat, jerk")
	}
}

func main() {
	var h hero
	http.ListenAndServe(":8080", h)
}
