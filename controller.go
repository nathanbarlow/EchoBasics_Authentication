package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	data := BuildDataMapFromContext(c)
	data["title"] = "Home"

	// return c.Render(http.StatusOK, "index.html", data)
	return c.Render(http.StatusOK, "index.html", data)
}

func notFound(c echo.Context) error {
	data := BuildDataMapFromContext(c)
	data["title"] = "Page Not Found"
	return c.Render(http.StatusNotFound, "404.html", data)
}

func serverError(c echo.Context) error {
	data := BuildDataMapFromContext(c)
	data["title"] = "Server Error"
	return c.Render(http.StatusNotFound, "500.html", data)
}

func login(c echo.Context) error {
	data := BuildDataMapFromContext(c)
	data["title"] = "Login"
	return c.Render(http.StatusOK, "login.html", data)
}

func planes(c echo.Context) error {
	data := BuildDataMapFromContext(c)
	data["title"] = "Planes"
	return c.Render(http.StatusOK, "planes.html", data)
}

func loginAuth(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Here you should validate the username and password against your user store
	if username == "joe" && password == "secret" {
		session, _ := store.Get(c.Request(), sessionName)
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusFound, "/planes")
	}

	data := BuildDataMapFromContext(c)
	return c.Render(http.StatusOK, "login.html", data)
}

func logout(c echo.Context) error {
	session, _ := store.Get(c.Request(), sessionName)
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}
