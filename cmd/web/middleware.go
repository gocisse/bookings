package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// This is an example and prototy of middleware functions
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("page got a hit ")
	})
}

// Csrf handler using nosurl to set cookie parameters
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Sessions loader
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
