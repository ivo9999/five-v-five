package main

import (
	"context"
	"errors"
	"fmt"
	"game-management-micro/data"
	proto "game-management-micro/riot"
	"net/http"
	"strconv"
	"time"

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

// Create a new game with balanced teams from a list of summoners
func (h *Handlers) InitializeGame(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Summoners []string `json:"summoners"`
		TeamRed   string   `json:"team_red"`
		TeamBlue  string   `json:"team_blue"`
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

	teamBlue := data.Team{Name: request.TeamBlue, Rating: 0, MasteryPoints: 0}
	teamRed := data.Team{Name: request.TeamRed, Rating: 0, MasteryPoints: 0}

	teamBlueID, err := data.InsertTeam(h.DB.DB, teamBlue)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	teamRedID, err := data.InsertTeam(h.DB.DB, teamRed)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	lanes := []string{"top", "jungle", "mid", "adc", "support"}

	for i, summonerName := range grpcResponse.Team1 {
		summoner := data.Summoner{Name: summonerName, Role: lanes[i%5], TeamID: teamBlueID}
		if _, err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	for i, summonerName := range grpcResponse.Team2 {
		summoner := data.Summoner{Name: summonerName, Role: lanes[i%5], TeamID: teamRedID}
		if _, err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	game := data.GameWithID{
		TeamBlue: teamBlueID,
		TeamRed:  teamRedID,
		Winner:   "",
		Date:     time.Now().Format("2006-01-02"),
	}

	gameID, err := data.InsertGame(h.DB.DB, game)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusCreated, map[string]int{"game_id": gameID})
}

// Get a game by its ID, including the details of both teams and summoners
func (h *Handlers) GetGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if game == nil {
		h.Config.errorJSON(w, errors.New("game not found"), http.StatusNotFound)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, game)
}

// Shuffle the teams in a game
func (h *Handlers) ShuffleTeams(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if game == nil {
		h.Config.errorJSON(w, errors.New("game not found"), http.StatusNotFound)
		return
	}

	summonerNames := getSummonerNames(game)

	grpcRequest := &proto.GetTeamsRequest{
		Summoners: summonerNames,
	}

	fmt.Println(grpcRequest)
	fmt.Println(summonerNames)

	grpcResponse, err := h.RiotAPI.GetTeams(r.Context(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := data.ClearTeams(h.DB.DB, game.ID); err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	fmt.Println(grpcResponse)

	// Assign new teams
	lanes := []string{"top", "jungle", "mid", "adc", "support"}

	for i, summonerName := range grpcResponse.Team1 {
		summoner := data.Summoner{Name: summonerName, Role: lanes[i%5], TeamID: game.TeamBlue.ID}
		if _, err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	for i, summonerName := range grpcResponse.Team2 {
		summoner := data.Summoner{Name: summonerName, Role: lanes[i%5], TeamID: game.TeamRed.ID}
		if _, err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	game, err = data.GetGame(h.DB.DB, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Teams shuffled successfully",
		Data:    game,
	})
}

// Swap summoners between teams in a game
func (h *Handlers) SwapSummoners(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var request struct {
		Summoner1 string `json:"summoner1"`
		Summoner2 string `json:"summoner2"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	err = data.SwapSummoners(h.DB.DB, gameID, request.Summoner1, request.Summoner2)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Summoners swapped successfully",
		Data:    game,
	})
}

// GetChampions handles the request to get balanced champions for a game.
func (h *Handlers) GetChampions(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if game == nil {
		h.Config.errorJSON(w, errors.New("game not found"), http.StatusNotFound)
		return
	}

	team1, team2 := getSummonerInfo(game)

	// Construct the gRPC request
	grpcRequest := &proto.GetChampionsByTeamsRequest{
		Team1: make([]*proto.SummonerLane, len(team1)),
		Team2: make([]*proto.SummonerLane, len(team2)),
	}

	for i, summoner := range team1 {
		grpcRequest.Team1[i] = &proto.SummonerLane{
			SummonerName: summoner.SummonerName,
			Lane:         summoner.Lane,
		}
	}

	for i, summoner := range team2 {
		grpcRequest.Team2[i] = &proto.SummonerLane{
			SummonerName: summoner.SummonerName,
			Lane:         summoner.Lane,
		}
	}

	grpcResponse, err := h.RiotAPI.GetChampionsByTeams(context.Background(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Update the summoners with the received champions
	var updatedSummoners []data.Summoner
	for _, champion := range grpcResponse.Team1Champions {
		updatedSummoners = append(updatedSummoners, data.Summoner{
			Name:     champion.SummonerName,
			Role:     champion.Lane,
			Champion: champion.ChampionName,
			TeamID:   game.TeamBlue.ID, // Assuming Team1 is TeamBlue
		})
	}

	for _, champion := range grpcResponse.Team2Champions {
		updatedSummoners = append(updatedSummoners, data.Summoner{
			Name:     champion.SummonerName,
			Role:     champion.Lane,
			Champion: champion.ChampionName,
			TeamID:   game.TeamRed.ID, // Assuming Team2 is TeamRed
		})
	}

	fmt.Printf("Updated Summoners: %+v\n", updatedSummoners)

	err = data.UpdateSummoners(h.DB.DB, updatedSummoners)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	game, _ = data.GetGame(h.DB.DB, gameID)

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Balanced champions retrieved and updated successfully",
		Data:    game,
	})
}

// Get a new champion for a summoner in a game
func (h *Handlers) GetNewChampion(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var request struct {
		SummonerName string `json:"summoner_name"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	// Retrieve the summoner's lane from the database
	lane, err := data.GetSummonerLane(h.DB.DB, request.SummonerName, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Construct the gRPC request to get a new champion for the summoner and lane
	grpcRequest := &proto.ChampionBySummonerAndLaneRequest{
		SummonerId: request.SummonerName,
		Lane:       lane,
	}

	grpcResponse, err := h.RiotAPI.GetChampionBySummonerAndLane(context.Background(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Update the summoner with the new champion
	err = data.UpdateSummonerChampion(h.DB.DB, request.SummonerName, grpcResponse.Champion)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	game, err := data.GetGame(h.DB.DB, gameID)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "New champion assigned successfully",
		Data:    game,
	})
}

// Set the winner of a game
func (h *Handlers) SetWinner(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(chi.URLParam(r, "gameId"))
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

	err = data.SetWinner(h.DB.DB, gameID, request.Winner)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Winner set successfully",
	})
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
