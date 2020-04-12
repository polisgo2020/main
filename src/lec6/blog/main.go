package main

import (
	"net/http"
	"time"

	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/config"
	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/controller"
	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/model"
	"github.com/polis-mail-ru-golang-1/examples/lec6/blog/view"

	"github.com/go-pg/pg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	die(err)

	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	zerolog.MessageFieldName = "msg"
	log.Level(level)

	log.Print(cfg)

	pgOpt, err := pg.ParseURL(cfg.PgSQL)
	die(err)
	pgdb := pg.Connect(pgOpt)
	defer pgdb.Close()

	m := model.New(pgdb)
	v, err := view.New()
	die(err)
	c := controller.New(v, m)
	s := server{
		controller: c,
		listen:     cfg.Listen,
	}
	die(s.Start())
}

type server struct {
	controller controller.Controller
	listen     string
}

func (s server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.controller.Posts)
	mux.HandleFunc("/post", s.controller.Post)

	server := http.Server{
		Addr:         s.listen,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
