package main

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	// Key for session name and secret (replace with your own secret)
	sessionName = "basic_authentication-test-session-G4lb#mMg$ERXY!n2"
	sessionKey  = "eUCnaoV7ln9&3DeLd&BIWnhy21S9GC"
	store       = sessions.NewCookieStore([]byte(sessionKey))
)

func main() {
	fmt.Println("https://localhost:5050")
	// fmt.Println("http://localhost:5000")

	e := echo.New()

	// Load the session securtiy configurations
	sessionConfig()

	// Load the templates, middleware, error handler and routes
	loadTemplates(e)
	loadMiddleware(e)
	loadErrorHandler(e)
	loadRoutes(e)

	// go e.Logger.Fatal(e.Start(":5000"))

	e.Logger.Fatal(e.StartTLS(
		":5050",
		"certificate/NathanBarlow-2024-02-09-141731.cer",
		"certificate/NathanBarlow-2024-02-09-141731.pkey"))
}
