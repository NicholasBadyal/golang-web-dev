package main

import (
	"fmt"
	"net/http"
	"time"
)

func getUser(req *http.Request) user {
	var u user

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un]
	return ok
}

func cleanSessions(){
	fmt.Println("Before clean:")	// for demonstration purposes
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("After clean:")
	showSessions()
}

func showSessions() {
	fmt.Println("**********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}