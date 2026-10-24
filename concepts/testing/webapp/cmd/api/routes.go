package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// protected routes
	mux.Route("/users", func(r chi.Router) {
		// use auth middleware
		mux.Use(app.authRequired)

		mux.Get("/", app.allUsers)
		mux.Get("/{userId}", app.getUser)
		mux.Delete("/{userId}", app.deleteUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/{userId}", app.updateUser)
	})

	return mux
}
