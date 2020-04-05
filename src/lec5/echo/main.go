package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/", hello)
	// Named route "foobar"
	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly!",
		})
	})
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, struct {
		ID   int
		Name string
	}{
		ID:   id,
		Name: "name",
	})
}
