package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gocisse/bookings/pkg/handlers"
	"github.com/gocisse/bookings/pkg/render"
	"github.com/gocisse/bookings/pkg/config"
)

const portNumber = ":4000"

//declare our app config
var app config.AppConfig

// declare a session manager to use in main
var session *scs.SessionManager

func main() {

	// set in production global variable
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	// add our session to our app
	app.Session = session

	// create a template cache to store in our app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("failed to create template cache")
	}

	// set our template to our new create cache
	app.TemplateCache = tc

	// set dev or production variable
	app.UseCache = false

	// Create a repo for the handlers
	// Add Repo to new handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//Create our NewTemplate cache based on info above
	render.NewTemplate(&app)

	fmt.Println("Starting server on port ", portNumber)
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	_ = srv.ListenAndServe()

}
