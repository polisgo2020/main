package view

import (
	"html/template"
	"io"

	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/model"
)

type View struct {
	postsT *template.Template
	postT  *template.Template
	addT   *template.Template
	errorT *template.Template
}

func New() (View, error) {
	v := View{}
	var err error

	v.postsT, err = template.ParseFiles("view/templates/layout.html", "view/templates/posts.html")
	if err != nil {
		return v, err
	}
	v.postT, err = template.ParseFiles("view/templates/layout.html", "view/templates/post.html")
	if err != nil {
		return v, err
	}
	v.addT, err = template.ParseFiles("view/templates/layout.html", "view/templates/add.html")
	if err != nil {
		return v, err
	}
	v.errorT, err = template.ParseFiles("view/templates/layout.html", "view/templates/error.html")
	if err != nil {
		return v, err
	}

	return v, nil
}

func (v View) Posts(posts []model.Post, wr io.Writer) error {
	return v.postsT.ExecuteTemplate(wr, "layout",
		struct {
			Title string
			Posts []model.Post
		}{
			Title: "Posts",
			Posts: posts,
		})
}

func (v View) Post(post model.Post, comments []model.Comment, wr io.Writer) error {
	return v.postT.ExecuteTemplate(wr, "layout",
		struct {
			Title    string
			Post     model.Post
			Comments []model.Comment
		}{
			Title:    post.Title,
			Post:     post,
			Comments: comments,
		})
}

func (v View) Error(displayErr string, status int, wr io.Writer) error {
	return v.errorT.ExecuteTemplate(wr, "layout",
		struct {
			Title  string
			Status int
			Error  string
		}{
			Title:  "error",
			Status: status,
			Error:  displayErr,
		})
}
