package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func loadMiddleware(e *echo.Echo) {
	if os.Getenv("ENV") == "development" {
		e.Debug = true
	}

	e.Static("/", "/public/static")

	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(session.MiddlewareWithConfig(session.Config{
		Store: store,
	}))

	e.Use(sessionToContextMiddleware)
}

// Middleware to check if the user is authenticated
func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), sessionName)
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			return c.Redirect(http.StatusFound, "/login")
		}
		return next(c)
	}
}

func sessionToContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), sessionName)
		userName, ok := session.Values["username"].(string)
		if !ok {
			// Handle missing session data if necessary
			userName = "Guest"
		}
		loggedIn, ok := session.Values["authenticated"].(bool)
		if !ok {
			// Handle missing session data if necessary
			loggedIn = false
		}

		// You could use a custom context here, but for simplicity, we'll use the standard one
		// and set a custom value that handlers can access.
		c.Set("userName", userName)
		c.Set("loggedIn", loggedIn)

		// Proceed with the request
		return next(c)
	}
}
