CREATE TABLE public.words
(
  id serial NOT NULL,
  word text,
  CONSTRAINT words_pk PRIMARY KEY (id)
);

CREATE TABLE public.files
(
  id serial NOT NULL,
  name text,
  CONSTRAINT files_pk PRIMARY KEY (id)
);

CREATE TABLE public.occurences
(
  id serial NOT NULL,
  word_id integer,
  file_id integer,
  count integer,
  CONSTRAINT occurences_pk PRIMARY KEY (id),
  CONSTRAINT occurence_file_id FOREIGN KEY (file_id)
      REFERENCES public.files (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT occurence_word_fk FOREIGN KEY (word_id)
      REFERENCES public.words (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE CASCADE
);
