package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gocisse/bookings/pkg/config"
	"github.com/gocisse/bookings/pkg/handlers"
)

// routes will hold our route and middleware
func routes(a *config.AppConfig) http.Handler {
	// declare our mux handler to return
	mux := chi.NewRouter()

	// Middlewares goes here
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//page handlers
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//serve static files with fileserver
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
