package main

import (
	"fmt"
	"net/http"
)

type hero int

func (h hero) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}