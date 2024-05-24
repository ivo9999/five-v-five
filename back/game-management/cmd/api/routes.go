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

	r.Post("/teams", h.CreateTeams)               // New endpoint to create balanced teams
	r.Post("/games", h.InitializeGame)            // Updated endpoint to create a game from two teams
	r.Get("/games/{gameId}", h.GetGame)           // New endpoint to get a game by ID
	r.Post("/games/champions", h.GetChamps)       // Updated endpoint to get champions for a list of players and lanes
	r.Get("/champion", h.GetChampion)             // Updated endpoint to get a champion for a player and lane
	r.Post("/games/{gameId}/winner", h.SetWinner) // Updated endpoint to set the winner of a game
	r.Post("/games/{gameId}/move", h.MovePlayer)  // Updated endpoint to move a player between teams
	r.Post("/dropdb", h.DropDB)
}
