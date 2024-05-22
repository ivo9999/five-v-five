package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"riot-micro/cmd/data"
)

const (
	baseURL = "https://europe.api.riotgames.com/"
	euwURL  = "https://euw1.api.riotgames.com/"
)

func FetchAccountByName(apiKey, summonerName, tag string) (*data.Summoner, error) {
	url := fmt.Sprintf("%sriot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", baseURL, summonerName, tag, apiKey)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var summoner data.Summoner
	if err := json.NewDecoder(resp.Body).Decode(&summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func FetchSummonerByName(apiKey, summonerName string) (*data.Summoner, error) {
	url := fmt.Sprintf("%slol/summoner/v4/summoners/by-puuid/%s?api_key=%s", euwURL, summonerName, apiKey)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var summoner data.Summoner
	if err := json.NewDecoder(resp.Body).Decode(&summoner); err != nil {
		return nil, err
	}
	return &summoner, nil
}

func FetchChampionMasteries(apiKey, summonerID string) ([]data.ChampionMastery, error) {
	url := fmt.Sprintf("%slol/champion-mastery/v4/champion-masteries/by-puuid/%s?api_key=%s", euwURL, summonerID, apiKey)
	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var masteries []data.ChampionMastery
	if err := json.NewDecoder(resp.Body).Decode(&masteries); err != nil {
		return nil, err
	}
	return masteries, nil
}

func FetchLeagueEntries(apiKey, summonerID string) ([]data.LeagueEntry, error) {
	url := fmt.Sprintf("%slol/league/v4/entries/by-summoner/%s?api_key=%s", euwURL, summonerID, apiKey)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entries []data.LeagueEntry
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, err
	}
	return entries, nil
}
