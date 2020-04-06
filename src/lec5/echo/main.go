package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/polisgo2020/main/src/lec5/echo/templates"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal("Invalid log level", err)
	}
	zerolog.SetGlobalLevel(logLevel)

	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// custom logger
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			zl.Debug().
				Str("method", req.Method).
				Str("remote", req.RemoteAddr).
				Str("path", req.URL.Path).
				Int("response", res.Status).
				Int("duration", int(stop.Sub(start))).
				Msgf("Called url %s", req.URL.Path)
			return nil
		}
	})

	// template renderer
	e.Renderer, err = templates.New()
	if err != nil {
		zl.Fatal().Err(err).Msg("Error initializing templates")
	}

	app := NewApp()

	// Routes
	e.GET("/", app.Index)
	e.GET("/something", app.Something)
	e.GET("/users/:id", app.User)

	err = e.Start(cfg.Listen)
	if err != nil {
		zl.Fatal().Err(err).Msg("Error starting echo")
	}
}
