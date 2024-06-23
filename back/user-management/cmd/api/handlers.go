package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	err = UpdateUserRiot(user.LeagueName, user.LeagueTag, app)
	if err != nil {
		fmt.Println("Error updating user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (app *Config) loginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc, err := data.GetUserByUsernameLogin(app.DB.DB, req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ValidPassword(req.Password, acc.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := data.CreateJWT(&acc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := data.LoginResponse{
		Username: acc.Username,
		ID:       acc.ID,
		Token:    token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
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

	err := updateUser(username, app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If everything is successful
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func updateUser(username string, app *Config) error {
	user, err := data.GetUserByUsername(app.DB.DB, username)
	if err != nil {
		return err
	}

	fmt.Println(user)

	err = UpdateUserRiot(user.LeagueName, user.LeagueTag, app)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserRiot(leagueName string, leagueTag string, app *Config) error {
	updateSumoner := &riot.SummonerByNameRequest{
		Name: leagueName,
		Tag:  leagueTag,
	}

	updateChampion := &riot.ChampionMasteriesRequest{
		SummonerId: leagueName,
	}

	updateLeagues := &riot.LeagueEntriesRequest{
		SummonerId: leagueName,
	}

	// Update summoner by name
	if _, err := app.RiotAPI.UpdateSummonerByName(context.Background(), updateSumoner); err != nil {
		return err
	}

	// Update champion masteries by summoner
	if _, err := app.RiotAPI.UpdateChampionMasteriesBySummoner(context.Background(), updateChampion); err != nil {
		return err
	}

	// Update league entries by summoner
	if _, err := app.RiotAPI.UpdateLeagueEntriesBySummoner(context.Background(), updateLeagues); err != nil {
		return err
	}

	return nil
}
