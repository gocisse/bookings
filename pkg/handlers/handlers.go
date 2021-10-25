package handlers

import (
	"net/http"

	"github.com/gocisse/bookings/pkg/config"
	"github.com/gocisse/bookings/pkg/models"
	"github.com/gocisse/bookings/pkg/render"
)

//Repo with hold our Repository
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// create a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//Create a new handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	remoteHost := r.Host
	m.App.Session.Get(r.Context(), r.Host)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		Flash: remoteHost,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "It works"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// add it to our render
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
