package main

import "github.com/labstack/echo/v4"

func BuildDataMapFromContext(c echo.Context) map[string]interface{} {
	data := map[string]interface{}{}
	data["userName"] = c.Get("userName")
	data["loggedIn"] = c.Get("loggedIn")
	return data
}
