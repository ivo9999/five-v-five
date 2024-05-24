package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	handlers := NewHandlers(app.DB, app.RiotAPI, app)

  fmt.Println("check check")
	handlers.RegisterRoutes(mux)

	return mux
}

func (h *Handlers) RegisterRoutes(r chi.Router) {
	r.Use(middleware.Logger)

	r.Post("/games", h.InitializeGame)
	r.Post("/games/{gameId}/add/{team}", h.AddPlayerToTeam)
	r.Post("/games/{gameId}/move", h.MovePlayer)
	r.Post("/games/{gameId}/winner", h.SetWinner)
	r.Get("/champion", h.GetChampion)
	r.Post("/games/{gameId}/champions", h.GetGameChampions)
}
