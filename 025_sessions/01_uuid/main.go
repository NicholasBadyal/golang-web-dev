package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:       "session",
			Value:      id.String(),
			Path:       "/",
			HttpOnly:   true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
