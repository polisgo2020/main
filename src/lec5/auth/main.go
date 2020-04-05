package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	password string
	admin    bool
}

var users = map[string]User{
	"admin": {
		password: "admin",
		admin:    true,
	},
	"user": {
		password: "user",
	},
}

var sessions map[string]string

func loginPage(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	fmt.Println("Try to authorise", login, password)
	if user, ok := users[login]; ok {
		fmt.Println("User found", login)
		if user.password == password {
			fmt.Println("Password correct", login)
			sessionId := strconv.Itoa(rand.Int())
			sessions[sessionId] = login

			cookie := http.Cookie{
				Name:  "session_id",
				Value: sessionId,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusFound)

			return
		}
	}

	tpl := template.Must(template.ParseFiles("login.html"))
	tpl.Execute(w, nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	sessionId, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if login, ok := sessions[sessionId.Value]; ok {
		user := users[login]
		fmt.Fprintf(w, "User: %s, is admin: %t", login, user.admin)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func main() {
	sessions = map[string]string{}

	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", indexPage)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
