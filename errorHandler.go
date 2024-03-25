package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func loadErrorHandler(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					if he.Code == http.StatusNotFound {
						return c.Redirect(http.StatusFound, "/404")
					} else if he.Code == http.StatusUnauthorized {
						return c.Redirect(http.StatusFound, "/login")
					}
				}
				return err
			}
			return nil
		}
	})
}
