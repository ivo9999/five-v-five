package main

import (
	"context"
	"fmt"
	"game-management-micro/data"
	"game-management-micro/proto"
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

func (h *Handlers) InitializeGame(w http.ResponseWriter, r *http.Request) {
	var game data.Game
	if err := h.Config.readJSON(w, r, &game); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	fmt.Println(game)

	for i := range game.Team1 {
		if err := h.populateSummonerData(&game.Team1[i]); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	for i := range game.Team2 {
		if err := h.populateSummonerData(&game.Team2[i]); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	gameId, err := data.CreateGame(h.DB.DB, game)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusCreated, map[string]int{"game_id": gameId})
}

func (h *Handlers) AddPlayerToTeam(w http.ResponseWriter, r *http.Request) {
	gameId, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var summoner data.Summoner
	if err := h.Config.readJSON(w, r, &summoner); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	if err := h.populateSummonerData(&summoner); err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	toTeam := chi.URLParam(r, "team")
	err = data.MoveSummonerTeam(h.DB.DB, gameId, summoner.SummonerName, "", toTeam)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Player added to team",
	})
}

func (h *Handlers) MovePlayer(w http.ResponseWriter, r *http.Request) {
	gameId, err := strconv.Atoi(chi.URLParam(r, "gameId"))
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var request struct {
		SummonerName string `json:"summoner_name"`
		FromTeam     string `json:"from_team"`
		ToTeam       string `json:"to_team"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	err = data.MoveSummonerTeam(h.DB.DB, gameId, request.SummonerName, request.FromTeam, request.ToTeam)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	h.Config.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Player moved between teams",
	})
}

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

	summoner := data.Summoner{
		SummonerName: request.SummonerName,
		Lane:         request.Lane,
		ChampionName: grpcResponse.Champion,
	}

	if err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: "Champion retrieved and stored successfully",
		Data:    grpcResponse.Champion,
	}

	h.Config.writeJSON(w, http.StatusOK, response)
}

func (h *Handlers) GetGameChampions(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Team1 []proto.SummonerLane `json:"team1"`
		Team2 []proto.SummonerLane `json:"team2"`
	}

	if err := h.Config.readJSON(w, r, &request); err != nil {
		h.Config.errorJSON(w, err)
		return
	}

	team1 := make([]*proto.SummonerLane, len(request.Team1))
	for i, lane := range request.Team1 {
		team1[i] = &lane
	}

	team2 := make([]*proto.SummonerLane, len(request.Team2))
	for i, lane := range request.Team2 {
		team2[i] = &lane
	}

	grpcRequest := &proto.GetChampionsByTeamsRequest{
		Team1: team1,
		Team2: team2,
	}

	grpcResponse, err := h.RiotAPI.GetChampionsByTeams(r.Context(), grpcRequest)
	if err != nil {
		h.Config.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	for _, champ := range grpcResponse.Team1Champions {
		summoner := data.Summoner{
			SummonerName:   champ.SummonerName,
			Lane:           champ.Lane,
			ChampionName:   champ.ChampionName,
			ChampionPoints: int(champ.ChampionPoints),
		}

		if err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	for _, champ := range grpcResponse.Team2Champions {
		summoner := data.Summoner{
			SummonerName:   champ.SummonerName,
			Lane:           champ.Lane,
			ChampionName:   champ.ChampionName,
			ChampionPoints: int(champ.ChampionPoints),
		}

		if err := data.InsertSummoner(h.DB.DB, summoner); err != nil {
			h.Config.errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	response := jsonResponse{
		Error:   false,
		Message: "Champions retrieved and stored successfully",
		Data:    grpcResponse,
	}

	h.Config.writeJSON(w, http.StatusOK, response)
}

func (h *Handlers) populateSummonerData(summoner *data.Summoner) error {
	grpcRequest := &proto.ChampionBySummonerAndLaneRequest{
		SummonerId: summoner.SummonerName,
		Lane:       summoner.Lane,
	}
	fmt.Println(grpcRequest)

	grpcResponse, err := h.RiotAPI.GetChampionBySummonerAndLane(context.Background(), grpcRequest)
	if err != nil {
		return err
	}

	fmt.Println(grpcResponse)

	summoner.ChampionName = grpcResponse.Champion

	return nil
}
