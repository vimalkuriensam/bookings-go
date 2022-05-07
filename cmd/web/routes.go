package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vimalkuriensam/bookings/pkg/config"
	"github.com/vimalkuriensam/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(WriteToConsole)
	mux.Use(noSurfHandler)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
