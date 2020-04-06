package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (a *App) Something(c echo.Context) error {
	return c.Render(http.StatusOK, "template.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func (a *App) User(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	return c.JSON(http.StatusOK, struct {
		ID   int
		Name string
	}{
		ID:   id,
		Name: "name",
	})
}
