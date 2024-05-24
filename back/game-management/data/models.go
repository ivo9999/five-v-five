package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
)

type Models struct {
	Summoner Summoner
	Game     Game
}

func NewModels(db *DB) Models {
	return Models{
		Summoner: Summoner{},
		Game:     Game{},
	}
}

type DB struct {
	*sql.DB
}

func NewDB(database *sql.DB) *DB {
	return &DB{DB: database}
}

type Summoner struct {
	SummonerName   string `json:"summoner_name"`
	Lane           string `json:"lane"`
	ChampionName   string `json:"champion_name"`
	Division       string `json:"division"`
	ChampionPoints int    `json:"champion_points"`
}

type Game struct {
	Mode   string     `json:"mode"`
	Winner string     `json:"winner"`
	State  string     `json:"state"`
	Team1  []Summoner `json:"team1"`
	Team2  []Summoner `json:"team2"`
	ID     int        `json:"id"`
}

func InitializeDatabase(db *sql.DB) error {
	ctx := context.Background()

	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS summoners (
		summoner_name VARCHAR(255),
		lane VARCHAR(50),
		champion_name VARCHAR(255),
		division VARCHAR(50),
		champion_points INT,
		PRIMARY KEY (summoner_name, champion_name)
	);`); err != nil {
		log.Fatalf("Error creating summoners table: %v", err)
		return err
	}

	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		mode VARCHAR(50),
		winner VARCHAR(255),
		state VARCHAR(50),
		team1 JSONB,
		team2 JSONB
	);`); err != nil {
		log.Fatalf("Error creating games table: %v", err)
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}

func InsertSummoner(db *sql.DB, summoner Summoner) error {
	ctx := context.Background()

	query := `
	INSERT INTO summoners (summoner_name, lane, champion_name, division, champion_points)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (summoner_name, champion_name) DO UPDATE SET
	lane = EXCLUDED.lane,
	division = EXCLUDED.division,
	champion_points = EXCLUDED.champion_points;
	`
	_, err := db.ExecContext(ctx, query, summoner.SummonerName, summoner.Lane, summoner.ChampionName, summoner.Division, summoner.ChampionPoints)
	return err
}

func GetSummoner(db *sql.DB, summonerName string, championName string) (*Summoner, error) {
	ctx := context.Background()

	query := `
	SELECT summoner_name, lane, champion_name, division, champion_points
	FROM summoners
	WHERE summoner_name = $1 AND champion_name = $2;
	`
	row := db.QueryRowContext(ctx, query, summonerName, championName)

	var summoner Summoner
	err := row.Scan(&summoner.SummonerName, &summoner.Lane, &summoner.ChampionName, &summoner.Division, &summoner.ChampionPoints)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result is not an error, simply no summoner found
		}
		return nil, err
	}
	return &summoner, nil
}

func CreateGame(db *sql.DB, game Game) (int, error) {
	ctx := context.Background()

	team1, err := json.Marshal(game.Team1)
	if err != nil {
		return 0, err
	}

	team2, err := json.Marshal(game.Team2)
	if err != nil {
		return 0, err
	}

	query := `
	INSERT INTO games (mode, winner, state, team1, team2)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`
	var gameId int
	err = db.QueryRowContext(ctx, query, game.Mode, game.Winner, game.State, team1, team2).Scan(&gameId)
	if err != nil {
		return 0, err
	}
	return gameId, nil
}

func MoveSummonerTeam(db *sql.DB, gameId int, summonerName string, fromTeam string, toTeam string) error {
	ctx := context.Background()

	var team1, team2 []Summoner

	query := `
	SELECT team1, team2
	FROM games
	WHERE id = $1;
	`
	row := db.QueryRowContext(ctx, query, gameId)
	err := row.Scan(&team1, &team2)
	if err != nil {
		return err
	}

	var summoner Summoner
	for i, s := range team1 {
		if s.SummonerName == summonerName {
			summoner = s
			team1 = append(team1[:i], team1[i+1:]...)
			break
		}
	}
	if summoner.SummonerName == "" {
		for i, s := range team2 {
			if s.SummonerName == summonerName {
				summoner = s
				team2 = append(team2[:i], team2[i+1:]...)
				break
			}
		}
	}

	if summoner.SummonerName == "" {
		return errors.New("summoner not found in any team")
	}

	if toTeam == "team1" {
		team1 = append(team1, summoner)
	} else {
		team2 = append(team2, summoner)
	}

	team1JSON, err := json.Marshal(team1)
	if err != nil {
		return err
	}

	team2JSON, err := json.Marshal(team2)
	if err != nil {
		return err
	}

	updateQuery := `
	UPDATE games
	SET team1 = $1, team2 = $2
	WHERE id = $3;
	`
	_, err = db.ExecContext(ctx, updateQuery, team1JSON, team2JSON, gameId)
	return err
}

func SetWinner(db *sql.DB, gameId int, winner string) error {
	ctx := context.Background()

	query := `
	UPDATE games
	SET winner = $1
	WHERE id = $2;
	`
	_, err := db.ExecContext(ctx, query, winner, gameId)
	return err
}
