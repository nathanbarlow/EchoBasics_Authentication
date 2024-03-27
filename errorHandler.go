package main

import (
	"fmt"
	"net/http"
	"runtime"

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
				} else {
					_, file, line, _ := runtime.Caller(1)
					return c.String(http.StatusInternalServerError, fmt.Sprintf("An error occurred: %v\nFile: %s\nLine: %d", err, file, line))
				}
				return err
			}
			return nil
		}
	})
}
