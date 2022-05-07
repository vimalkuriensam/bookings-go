package handlers

import (
	"net/http"

	"github.com/vimalkuriensam/bookings/pkg/config"
	"github.com/vimalkuriensam/bookings/pkg/models"
	"github.com/vimalkuriensam/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (r *Repository) Home(w http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	r.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About is the about page handler
func (r *Repository) About(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := r.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
