package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	Templates *template.Template
}

func New() (*TemplateRenderer, error) {
	tpls, err := template.ParseGlob("templates/*.html")
	return &TemplateRenderer{
		Templates: tpls,
	}, err
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
