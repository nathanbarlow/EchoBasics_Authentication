package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func loadTemplates(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*/*.html")),
	}
	e.Renderer = t
}
