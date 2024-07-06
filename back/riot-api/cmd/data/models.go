package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type DB struct {
	*sql.DB
}

type Models struct {
	LeagueEntry     LeagueEntry
	Summoner        Summoner
	ChampionMastery ChampionMastery
}

func NewModels(db *sql.DB) Models {
	return Models{
		LeagueEntry:     LeagueEntry{},
		Summoner:        Summoner{},
		ChampionMastery: ChampionMastery{},
	}
}

func NewDB(database *sql.DB) *DB {
	return &DB{DB: database}
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	SummonerLevel int    `json:"summonerLevel"`
}

type ChampionMastery struct {
	PUUID                        string `json:"puuid"`
	ChampionId                   int64  `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChestGranted                 bool   `json:"chestGranted"`
}

type LeagueEntry struct {
	LeagueID     string `json:"leagueId"`
	SummonerID   string `json:"summonerId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

type SummonerLane struct {
	SummonerName string `json:"summoner_name"`
	Lane         string `json:"lane"`
}

type SummonerChampion struct {
	SummonerName   string `json:"summoner_name"`
	Lane           string `json:"lane"`
	ChampionName   string `json:"champion_name"`
	ChampionPoints int    `json:"champion_points"`
}

func GetRankPoints(tier string, rank string) int {
	tierPoints := map[string]int{
		"IRON":        0,
		"BRONZE":      10,
		"SILVER":      20,
		"GOLD":        30,
		"PLATINUM":    40,
		"DIAMOND":     50,
		"MASTER":      60,
		"GRANDMASTER": 70,
		"CHALLENGER":  80,
	}

	rankPoints := map[string]int{
		"IV":  1,
		"III": 2,
		"II":  3,
		"I":   4,
	}

	return tierPoints[tier] + rankPoints[rank]
}

func InitializeDatabase(db *sql.DB) error {
	ctx := context.Background()

	// Ensure the 'summoners' table is set up with 'puuid' as a unique field
	if _, err := db.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS summoners (
        id VARCHAR(255) PRIMARY KEY,
        account_id VARCHAR(255),
        puuid VARCHAR(255) UNIQUE,
        name VARCHAR(255),
        tag VARCHAR(255),
        summoner_level INT
    );`); err != nil {
		log.Fatalf("Error creating summoners table: %v", err)
		return err
	}

	// Create the 'champion_masteries' table with a foreign key that references the 'puuid' in 'summoners'
	if _, err := db.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS champion_masteries (
        puuid VARCHAR(255),
        champion_id INT,
        champion_level INT,
        champion_points INT,
        last_play_time BIGINT,
        points_since_last_level INT,
        points_until_next_level INT,
        tokens_earned INT,
        chest_granted BOOLEAN,
        PRIMARY KEY (puuid, champion_id),
        FOREIGN KEY (puuid) REFERENCES summoners(puuid)
    );`); err != nil {
		log.Fatalf("Error creating champion masteries table: %v", err)
		return err
	}

	// Create the 'league_entries' table with a foreign key that references the 'id' in 'summoners'
	if _, err := db.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS league_entries (
        league_id VARCHAR(255),
        summoner_id VARCHAR(255),
        queue_type VARCHAR(50),
        tier VARCHAR(50),
        rank VARCHAR(50),
        league_points INT,
        wins INT,
        losses INT,
        veteran BOOLEAN,
        inactive BOOLEAN,
        fresh_blood BOOLEAN,
        hot_streak BOOLEAN,
        PRIMARY KEY (league_id, summoner_id),
        FOREIGN KEY (summoner_id) REFERENCES summoners(id)
    );`); err != nil {
		log.Fatalf("Error creating league entries table: %v", err)
		return err
	}

	// Creating lanes table
	if _, err := db.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS lanes (
        lane_name VARCHAR(50) PRIMARY KEY
    );`); err != nil {
		log.Fatalf("Error creating lanes table: %v", err)
		return err
	}

	// Create the 'champions' table
	if _, err := db.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS champions (
        champion_id INT PRIMARY KEY,
        name VARCHAR(255)
    );`); err != nil {
		log.Fatalf("Error creating champions table: %v", err)
		return err
	}

	// Create the 'champion_lanes' table
	if _, err := db.ExecContext(ctx, `
    CREATE TABLE IF NOT EXISTS champion_lanes (
        champion_id INT,
        name VARCHAR(255),
        lane_name VARCHAR(50),
        FOREIGN KEY (champion_id) REFERENCES champions(champion_id),
        FOREIGN KEY (lane_name) REFERENCES lanes(lane_name),
        PRIMARY KEY (champion_id, lane_name)
    );`); err != nil {
		log.Fatalf("Error creating champion_lanes table: %v", err)
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}

func SeedChampionLanes(db *sql.DB, champions []Champion) error {
	ctx := context.Background()

	for _, champ := range champions {
		for _, role := range champ.Roles {
			if _, err := db.ExecContext(ctx, "INSERT INTO champion_lanes (champion_id, name, lane_name) VALUES ($1, $2, $3)", champ.ChampionID, champ.Name, role); err != nil {
				log.Printf("Error inserting into champion_lanes: %v", err)
				return err
			}
		}
	}

	log.Println("champion_lanes table seeded successfully.")
	return nil
}

func SeedChampionsTable(db *sql.DB, champions []Champion) error {
	ctx := context.Background()

	for _, champ := range champions {
		if _, err := db.ExecContext(ctx, "INSERT INTO champions (champion_id, name) VALUES ($1, $2)", champ.ChampionID, champ.Name); err != nil {
			log.Printf("Error inserting into champions: %v", err)
			return err
		}
	}

	log.Println("champions table seeded successfully.")
	return nil
}

func SeedLanesTable(db *sql.DB) error {
	ctx := context.Background()

	lanes := []string{"top", "mid", "jungle", "adc", "support"}

	for _, lane := range lanes {
		if _, err := db.ExecContext(ctx, "INSERT INTO lanes (lane_name) VALUES ($1) ", lane); err != nil {
			log.Printf("Error inserting into lanes: %v", err)
			return err
		}
	}

	log.Println("lanes table seeded successfully.")
	return nil
}

func GetRandomChampionForLane(db *sql.DB, summonerName string, lane string) (string, int32, error) {
	ctx := context.Background()

	// Get the summoner's PUUID based on their name
	var puuid string
	err := db.QueryRowContext(ctx, "SELECT puuid FROM summoners WHERE name = $1", summonerName).Scan(&puuid)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No summoner found with name: %v", summonerName)
			return "", 0, nil // No summoner found
		}
		log.Printf("Error retrieving summoner PUUID: %v", err)
		return "", 0, err
	}

	// Get the list of champions for the given lane that the summoner has champion points on
	query := `
    SELECT c.name, cm.champion_points
    FROM champion_masteries cm
    JOIN champion_lanes cl ON cm.champion_id = cl.champion_id
    JOIN champions c ON cl.champion_id = c.champion_id
    WHERE cm.puuid = $1 AND cl.lane_name = $2
    `
	rows, err := db.QueryContext(ctx, query, puuid, lane)
	if err != nil {
		log.Printf("Error retrieving champions for lane: %v", err)
		return "", 0, err
	}
	defer rows.Close()

	var champions []struct {
		name           string
		championPoints int32
	}
	for rows.Next() {
		var champion struct {
			name           string
			championPoints int32
		}
		if err := rows.Scan(&champion.name, &champion.championPoints); err != nil {
			log.Printf("Error scanning champion name and points: %v", err)
			return "", 0, err
		}
		champions = append(champions, champion)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error reading champions: %v", err)
		return "", 0, err
	}

	// If no champions are found, return an empty result
	if len(champions) == 0 {
		return "", 0, nil
	}

	// Return a random champion from the list
	rand.Seed(time.Now().UnixNano())
	randomChampion := champions[rand.Intn(len(champions))]

	return randomChampion.name, randomChampion.championPoints, nil
}

// Insert a Summoner into the database.
func InsertSummoner(db *sql.DB, ctx context.Context, s Summoner) error {
	query := `
	INSERT INTO summoners (id, account_id, puuid, name, tag, summoner_level)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (id) DO UPDATE SET
	account_id = EXCLUDED.account_id,
	puuid = EXCLUDED.puuid,
	name = EXCLUDED.name,
  tag = EXCLUDED.tag,
	summoner_level = EXCLUDED.summoner_level;
	`
	_, err := db.ExecContext(ctx, query, s.ID, s.AccountID, s.PUUID, s.Name, s.Tag, s.SummonerLevel)
	return err
}

// UpdateSummoner updates an existing summoner in the database.
func UpdateSummoner(db *sql.DB, ctx context.Context, s Summoner) error {
	stmt := `
    UPDATE summoners
    SET account_id = $1, puuid = $2, name = $3, tag = $4, summoner_level = $5
    WHERE id = $6;
    `
	_, err := db.ExecContext(ctx, stmt, s.AccountID, s.PUUID, s.Name, s.Tag, s.SummonerLevel, s.ID)
	if err != nil {
		log.Printf("Error updating summoner: %v", err)
		return err
	}
	return nil
}

// GetSummoner retrieves a summoner by ID from the database.
func GetSummoner(db *sql.DB, ctx context.Context, name string) (*Summoner, error) {
	stmt := `
    SELECT id, account_id, puuid, name, tag, summoner_level
    FROM summoners
    WHERE name = $1;
    `
	row := db.QueryRowContext(ctx, stmt, name)

	var s Summoner
	err := row.Scan(&s.ID, &s.AccountID, &s.PUUID, &s.Name, &s.Tag, &s.SummonerLevel)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result is not an error, simply no summoner found
		}
		log.Printf("Error retrieving summoner: %v", err)
		return nil, err
	}
	return &s, nil
}

// Insert a Champion Mastery into the database.
func InsertChampionMastery(db *sql.DB, ctx context.Context, cm ChampionMastery) error {
	query := `
    INSERT INTO champion_masteries (
        puuid, champion_id, champion_level, champion_points, last_play_time,
        points_since_last_level, points_until_next_level, tokens_earned, chest_granted
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    ON CONFLICT (puuid, champion_id) DO UPDATE SET
        champion_level = EXCLUDED.champion_level,
        champion_points = EXCLUDED.champion_points,
        last_play_time = EXCLUDED.last_play_time,
        points_since_last_level = EXCLUDED.points_since_last_level,
        points_until_next_level = EXCLUDED.points_until_next_level,
        tokens_earned = EXCLUDED.tokens_earned,
        chest_granted = EXCLUDED.chest_granted;
    `
	_, err := db.ExecContext(ctx, query, cm.PUUID, cm.ChampionId, cm.ChampionLevel, cm.ChampionPoints, cm.LastPlayTime,
		cm.ChampionPointsSinceLastLevel, cm.ChampionPointsUntilNextLevel, cm.TokensEarned, cm.ChestGranted)
	return err
}

// UpdateChampionMastery updates an existing champion mastery in the database.
func UpdateChampionMastery(db *sql.DB, ctx context.Context, cm ChampionMastery) error {
	query := `
    UPDATE champion_masteries SET
        champion_level = $3,
        champion_points = $4,
        last_play_time = $5,
        points_since_last_level = $6,
        points_until_next_level = $7,
        tokens_earned = $8,
        chest_granted = $9
    WHERE puuid = $1 AND champion_id = $2;
    `
	_, err := db.ExecContext(ctx, query, cm.PUUID, cm.ChampionId, cm.ChampionLevel, cm.ChampionPoints, cm.LastPlayTime,
		cm.ChampionPointsSinceLastLevel, cm.ChampionPointsUntilNextLevel, cm.TokensEarned, cm.ChestGranted)
	if err != nil {
		log.Printf("Error updating champion mastery: %v", err)
		return err
	}
	return nil
}

// GetChampionMastery retrieves a specific champion mastery from the database.

func GetChampionMasteries(db *sql.DB, ctx context.Context, puuid string) ([]ChampionMastery, error) {
	const query = `
    SELECT puuid, champion_id, champion_level, champion_points, last_play_time,
           points_since_last_level, points_until_next_level, tokens_earned, chest_granted
    FROM champion_masteries
    WHERE puuid = $1;
    `
	rows, err := db.QueryContext(ctx, query, puuid)
	if err != nil {
		log.Printf("Error retrieving champion masteries: %v", err)
		return nil, err
	}
	defer rows.Close()

	var masteries []ChampionMastery
	for rows.Next() {
		var cm ChampionMastery
		if err := rows.Scan(&cm.PUUID, &cm.ChampionId, &cm.ChampionLevel, &cm.ChampionPoints, &cm.LastPlayTime,
			&cm.ChampionPointsSinceLastLevel, &cm.ChampionPointsUntilNextLevel, &cm.TokensEarned, &cm.ChestGranted); err != nil {
			log.Printf("Error scanning champion mastery: %v", err)
			return nil, err
		}
		masteries = append(masteries, cm)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error reading champion masteries: %v", err)
		return nil, err
	}
	return masteries, nil
}

// Insert a League Entry into the database.
func InsertLeagueEntry(db *sql.DB, ctx context.Context, le LeagueEntry) error {
	query := `
    INSERT INTO league_entries (
        league_id, summoner_id, queue_type, tier, rank, league_points, wins, losses, veteran, inactive, fresh_blood, hot_streak
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    ON CONFLICT (league_id, summoner_id) DO UPDATE SET
        queue_type = EXCLUDED.queue_type,
        tier = EXCLUDED.tier,
        rank = EXCLUDED.rank,
        league_points = EXCLUDED.league_points,
        wins = EXCLUDED.wins,
        losses = EXCLUDED.losses,
        veteran = EXCLUDED.veteran,
        inactive = EXCLUDED.inactive,
        fresh_blood = EXCLUDED.fresh_blood,
        hot_streak = EXCLUDED.hot_streak;
    `
	_, err := db.ExecContext(ctx, query,
		le.LeagueID, le.SummonerID, le.QueueType, le.Tier, le.Rank, le.LeaguePoints,
		le.Wins, le.Losses, le.Veteran, le.Inactive, le.FreshBlood, le.HotStreak)
	return err
}

// UpdateLeagueEntry updates an existing league entry in the database.
func UpdateLeagueEntry(db *sql.DB, ctx context.Context, le LeagueEntry) error {
	stmt := `
    UPDATE league_entries SET
        queue_type = $3, tier = $4, rank = $5, league_points = $6,
        wins = $7, losses = $8, veteran = $9, inactive = $10, fresh_blood = $11, hot_streak = $12
    WHERE league_id = $1 AND summoner_id = $2;
    `
	_, err := db.ExecContext(ctx, stmt,
		le.LeagueID, le.SummonerID, le.QueueType, le.Tier, le.Rank, le.LeaguePoints,
		le.Wins, le.Losses, le.Veteran, le.Inactive, le.FreshBlood, le.HotStreak)
	return err
}

// GetLeagueEntry retrieves a specific league entry from the database.
func GetLeagueEntries(db *sql.DB, ctx context.Context, summonerID string) ([]LeagueEntry, error) {
	const query = `
        SELECT league_id, summoner_id, queue_type, tier, rank, league_points,
               wins, losses, veteran, inactive, fresh_blood, hot_streak
        FROM league_entries
        WHERE summoner_id = $1;
    `
	rows, err := db.QueryContext(ctx, query, summonerID)
	if err != nil {
		log.Printf("Error retrieving league entries: %v", err)
		return nil, err
	}
	defer rows.Close()

	var entries []LeagueEntry
	for rows.Next() {
		var le LeagueEntry
		err := rows.Scan(&le.LeagueID, &le.SummonerID, &le.QueueType, &le.Tier, &le.Rank, &le.LeaguePoints,
			&le.Wins, &le.Losses, &le.Veteran, &le.Inactive, &le.FreshBlood, &le.HotStreak)
		if err != nil {
			log.Printf("Error scanning league entry: %v", err)
			return nil, err
		}
		entries = append(entries, le)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error reading league entries: %v", err)
		return nil, err
	}
	return entries, nil
}

func GetBalancedChampionsForLanes(db *sql.DB, team1 []SummonerLane, team2 []SummonerLane) ([]SummonerChampion, []SummonerChampion, error) {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	const maxAttempts = 10

	for attempt := 0; attempt < maxAttempts; attempt++ {
		team1Champions, err := getChampionsForTeam(ctx, db, team1)
		if err != nil {
			return nil, nil, err
		}

		team2Champions, err := getChampionsForTeam(ctx, db, team2)
		if err != nil {
			return nil, nil, err
		}

		balanced := balanceTeams(team1Champions, team2Champions)
		if balanced {
			return team1Champions, team2Champions, nil
		}
	}

	return nil, nil, errors.New("unable to balance teams within the 25% champion points difference after 10 attempts")
}

func getChampionsForTeam(ctx context.Context, db *sql.DB, team []SummonerLane) ([]SummonerChampion, error) {
	var teamChampions []SummonerChampion

	for _, sl := range team {
		champion, err := getRandomChampionForSummonerLane(ctx, db, sl.SummonerName, sl.Lane)
		if err != nil {
			return nil, err
		}
		teamChampions = append(teamChampions, champion)
	}

	return teamChampions, nil
}

func balanceTeams(team1 []SummonerChampion, team2 []SummonerChampion) bool {
	points1 := totalChampionPoints(team1)
	points2 := totalChampionPoints(team2)

	difference := float64(points1-points2) / float64(points1+points2)
	return difference <= 0.25
}

func totalChampionPoints(team []SummonerChampion) int {
	totalPoints := 0
	for _, sc := range team {
		totalPoints += sc.ChampionPoints
	}
	return totalPoints
}

func GetChampionPointsByUser(ctx context.Context, db *sql.DB, summonerName, championName string) (int, error) {
	var puuid string
	var points int
	var championID int

	err := db.QueryRowContext(ctx, "SELECT puuid FROM summoners WHERE name = $1", summonerName).Scan(&puuid)
	if err != nil {
		return 0, err
	}

	err = db.QueryRowContext(ctx, "SELECT champion_id FROM champions WHERE name = $1", championName).Scan(&championID)
	if err != nil {
		return 0, err
	}

	query := `
    SELECT champion_points
    FROM champion_masteries
    WHERE puuid = $1 AND champion_id = $2
    `
	err = db.QueryRowContext(ctx, query, puuid, championID).Scan(&points)
	if err != nil {
		return 0, err
	}
	return points, nil
}

func GetEloByUser(ctx context.Context, db *sql.DB, summonerName string) (int, error) {
	var elo int

	summoner, err := GetSummoner(db, ctx, summonerName)
	if err != nil {
		return 0, err
	}
	if summoner == nil {
		return 0, fmt.Errorf("summoner %s not found", summonerName)
	}

	entries, err := GetLeagueEntries(db, ctx, summoner.ID)
	if err != nil {
		return 0, err
	}

	if len(entries) == 0 {
		return 0, nil
	}

	highestRank := entries[0]
	for _, entry := range entries {
		if GetRankPoints(entry.Tier, entry.Rank) > GetRankPoints(highestRank.Tier, highestRank.Rank) {
			highestRank = entry
		}
	}

	elo = GetRankPoints(highestRank.Tier, highestRank.Rank)
	return elo, nil
}

func getRandomChampionForSummonerLane(ctx context.Context, db *sql.DB, summonerName string, lane string) (SummonerChampion, error) {
	var sc SummonerChampion

	var puuid string
	err := db.QueryRowContext(ctx, "SELECT puuid FROM summoners WHERE name = $1", summonerName).Scan(&puuid)
	if err != nil {
		return sc, err
	}

	query := `
    SELECT c.name, cm.champion_points
    FROM champion_masteries cm
    JOIN champion_lanes cl ON cm.champion_id = cl.champion_id
    JOIN champions c ON cl.champion_id = c.champion_id
    WHERE cm.puuid = $1 AND cl.lane_name = $2
    `
	rows, err := db.QueryContext(ctx, query, puuid, lane)
	if err != nil {
		return sc, err
	}
	defer rows.Close()

	var champions []SummonerChampion
	for rows.Next() {
		var champion SummonerChampion
		if err := rows.Scan(&champion.ChampionName, &champion.ChampionPoints); err != nil {
			return sc, err
		}
		champion.SummonerName = summonerName
		champion.Lane = lane
		champions = append(champions, champion)
	}
	if err = rows.Err(); err != nil {
		return sc, err
	}

	if len(champions) == 0 {
		return sc, errors.New("no champions found for the summoner and lane")
	}

	randomChampion := champions[rand.Intn(len(champions))]

	return randomChampion, nil
}
