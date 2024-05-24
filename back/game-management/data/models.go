package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Models struct {
	Summoner Summoner
	Game     Game
	Team     Team
}

func NewModels(db *DB) Models {
	return Models{
		Summoner: Summoner{},
		Game:     Game{},
		Team:     Team{},
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

type Team struct {
	TeamName  string     `json:"team_name"`
	Summoners []Summoner `json:"summoners"`
	ID        int        `json:"id"`
	GameID    int        `json:"game_id"`
}

type Game struct {
	Mode   string `json:"mode"`
	Winner string `json:"winner"`
	State  string `json:"state"`
	ID     int    `json:"id"`
}

type Teams struct {
	Team1 Team
	Team2 Team
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
		state VARCHAR(50)
	);`); err != nil {
		log.Fatalf("Error creating games table: %v", err)
		return err
	}

	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS teams (
		id SERIAL PRIMARY KEY,
		game_id INT,
		team_name VARCHAR(50),
		FOREIGN KEY (game_id) REFERENCES games(id)
	);`); err != nil {
		log.Fatalf("Error creating teams table: %v", err)
		return err
	}

	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS team_summoners (
		team_id INT,
		summoner_name VARCHAR(255),
		lane VARCHAR(50),
		champion_name VARCHAR(255),
		division VARCHAR(50),
		champion_points INT,
		FOREIGN KEY (team_id) REFERENCES teams(id),
		FOREIGN KEY (summoner_name, champion_name) REFERENCES summoners(summoner_name, champion_name),
		PRIMARY KEY (team_id, summoner_name, champion_name)
	);`); err != nil {
		log.Fatalf("Error creating team_summoners table: %v", err)
		return err
	}

	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS game_teams (
		game_id INT,
		team_id INT,
		FOREIGN KEY (game_id) REFERENCES games(id),
		FOREIGN KEY (team_id) REFERENCES teams(id),
		PRIMARY KEY (game_id, team_id)
	);`); err != nil {
		log.Fatalf("Error creating game_teams table: %v", err)
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

	query := `
	INSERT INTO games (mode, winner, state)
	VALUES ($1, $2, $3)
	RETURNING id;
	`
	var gameId int
	err := db.QueryRowContext(ctx, query, game.Mode, game.Winner, game.State).Scan(&gameId)
	if err != nil {
		return 0, err
	}
	return gameId, nil
}

func CreateTeam(db *sql.DB, team Team) (int, error) {
	ctx := context.Background()

	query := `
	INSERT INTO teams (team_name, game_id)
	VALUES ($1, $2)
	RETURNING id;
	`
	var teamId int
	err := db.QueryRowContext(ctx, query, team.TeamName, team.GameID).Scan(&teamId)
	if err != nil {
		return 0, err
	}

	for _, summoner := range team.Summoners {
		if err := InsertSummonerToTeam(db, teamId, summoner); err != nil {
			return 0, err
		}
	}

	return teamId, nil
}

func InsertSummonerToTeam(db *sql.DB, teamId int, summoner Summoner) error {
	ctx := context.Background()

	query := `
	INSERT INTO team_summoners (team_id, summoner_name, lane, champion_name, division, champion_points)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (team_id, summoner_name, champion_name) DO UPDATE SET
	lane = EXCLUDED.lane,
	division = EXCLUDED.division,
	champion_points = EXCLUDED.champion_points;
	`
	_, err := db.ExecContext(ctx, query, teamId, summoner.SummonerName, summoner.Lane, summoner.ChampionName, summoner.Division, summoner.ChampionPoints)
	return err
}

func InsertGameTeam(db *sql.DB, gameId int, teamId int) error {
	ctx := context.Background()

	query := `
	INSERT INTO game_teams (game_id, team_id)
	VALUES ($1, $2);
	`
	_, err := db.ExecContext(ctx, query, gameId, teamId)
	return err
}

func GetTeamIDByGameAndName(db *sql.DB, gameId int, teamName string) (int, error) {
	ctx := context.Background()

	query := `
	SELECT id FROM teams WHERE team_name = $1 AND game_id = $2;
	`
	var teamId int
	err := db.QueryRowContext(ctx, query, teamName, gameId).Scan(&teamId)
	if err != nil {
		return 0, err
	}
	return teamId, nil
}

func GetTeamIDBySummoner(db *sql.DB, summonerName string) (int, error) {
	ctx := context.Background()

	query := `
	SELECT team_id FROM team_summoners WHERE summoner_name = $1;
	`
	var teamId int
	err := db.QueryRowContext(ctx, query, summonerName).Scan(&teamId)
	if err != nil {
		return 0, err
	}
	return teamId, nil
}

func MoveSummonerTeam(db *sql.DB, fromTeamId int, toTeamId int, summonerName string) error {
	ctx := context.Background()

	query := `
	UPDATE team_summoners SET team_id = $1 WHERE team_id = $2 AND summoner_name = $3;
	`
	_, err := db.ExecContext(ctx, query, toTeamId, fromTeamId, summonerName)
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

func GetGame(db *sql.DB, gameId int) (*Game, error) {
	ctx := context.Background()

	query := `
	SELECT id, mode, winner, state
	FROM games
	WHERE id = $1;
	`
	row := db.QueryRowContext(ctx, query, gameId)

	var game Game
	err := row.Scan(&game.ID, &game.Mode, &game.Winner, &game.State)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result is not an error, simply no game found
		}
		return nil, err
	}
	return &game, nil
}

func GetTeamsByGameID(db *sql.DB, gameId int) (*Teams, error) {
	ctx := context.Background()

	query := `
	SELECT t.id, t.team_name, ts.summoner_name, ts.lane, ts.champion_name, ts.division, ts.champion_points
	FROM teams t
	JOIN team_summoners ts ON t.id = ts.team_id
	WHERE t.game_id = $1;
	`
	rows, err := db.QueryContext(ctx, query, gameId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := &Teams{
		Team1: Team{Summoners: []Summoner{}},
		Team2: Team{Summoners: []Summoner{}},
	}
	for rows.Next() {
		var teamName, summonerName, lane, championName, division string
		var championPoints, teamID int
		if err := rows.Scan(&teamID, &teamName, &summonerName, &lane, &championName, &division, &championPoints); err != nil {
			return nil, err
		}
		summoner := Summoner{
			SummonerName:   summonerName,
			Lane:           lane,
			ChampionName:   championName,
			Division:       division,
			ChampionPoints: championPoints,
		}
		if teamName == teams.Team1.TeamName {
			teams.Team1.Summoners = append(teams.Team1.Summoners, summoner)
		} else {
			teams.Team2.Summoners = append(teams.Team2.Summoners, summoner)
		}
	}
	return teams, nil
}

func DropDatabase(db *sql.DB) error {
	ctx := context.Background()

	tables := []string{"game_teams", "team_summoners", "teams", "games", "summoners"}

	for _, table := range tables {
		if _, err := db.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", table)); err != nil {
			log.Fatalf("Error dropping table %s: %v", table, err)
			return err
		}
	}

	log.Println("Database dropped successfully.")
	return nil
}
