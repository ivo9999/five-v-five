package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management-micro/data"
	"user-management-micro/riot"

	"github.com/go-chi/chi/v5"
)

// createAccount handler
func (app *Config) createAccount(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := data.InsertUser(app.DB.DB, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// updateAccount handler
func (app *Config) updateAccount(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.UpdateUser(app.DB.DB, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// getAccountById handler
func (app *Config) getAccountById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := data.GetUser(app.DB.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// getAllUsers
func (app *Config) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := data.GetAllUsers(app.DB.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusOK, users)
}

// getAccountByUsername handler
func (app *Config) getAccountByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := data.GetUserByUsername(app.DB.DB, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *Config) fillUserDataRiot(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := data.GetUserByUsername(app.DB.DB, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateSumoner := &riot.SummonerByNameRequest{
		Name: user.LeagueName,
		Tag:  user.LeagueTag,
	}

	updateChampion := &riot.ChampionMasteriesRequest{
		SummonerId: user.LeagueName,
	}

	updateLeagues := &riot.LeagueEntriesRequest{
		SummonerId: user.LeagueName,
	}

	// Update summoner by name
	if _, err := app.RiotAPI.UpdateSummonerByName(r.Context(), updateSumoner); err != nil {
		http.Error(w, "Failed to update summoner by name", http.StatusInternalServerError)
		return
	}

	// Update champion masteries by summoner
	if _, err := app.RiotAPI.UpdateChampionMasteriesBySummoner(r.Context(), updateChampion); err != nil {
		http.Error(w, "Failed to update champion masteries by summoner", http.StatusInternalServerError)
		return
	}

	// Update league entries by summoner
	if _, err := app.RiotAPI.UpdateLeagueEntriesBySummoner(r.Context(), updateLeagues); err != nil {
		http.Error(w, "Failed to update league entries by summoner", http.StatusInternalServerError)
		return
	}

	// If everything is successful
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
