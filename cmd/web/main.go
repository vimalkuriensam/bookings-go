package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vimalkuriensam/bookings/pkg/config"
	"github.com/vimalkuriensam/bookings/pkg/handlers"
	"github.com/vimalkuriensam/bookings/pkg/render"
)

const PORT = 8080

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	render.NewTemplates(&app)
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: routes(&app),
	}
	fmt.Printf("Server running on port %d\n", PORT)
	log.Fatal(srv.ListenAndServe())
}
