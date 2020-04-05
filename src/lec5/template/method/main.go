package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func (u *User) PrintActive() string {
	if !u.Active {
		return ""
	}
	return "method says user " + u.Name + " active"
}

func main() {
	tmpl, err := template.
		New("").
		ParseFiles("method.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "User 1", true},
		User{2, "User 2", false},
		User{3, "User 3", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "method.html",
			struct {
				Users []User
			}{
				users,
			})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
