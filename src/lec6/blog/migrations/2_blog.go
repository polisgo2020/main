package main

import (
	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`CREATE TABLE public.posts
		(
		  id serial NOT NULL,
		  title text,
		  content text,
		  created_at timestamp,
		  CONSTRAINT post_pk PRIMARY KEY (id)
		);`)
		if err != nil {
			return err
		}
		_, err = db.Exec(`CREATE TABLE public.comments
		(
		  id serial NOT NULL,
		  post_id integer,
		  author text,
		  content text,
		  CONSTRAINT comment_pk PRIMARY KEY (id),
		  CONSTRAINT comment_post_id FOREIGN KEY (post_id)
			  REFERENCES public.posts (id) MATCH SIMPLE
			  ON UPDATE NO ACTION ON DELETE CASCADE
		);
		`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`DROP TABLE public.posts;`)
		if err != nil {
			return err
		}
		_, err = db.Exec(`DROP TABLE public.comments;`)
		return err
	})
}
