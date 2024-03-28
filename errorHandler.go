package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/labstack/echo/v4"
)

func loadErrorHandler(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = next(c)
			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					switch he.Code {
					case http.StatusNotFound:
						return c.Redirect(http.StatusFound, "/404")
					case http.StatusUnauthorized:
						return c.Redirect(http.StatusFound, "/login")
					default:
						return handleInternalServerError(c, err)
					}
				}
				return handleInternalServerError(c, err)
			}
			return nil
		}
	})
}

func handleInternalServerError(c echo.Context, err error) error {
	if os.Getenv("ENV") == "development" {
		_, file, line, _ := runtime.Caller(0)
		details := fmt.Sprintf("An error occurred: %v\nFile: %s\nLine: %d", err, file, line)
		log.Println(details)
		return c.String(http.StatusInternalServerError, details)
	}
	log.Println("Internal Server Error: ", err)
	return c.Redirect(http.StatusFound, "/500")
}
