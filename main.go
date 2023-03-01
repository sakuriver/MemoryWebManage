package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("static/template/*.html")),
	}

	e.Renderer = t

	e.GET("/top", TopHtml)

	e.GET("/client_data_top", ClientDataTopHtml)

	e.GET("/season_top", SeasonHtml)

	e.GET("/season_top", SeasonHtml)

	e.GET("/account_search", AccountSearchHtml)

	e.GET("/log_search", LogSearchHtml)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type DataForm struct{}

func TopHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "top", DataForm{})
}

func ClientDataTopHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "client_data_top", DataForm{})
}

func SeasonHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "season_top", DataForm{})
}

func AccountSearchHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "account_search", DataForm{})
}

func LogSearchHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "log_search", DataForm{})
}
