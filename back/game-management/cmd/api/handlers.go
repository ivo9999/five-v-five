package main

import (
	"errors"
	"game-management-micro/data"
	proto "game-management-micro/riot"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	DB      *data.DB
	RiotAPI proto.RiotAPIServiceClient
	Config  *Config
}

func NewHandlers(db *data.DB, riotAPI proto.RiotAPIServiceClient, config *Config) *Handlers {
	return &Handlers{
		DB:      db,
		RiotAPI: riotAPI,
		Config:  config,
	}
}

// Create balanced teams from a list of summoners
func (h *Handlers) CreateTeams(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Summoners []string `json:"summoners"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	grpcRequest := &proto.GetTeamsRequest{
		Summoners: request.Summoners,
	}

	grpcResponse, err := h.RiotAPI.GetTeams(r.Context(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Teams created successfully",
		Data:    grpcResponse,
	}

	h.Config.writeJSON(w, http.StatusOK, response)
}

// Initialize a new game with given teams
func (h *Handlers) InitializeGame(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Team1 struct {
			TeamName  string   `json:"team_name"`
			Summoners []string `json:"summoners"`
		} `json:"team1"`
		Team2 struct {
			TeamName  string   `json:"team_name"`
			Summoners []string `json:"summoners"`
		} `json:"team2"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	// Create the game
	game := data.Game{
		Mode:  "standard", // or any other mode
		State: "new",
	}

	gameId, err := data.CreateGame(h.DB.DB, game)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Assign the game ID to the teams and determine lanes based on summoner index
	team1 := data.Team{GameID: gameId, TeamName: request.Team1.TeamName, Summoners: []data.Summoner{}}
	team2 := data.Team{GameID: gameId, TeamName: request.Team2.TeamName, Summoners: []data.Summoner{}}

	lanes := []string{"top", "jungle", "mid", "adc", "support"}

	for i, summonerName := range request.Team1.Summoners {
		summoner := data.Summoner{SummonerName: summonerName, Lane: lanes[i]}
		team1.Summoners = append(team1.Summoners, summoner)

		// Insert summoner into the summoners table
		if err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	for i, summonerName := range request.Team2.Summoners {
		summoner := data.Summoner{SummonerName: summonerName, Lane: lanes[i]}
		team2.Summoners = append(team2.Summoners, summoner)

		// Insert summoner into the summoners table
		if err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	// Create teams and insert summoners
	team1ID, err := data.CreateTeam(h.DB.DB, team1)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	team2ID, err := data.CreateTeam(h.DB.DB, team2)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Insert into game_teams
	if err := data.InsertGameTeam(h.DB.DB, gameId, team1ID); err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	if err := data.InsertGameTeam(h.DB.DB, gameId, team2ID); err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusCreated, map[string]int{"game_id": gameId})
}

// Get a game by its ID, including the details of both teams
func (h *Handlers) GetGame(w http.ResponseWriter, r *http.Request) {
	gameId, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameId)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if game == nil {
		h.Config.errorJSON(w, errors.New("Game not found"), http.StatusNotFound)
		return
	}

	teams, err := data.GetTeamsByGameID(h.DB.DB, gameId)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := struct {
		ID     int         `json:"id"`
		Mode   string      `json:"mode"`
		Winner string      `json:"winner"`
		State  string      `json:"state"`
		Teams  *data.Teams `json:"teams"`
	}{
		ID:     game.ID,
		Mode:   game.Mode,
		Winner: game.Winner,
		State:  game.State,
		Teams:  teams,
	}

	h.Config.writeJSON(w, http.StatusOK, response)
}

// Get champions for a game
func (h *Handlers) GetChamps(w http.ResponseWriter, r *http.Request) {
	gameId, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	teams, err := data.GetTeamsByGameID(h.DB.DB, gameId)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	var team1Summoners []*proto.SummonerLane
	var team2Summoners []*proto.SummonerLane

	for _, summoner := range teams.Team1.Summoners {
		team1Summoners = append(team1Summoners, &proto.SummonerLane{
			SummonerName: summoner.SummonerName,
			Lane:         summoner.Lane,
		})
	}

	for _, summoner := range teams.Team2.Summoners {
		team2Summoners = append(team2Summoners, &proto.SummonerLane{
			SummonerName: summoner.SummonerName,
			Lane:         summoner.Lane,
		})
	}

	grpcRequest := &proto.GetChampionsByTeamsRequest{
		Team1: team1Summoners,
		Team2: team2Summoners,
	}

	grpcResponse, err := h.RiotAPI.GetChampionsByTeams(r.Context(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Champions retrieved successfully",
		Data:    grpcResponse,
	}

	h.Config.writeJSON(w, http.StatusOK, response)
}

// Get a champion for a player and lane
func (h *Handlers) GetChampion(w http.ResponseWriter, r *http.Request) {
	var request struct {
		SummonerName string `json:"summoner_name"`
		Lane         string `json:"lane"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	grpcRequest := &proto.ChampionBySummonerAndLaneRequest{
		SummonerId: request.SummonerName,
		Lane:       request.Lane,
	}

	grpcResponse, err := h.RiotAPI.GetChampionBySummonerAndLane(r.Context(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Champion retrieved successfully",
		Data:    grpcResponse,
	}

	h.Config.writeJSON(w, http.StatusOK, response)
}

// Set the winner of a game
func (h *Handlers) SetWinner(w http.ResponseWriter, r *http.Request) {
	gameId, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var request struct {
		Winner string `json:"winner"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	err = data.SetWinner(h.DB.DB, gameId, request.Winner)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Winner set successfully",
	})
}

// Move a player between teams
func (h *Handlers) MovePlayer(w http.ResponseWriter, r *http.Request) {
	gameId, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var request struct {
		SummonerName string `json:"summoner_name"`
		ToTeam       string `json:"to_team"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	fromTeamID, err := data.GetTeamIDBySummoner(h.DB.DB, request.SummonerName)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	toTeamID, err := data.GetTeamIDByGameAndName(h.DB.DB, gameId, request.ToTeam)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = data.MoveSummonerTeam(h.DB.DB, fromTeamID, toTeamID, request.SummonerName)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameId)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, game)
}

// Drop the database
func (h *Handlers) DropDB(w http.ResponseWriter, r *http.Request) {
	err := data.DropDatabase(h.DB.DB)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Database dropped successfully",
	})
}
