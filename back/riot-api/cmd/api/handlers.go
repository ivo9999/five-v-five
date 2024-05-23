package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"riot-micro/cmd/data"
	"riot-micro/riot"

	"github.com/joho/godotenv"
)

func (s *RiotAPIServer) GetSummonerByName(ctx context.Context, req *riot.SummonerByNameRequest) (*riot.SummonerResponse, error) {
	summoner, err := data.GetSummoner(s.db, ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &riot.SummonerResponse{
		Id:            summoner.ID,
		AccountId:     summoner.AccountID,
		Puuid:         summoner.PUUID,
		Name:          summoner.Name,
		Tag:           summoner.Tag,
		SummonerLevel: int32(summoner.SummonerLevel),
	}, nil
}

func (s *RiotAPIServer) GetChampionBySummonerAndLane(ctx context.Context, req *riot.ChampionBySummonerAndLaneRequest) (*riot.ChampionBySummonerAndLaneResponse, error) {
	champ, _, err := data.GetRandomChampionForLane(s.db, req.SummonerId, req.Lane)
	if err != nil {
		return nil, err
	}

	return &riot.ChampionBySummonerAndLaneResponse{
		Champion: champ,
	}, nil
}

func (s *RiotAPIServer) GetChampionsByTeams(ctx context.Context, req *riot.GetChampionsByTeamsRequest) (*riot.GetChampionsByTeamsResponse, error) {
	var team1Champions []*riot.SummonerChampion
	var team2Champions []*riot.SummonerChampion

	var bestTeam1Champions []*riot.SummonerChampion
	var bestTeam2Champions []*riot.SummonerChampion
	minDifference := 999999.0

	for i := 0; i < 20; i++ {
		team1Champions = []*riot.SummonerChampion{}
		team2Champions = []*riot.SummonerChampion{}
		var totalPointsTeam1 int32 = 0
		var totalPointsTeam2 int32 = 0
		success := true

		for _, sl := range req.Team1 {
			champ, points, err := data.GetRandomChampionForLane(s.db, sl.SummonerName, sl.Lane)
			if err != nil {
				success = false
				break
			}
			team1Champions = append(team1Champions, &riot.SummonerChampion{
				SummonerName:   sl.SummonerName,
				Lane:           sl.Lane,
				ChampionName:   champ,
				ChampionPoints: points,
			})
			totalPointsTeam1 += points
		}

		for _, sl := range req.Team2 {
			champ, points, err := data.GetRandomChampionForLane(s.db, sl.SummonerName, sl.Lane)
			if err != nil {
				success = false
				break
			}
			team2Champions = append(team2Champions, &riot.SummonerChampion{
				SummonerName:   sl.SummonerName,
				Lane:           sl.Lane,
				ChampionName:   champ,
				ChampionPoints: points,
			})
			totalPointsTeam2 += points
		}

		if success {
			difference := math.Abs(float64(totalPointsTeam1) - float64(totalPointsTeam2))
			if difference < minDifference {
				minDifference = difference
				bestTeam1Champions = append([]*riot.SummonerChampion{}, team1Champions...)
				bestTeam2Champions = append([]*riot.SummonerChampion{}, team2Champions...)
			}

			if float64(totalPointsTeam1)/float64(totalPointsTeam2) <= 1.25 && float64(totalPointsTeam2)/float64(totalPointsTeam1) <= 1.25 {
				break
			}
		}
	}

	return &riot.GetChampionsByTeamsResponse{
		Team1Champions: bestTeam1Champions,
		Team2Champions: bestTeam2Champions,
	}, nil
}

func (s *RiotAPIServer) UpdateSummonerByName(ctx context.Context, req *riot.SummonerByNameRequest) (*riot.SummonerResponse, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		fmt.Println("API key not set in the environment variables")
	}

	account, err := FetchAccountByName(apiKey, req.Name, req.Tag)
	if err != nil {
		return nil, err
	}

	summoner, err := FetchSummonerByName(apiKey, account.PUUID)
	if err != nil {
		return nil, err
	}

	if err := data.InsertSummoner(s.db, ctx, data.Summoner{
		ID:            summoner.ID,
		AccountID:     summoner.AccountID,
		PUUID:         summoner.PUUID,
		Name:          req.Name,
		Tag:           req.Tag,
		SummonerLevel: summoner.SummonerLevel,
	}); err != nil {
		return nil, err
	}

	return &riot.SummonerResponse{
		Id:            summoner.ID,
		AccountId:     summoner.AccountID,
		Puuid:         summoner.PUUID,
		Name:          req.Name,
		Tag:           req.Tag,
		SummonerLevel: int32(summoner.SummonerLevel),
	}, nil
}

func (s *RiotAPIServer) GetLeagueEntriesBySummoner(ctx context.Context, req *riot.LeagueEntriesRequest) (*riot.LeagueEntriesResponse, error) {
	summoner, err := data.GetSummoner(s.db, ctx, req.SummonerId)
	if err != nil {
		return nil, err
	}

	entries, err := data.GetLeagueEntries(s.db, ctx, summoner.ID)
	if err != nil {
		return nil, err
	}
	response := &riot.LeagueEntriesResponse{}
	for _, entry := range entries {
		grpcEntry := &riot.LeagueEntry{
			LeagueId:     entry.LeagueID,
			QueueType:    entry.QueueType,
			Tier:         entry.Tier,
			Rank:         entry.Rank,
			LeaguePoints: int32(entry.LeaguePoints),
			Wins:         int32(entry.Wins),
			Losses:       int32(entry.Losses),
			Veteran:      entry.Veteran,
			Inactive:     entry.Inactive,
			FreshBlood:   entry.FreshBlood,
			HotStreak:    entry.HotStreak,
		}
		response.Entries = append(response.Entries, grpcEntry)
	}
	return response, nil
}

func (s *RiotAPIServer) UpdateLeagueEntriesBySummoner(ctx context.Context, req *riot.LeagueEntriesRequest) (*riot.LeagueEntriesResponse, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get API key from environment variables
	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		fmt.Println("API key not set in the environment variables")
	}

	summoner, err := data.GetSummoner(s.db, ctx, req.SummonerId)
	if err != nil {
		return nil, err
	}

	entries, err := FetchLeagueEntries(apiKey, summoner.ID)
	if err != nil {
		return nil, err
	}

	var grpcEntries []*riot.LeagueEntry
	for _, entry := range entries {
		// Update the database with the fetched data
		if err := data.InsertLeagueEntry(s.db, ctx, data.LeagueEntry{
			LeagueID:     entry.LeagueID,
			SummonerID:   summoner.ID,
			QueueType:    entry.QueueType,
			Tier:         entry.Tier,
			Wins:         entry.Wins,
			Losses:       entry.Losses,
			Veteran:      entry.Veteran,
			Inactive:     entry.Inactive,
			FreshBlood:   entry.FreshBlood,
			Rank:         entry.Rank,
			LeaguePoints: entry.LeaguePoints,
		}); err != nil {
			return nil, err
		}
		// Convert data.LeagueEntry to riot.LeagueEntry and append to the grpcEntries slice
		grpcEntries = append(grpcEntries, &riot.LeagueEntry{
			LeagueId:     entry.LeagueID,
			QueueType:    entry.QueueType,
			Tier:         entry.Tier,
			Wins:         int32(entry.Wins),
			Losses:       int32(entry.Losses),
			Rank:         entry.Rank,
			Veteran:      entry.Veteran,
			Inactive:     entry.Inactive,
			FreshBlood:   entry.FreshBlood,
			LeaguePoints: int32(entry.LeaguePoints),
		})
	}

	fmt.Println(grpcEntries)
	// Return the response with the converted slice
	return &riot.LeagueEntriesResponse{
		Entries: grpcEntries,
	}, nil
}

func (s *RiotAPIServer) GetChampionMasteriesBySummoner(ctx context.Context, req *riot.ChampionMasteriesRequest) (*riot.ChampionMasteriesResponse, error) {
	summoner, err := data.GetSummoner(s.db, ctx, req.SummonerId)
	if err != nil {
		return nil, err
	}

	masteries, err := data.GetChampionMasteries(s.db, ctx, summoner.PUUID)
	if err != nil {
		return nil, err
	}
	fmt.Println(masteries)
	response := &riot.ChampionMasteriesResponse{}
	for _, mastery := range masteries {
		grpcMastery := &riot.ChampionMastery{
			ChampionId:     mastery.ChampionId,
			ChampionLevel:  int32(mastery.ChampionLevel),
			ChampionPoints: int32(mastery.ChampionPoints),
			LastPlayTime:   mastery.LastPlayTime,
			TokensEarned:   int32(mastery.TokensEarned),
			ChestGranted:   mastery.ChestGranted,
		}
		response.Masteries = append(response.Masteries, grpcMastery)
	}
	return response, nil
}

func (s *RiotAPIServer) UpdateChampionMasteriesBySummoner(ctx context.Context, req *riot.ChampionMasteriesRequest) (*riot.ChampionMasteriesResponse, error) {
	err := godotenv.Load() // Load environment variables from the .env file
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("RIOT_API_KEY") // Get API key from environment variables
	if apiKey == "" {
		fmt.Println("API key not set in the environment variables")
	}

	summoner, err := data.GetSummoner(s.db, ctx, req.SummonerId)
	if err != nil {
		return nil, err
	}

	masteries, err := FetchChampionMasteries(apiKey, summoner.PUUID) // Fetch masteries from the API
	if err != nil {
		return nil, err
	}

	var grpcMasteries []*riot.ChampionMastery
	for _, mastery := range masteries {
		// Update the database with the fetched data
		if err := data.InsertChampionMastery(s.db, ctx, data.ChampionMastery{
			PUUID:                        mastery.PUUID,
			ChampionId:                   mastery.ChampionId,
			ChampionLevel:                mastery.ChampionLevel,
			ChampionPoints:               mastery.ChampionPoints,
			LastPlayTime:                 mastery.LastPlayTime,
			ChampionPointsSinceLastLevel: mastery.ChampionPointsSinceLastLevel,
			ChampionPointsUntilNextLevel: mastery.ChampionPointsUntilNextLevel,
			TokensEarned:                 mastery.TokensEarned,
			ChestGranted:                 mastery.ChestGranted,
		}); err != nil {
			return nil, err
		}
		// Convert data.ChampionMastery to *riot.ChampionMastery for the gRPC response
		grpcMastery := &riot.ChampionMastery{
			ChampionId:     mastery.ChampionId,
			ChampionLevel:  int32(mastery.ChampionLevel),
			ChampionPoints: int32(mastery.ChampionPoints),
			LastPlayTime:   mastery.LastPlayTime,
			TokensEarned:   int32(mastery.TokensEarned),
			ChestGranted:   mastery.ChestGranted,
		}
		grpcMasteries = append(grpcMasteries, grpcMastery)
	}

	// Return the response with the converted slice
	return &riot.ChampionMasteriesResponse{
		Masteries: grpcMasteries,
	}, nil
}
