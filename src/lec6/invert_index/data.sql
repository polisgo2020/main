INSERT INTO public.files (id, name) VALUES (1, 'document1');
INSERT INTO public.files (id, name) VALUES (2, 'document2');
INSERT INTO public.files (id, name) VALUES (3, 'document3');


INSERT INTO public.words (id, word) VALUES (1, 'hello');
INSERT INTO public.words (id, word) VALUES (2, 'cat');
INSERT INTO public.words (id, word) VALUES (3, 'mother');
INSERT INTO public.words (id, word) VALUES (4, 'wash');
INSERT INTO public.words (id, word) VALUES (5, 'window');


INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (1, 1, 1, 2);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (2, 2, 1, 1);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (3, 3, 2, 1);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (4, 4, 2, 1);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (5, 2, 2, 1);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (6, 3, 3, 2);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (7, 4, 3, 1);
INSERT INTO public.occurences (id, word_id, file_id, count) VALUES (8, 5, 3, 1);


SELECT pg_catalog.setval('public.files_id_seq', 3, true);
SELECT pg_catalog.setval('public.occurences_id_seq', 8, true);
SELECT pg_catalog.setval('public.words_id_seq', 5, true);
