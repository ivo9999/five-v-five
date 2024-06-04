package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	*sql.DB
}

func NewDB(database *sql.DB) *DB {
	return &DB{DB: database}
}

type Models struct {
	Summoner
	Team
	Game
}

func NewModels(db *DB) Models {
	return Models{
		Summoner: Summoner{},
		Team:     Team{},
		Game:     Game{},
	}
}

type Summoner struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Champion string `json:"champion"`
	ID       int    `json:"id"`
	TeamID   int    `json:"team_id"`
}

type Team struct {
	Name          string     `json:"name"`
	Summoners     []Summoner `json:"summoners"`
	ID            int        `json:"id"`
	Rating        int        `json:"rating"`
	MasteryPoints int        `json:"mastery_points"`
}

type Game struct {
	Winner   string `json:"winner"`
	Date     string `json:"date"`
	TeamBlue Team   `json:"team_blue"`
	TeamRed  Team   `json:"team_red"`
	ID       int    `json:"id"`
}

type GameWithID struct {
	Winner   string `json:"winner"`
	Date     string `json:"date"`
	TeamBlue int    `json:"team_blue"`
	TeamRed  int    `json:"team_red"`
	ID       int    `json:"id"`
}

// InitializeDatabase initializes the database schema.
func InitializeDatabase(db *sql.DB) error {
	ctx := context.Background()

	// Create Teams table
	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS teams (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		rating INT,
		mastery_points INT
	);`); err != nil {
		log.Fatalf("Error creating teams table: %v", err)
		return err
	}

	// Create Summoners table
	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS summoners (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		role VARCHAR(50),
		champion VARCHAR(255),
		team_id INT,
		FOREIGN KEY (team_id) REFERENCES teams(id)
	);`); err != nil {
		log.Fatalf("Error creating summoners table: %v", err)
		return err
	}

	// Create Games table
	if _, err := db.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		team_blue INT,
		team_red INT,
		winner VARCHAR(255),
		date VARCHAR(50),
		FOREIGN KEY (team_blue) REFERENCES teams(id),
		FOREIGN KEY (team_red) REFERENCES teams(id)
	);`); err != nil {
		log.Fatalf("Error creating games table: %v", err)
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}

// InsertSummoner inserts a new summoner into the database.
func InsertSummoner(db *sql.DB, summoner Summoner) (int, error) {
	ctx := context.Background()

	query := `
	INSERT INTO summoners (name, role, champion, team_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	var summonerID int
	err := db.QueryRowContext(ctx, query, summoner.Name, summoner.Role, summoner.Champion, summoner.TeamID).Scan(&summonerID)
	return summonerID, err
}

// InsertTeam inserts a new team into the database.
func InsertTeam(db *sql.DB, team Team) (int, error) {
	ctx := context.Background()

	query := `
	INSERT INTO teams (name, rating, mastery_points)
	VALUES ($1, $2, $3)
	RETURNING id;
	`
	var teamID int
	err := db.QueryRowContext(ctx, query, team.Name, team.Rating, team.MasteryPoints).Scan(&teamID)
	return teamID, err
}

// InsertGame inserts a new game into the database.
func InsertGame(db *sql.DB, game GameWithID) (int, error) {
	ctx := context.Background()

	query := `
	INSERT INTO games (team_blue, team_red, winner, date)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	var gameID int
	err := db.QueryRowContext(ctx, query, game.TeamBlue, game.TeamRed, "", game.Date).Scan(&gameID)
	return gameID, err
}

// UpdateTeam updates a team in the database
func UpdateTeam(db *sql.DB, team Team, rating, mastery int) error {
	ctx := context.Background()

	query := `
  UPDATE teams
  SET name = $1, rating = $2, mastery_points = $3
  WHERE id = $4;
  `
	_, err := db.ExecContext(ctx, query, team.Name, rating, mastery, team.ID)
	return err
}

// GetGame retrieves a game by its ID from the database.
func GetGame(db *sql.DB, gameID int) (*Game, error) {
	ctx := context.Background()

	query := `
	SELECT id, team_blue, team_red, winner, date
	FROM games
	WHERE id = $1;
	`
	row := db.QueryRowContext(ctx, query, gameID)

	var game Game
	err := row.Scan(&game.ID, &game.TeamBlue.ID, &game.TeamRed.ID, &game.Winner, &game.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result is not an error, simply no game found
		}
		return nil, err
	}

	// Fetch teams and their summoners
	game.TeamBlue, err = GetTeamFull(db, game.TeamBlue.ID)
	if err != nil {
		return nil, err
	}

	game.TeamRed, err = GetTeamFull(db, game.TeamRed.ID)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// GetTeamFull retrieves a team and its summoners by the team ID.
func GetTeamFull(db *sql.DB, teamID int) (Team, error) {
	ctx := context.Background()

	var team Team
	err := db.QueryRowContext(ctx, "SELECT id, name, rating, mastery_points FROM teams WHERE id = $1", teamID).Scan(&team.ID, &team.Name, &team.Rating, &team.MasteryPoints)
	if err != nil {
		return team, err
	}

	rows, err := db.QueryContext(ctx, "SELECT id, name, role, champion, team_id FROM summoners WHERE team_id = $1", teamID)
	if err != nil {
		return team, err
	}
	defer rows.Close()

	for rows.Next() {
		var summoner Summoner
		if err := rows.Scan(&summoner.ID, &summoner.Name, &summoner.Role, &summoner.Champion, &summoner.TeamID); err != nil {
			return team, err
		}
		team.Summoners = append(team.Summoners, summoner)
	}

	return team, nil
}

// UpdateSummoners updates the champion for a list of summoners in the database.
func UpdateSummoners(db *sql.DB, summoners []Summoner) error {
	ctx := context.Background()

	query := `
	UPDATE summoners
	SET champion = $1
	WHERE name = $2 AND team_id = $3;
	`

	for _, summoner := range summoners {
		_, err := db.ExecContext(ctx, query, summoner.Champion, summoner.Name, summoner.TeamID)
		if err != nil {
			return fmt.Errorf("error updating summoner %s: %w", summoner.Name, err)
		}
	}

	return nil
}

// SwapSummoners swaps the teams of two summoners for a specific game.
func SwapSummoners(db *sql.DB, gameID int, summoner1Name string, summoner2Name string) error {
	ctx := context.Background()

	// Retrieve team IDs of both summoners and ensure they belong to the specified game
	var teamID1, teamID2 int
	query := `
		SELECT s1.team_id, s2.team_id
		FROM summoners s1, summoners s2
		WHERE s1.name = $1 AND s2.name = $2
		AND (s1.team_id = (SELECT team_blue FROM games WHERE id = $3) OR s1.team_id = (SELECT team_red FROM games WHERE id = $3))
		AND (s2.team_id = (SELECT team_blue FROM games WHERE id = $3) OR s2.team_id = (SELECT team_red FROM games WHERE id = $3));
	`
	err := db.QueryRowContext(ctx, query, summoner1Name, summoner2Name, gameID).Scan(&teamID1, &teamID2)
	if err != nil {
		return err
	}

	// Swap the teams
	_, err = db.ExecContext(ctx, "UPDATE summoners SET team_id = $1 WHERE name = $2", teamID2, summoner1Name)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, "UPDATE summoners SET team_id = $1 WHERE name = $2", teamID1, summoner2Name)
	return err
}

// Clear teams in a game (for shuffling)
func ClearTeams(db *sql.DB, gameID int) error {
	ctx := context.Background()
	query := `
    DELETE FROM summoners WHERE team_id IN (
        SELECT team_blue FROM games WHERE id = $1
        UNION
        SELECT team_red FROM games WHERE id = $1
    );
    `
	_, err := db.ExecContext(ctx, query, gameID)
	return err
}

// SetWinner sets the winner of a game.
func SetWinner(db *sql.DB, gameID int, winner string) error {
	ctx := context.Background()

	query := `
	UPDATE games
	SET winner = $1
	WHERE id = $2;
	`
	_, err := db.ExecContext(ctx, query, winner, gameID)
	return err
}

// GetSummonerLane retrieves the lane of a summoner for a specific game.
func GetSummonerLane(db *sql.DB, summonerName string, gameID int) (string, error) {
	ctx := context.Background()

	var lane string
	query := `
		SELECT role
		FROM summoners
		WHERE name = $1 AND team_id IN (
			SELECT team_blue FROM games WHERE id = $2
			UNION
			SELECT team_red FROM games WHERE id = $2
		);
	`
	err := db.QueryRowContext(ctx, query, summonerName, gameID).Scan(&lane)
	if err != nil {
		return "", err
	}

	return lane, nil
}

// UpdateSummonerChampion updates the champion of a summoner.
func UpdateSummonerChampion(db *sql.DB, summonerName string, champion string) error {
	ctx := context.Background()

	query := `
		UPDATE summoners
		SET champion = $1
		WHERE name = $2;
	`
	_, err := db.ExecContext(ctx, query, champion, summonerName)
	return err
}

// DropDatabase drops all tables from the database.
func DropDatabase(db *sql.DB) error {
	ctx := context.Background()

	tables := []string{"summoners", "teams", "games"}

	for _, table := range tables {
		if _, err := db.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", table)); err != nil {
			log.Fatalf("Error dropping table %s: %v", table, err)
			return err
		}
	}

	log.Println("Database dropped successfully.")
	return nil
}
