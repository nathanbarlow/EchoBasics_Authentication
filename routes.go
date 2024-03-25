package main

import "github.com/labstack/echo/v4"

func loadRoutes(e *echo.Echo) {
	e.GET("/", index)
	e.GET("/404", notFound)
	e.GET("/login", login)
	e.POST("/login", loginAuth)
	e.GET("/logout", logout)

	e.GET("/planes", planes, isAuthenticated)

	// e.GET("/users", getUsers)
	// e.GET("/users/:id", getUser)
	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
}
