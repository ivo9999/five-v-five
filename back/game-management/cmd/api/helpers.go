package main

import (
	"encoding/json"
	"errors"
	"game-management-micro/data"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJSON tries to read the body of a request and converts it into JSON
func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// writeJSON takes a response status code and arbitrary data and writes a json response to the client
func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}

func getSummonerNames(game *data.Game) []string {
	var summonerNames []string

	// Extract names from TeamBlue
	for _, summoner := range game.TeamBlue.Summoners {
		summonerNames = append(summonerNames, summoner.Name)
	}

	// Extract names from TeamRed
	for _, summoner := range game.TeamRed.Summoners {
		summonerNames = append(summonerNames, summoner.Name)
	}

	return summonerNames
}

func getSummonerNamesFromTeam(team *data.Team) []string {
	var summonerNames []string
	for _, summoner := range team.Summoners {
		summonerNames = append(summonerNames, summoner.Name)
	}
	return summonerNames
}

type SummonerInfo struct {
	Lane         string `json:"lane"`
	SummonerName string `json:"summonerName"`
}

func getSummonerInfo(game *data.Game) ([]SummonerInfo, []SummonerInfo) {
	var team1Info []SummonerInfo
	var team2Info []SummonerInfo

	for _, summoner := range game.TeamBlue.Summoners {
		team1Info = append(team1Info, SummonerInfo{Lane: summoner.Role, SummonerName: summoner.Name})
	}

	for _, summoner := range game.TeamRed.Summoners {
		team2Info = append(team2Info, SummonerInfo{Lane: summoner.Role, SummonerName: summoner.Name})
	}

	return team1Info, team2Info
}
