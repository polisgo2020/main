package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"

	zl "github.com/rs/zerolog/log"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

type LogMiddle struct {
	StdLogger *log.Logger
}

func (lm *LogMiddle) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		fmt.Printf("FMT [%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))

		log.Printf("LOG [%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))

		lm.StdLogger.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))

		zl.Debug().
			Str("method", r.Method).
			Str("remote", r.RemoteAddr).
			Str("path", r.URL.Path).
			Int("duration", int(time.Since(start))).
			Msgf("Called url %s", r.URL.Path)
	})
}

// -----------

func main() {

	addr := "localhost"
	port := 8080

	// std
	fmt.Printf("STD starting server at %s:%d\n", addr, port)

	// std
	log.Printf("STD starting server at %s:%d", addr, port)

	lm := new(LogMiddle)

	// std
	lm.StdLogger = log.New(os.Stdout, "STD ", log.LUTC|log.Lshortfile)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zl.Logger = zl.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// server stuff
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/", mainPage)
	siteHandler := lm.logMiddleware(siteMux)
	http.ListenAndServe(":8080", siteHandler)
}
