package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First string
	Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}			// user ID, user
var dbSessions = map[string]string{}	// session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/barinfo", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	fmt.Println(dbUsers)
	fmt.Println(dbSessions)

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("Cookie not found")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		fmt.Println("Session not found")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
