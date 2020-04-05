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

func IsUserOdd(u *User) bool {
	return u.ID%2 != 0
}

func main() {
	tmplFuncs := template.FuncMap{
		"OddUser": IsUserOdd,
	}

	tmpl, err := template.
		New("").
		Funcs(tmplFuncs).
		ParseFiles("func.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "User 1", true},
		User{2, "User 2", false},
		User{3, "User 3", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "func.html",
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
