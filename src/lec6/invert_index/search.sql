SELECT *
FROM (
	SELECT occurences.file_id, SUM(occurences.count) as sum, 
		array_agg(occurences.word_id) as words 
	FROM occurences
	JOIN words on occurences.word_id = words.id
	WHERE words.word IN ('hello', 'cat')
	GROUP BY occurences.file_id
	ORDER BY sum DESC
) t1
WHERE array_length(words, 1) = 2
