package main

import "github.com/gorilla/sessions"

func sessionConfig() {
	store.Options = &sessions.Options{
		Path: "/", // Available globally
		// Domain:   "example.com",
		MaxAge:   86400 * 7, // 7 days
		Secure:   true,      // Only send over HTTPS
		HttpOnly: true,      // Prevent access via JavaScript
	}
}
