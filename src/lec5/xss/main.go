package main

import (
	"fmt"
	"net/http"
	"strings"
)

// comment<script>alert('111');</script>

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body>
		<form action='/comment' method='post'>
			<input type='text' name='comment'>
			<input type='submit'>
		</form>
		</body></html>`))
	w.Write([]byte(`</body></html>`))
}

func comment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><head><meta charset="utf-8"></head><body>`))
	comment := r.FormValue("comment")
	//comment = clean(comment)
	w.Write([]byte(fmt.Sprintf("Comment: %s<br/>", comment)))
	w.Write([]byte(fmt.Sprintf("End")))
	w.Write([]byte(`</body></html>`))
}

func clean(in string) string {
	in = strings.Replace(in, "&", "&amp;", -1)
	in = strings.Replace(in, "<", "&lt;", -1)
	in = strings.Replace(in, ">", "&gt;", -1)
	in = strings.Replace(in, "'", "&prime;", -1)
	return strings.Replace(in, `"`, "&quot;", -1)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/comment", comment)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
