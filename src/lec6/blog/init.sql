CREATE TABLE public.posts
(
  id serial NOT NULL,
  title text,
  content text,
  created_at timestamp,
  CONSTRAINT post_pk PRIMARY KEY (id)
);

CREATE TABLE public.comments
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
