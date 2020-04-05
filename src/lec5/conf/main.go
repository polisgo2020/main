package main

import (
	"fmt"
	"net/http"
)

var cfg Config

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, cfg.Greeting)
	w.Write([]byte("!!!"))
}

func main() {
	// cfg = Load()
	cfg = Load2()

	http.HandleFunc("/", handler)
	fmt.Println("starting server at ", cfg.Listen)
	http.ListenAndServe(cfg.Listen, nil)
}
