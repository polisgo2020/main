package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
)

type Config struct {
	Listen   string `env:"LISTEN" envDefault:"localhost:8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

func loadConfig() (*Config, error) {
	cfg := &Config{}
	return cfg, env.Parse(cfg)
}

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

	app := NewApp()

	// Routes
	e.POST("/users", app.CreateUser)
	e.GET("/users/:id", app.GetUser)

	err = e.Start(cfg.Listen)
	if err != nil {
		zl.Fatal().Err(err).Msg("Error starting echo")
	}
}

type App struct{}

func NewApp() *App {
	return &App{}
}

type User struct {
	Name string `json:"name" form:"name" query:"name"`
	ID   int    `json:"id" form:"id" query:"id"`
}

func (a *App) CreateUser(c echo.Context) (err error) {
	u := &User{}
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func (a *App) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	return c.JSON(http.StatusOK, User{
		Name: "Name",
		ID:   id,
	})
}
