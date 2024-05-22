package data

import (
	"context"
	"database/sql"
	"log"
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

	log.Println("Database initialized successfully.")
	return nil
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
