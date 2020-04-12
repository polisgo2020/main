package model

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Post empty
type Post struct {
	ID        int       `sql:"id,pk"`
	Title     string    `sql:"title"`
	Content   string    `sql:"content"`
	CreatedAt time.Time `sql:"created_at"`
}

func (p *Post) BeforeInsert(db orm.DB) error {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	return nil
}

// Comment empty
type Comment struct {
	ID      int    `sql:"id,pk"`
	Author  string `sql:"author"`
	Content string `sql:"content"`
	PostID  int
}

type Model struct {
	pg *pg.DB
}

func New(pg *pg.DB) Model {
	return Model{
		pg: pg,
	}
}

func (m Model) Posts() ([]Post, error) {
	var posts []Post
	err := m.pg.Model(&posts).Order("created_at ASC").Select()
	return posts, err
}

func (m Model) Post(id int) (Post, error) {
	post := Post{ID: id}
	err := m.pg.Select(&post)
	return post, err
}

func (m Model) AddPost(post Post) (Post, error) {
	_, err := m.pg.Model(&post).Returning("*").Insert()
	return post, err
}

func (m Model) Comments(post Post) ([]Comment, error) {
	comments := []Comment{}
	err := m.pg.Model(&comments).Where("post_id = ?", post.ID).Order("id ASC").Select()
	return comments, err
}

func (m Model) AddComment(comment Comment) (Comment, error) {
	var id int
	_, err := m.pg.Query(
		pg.Scan(&id),
		"INSERT INTO \"comments\" (author, content, post_id) VALUES (?, ?, ?) RETURNING \"id\"",
		comment.Author, comment.Content, comment.PostID,
	)
	comment.ID = id
	return comment, err
}
