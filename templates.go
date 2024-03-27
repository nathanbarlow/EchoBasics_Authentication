package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func loadTemplates(e *echo.Echo) {
	// Initialize the template renderer
	templates := template.Must(template.ParseGlob("public/views/*/*.html"))

	renderer := &TemplateRenderer{
		templates: templates,
	}
	e.Renderer = renderer

}
