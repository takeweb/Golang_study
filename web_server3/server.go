package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Template テンプレート
type Template struct {
	templates *template.Template
}

// Render レンダー
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Hello ハロー
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = t
	e.GET("/hello", Hello)

	e.Logger.Fatal(e.Start(":8000"))
}
