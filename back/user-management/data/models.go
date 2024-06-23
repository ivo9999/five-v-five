package data

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	*sql.DB
}

func NewDB(database *sql.DB) *DB {
	return &DB{DB: database}
}

type Models struct {
	User
}

func NewModels(db *DB) Models {
	return Models{
		User: User{},
	}
}

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	LeagueName  string `json:"league_name"`
	LeagueTag   string `json:"league_tag"`
	DiscordName string `json:"discord_name"`
	ID          int    `json:"id"`
}

func encryptPassword(pw string) (string, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encpw), nil
}

func InitializeDatabase(db *sql.DB) error {
	ctx := context.Background()

	if _, err := db.ExecContext(ctx, createUsersTable); err != nil {
		return err
	}
	return nil
}

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username varchar(32),
  password varchar(255),
  league_name varchar(32),
  league_tag varchar(32),
  discord_name varchar(32)
);`

func InsertUser(db *sql.DB, user User) (int, error) {
	ctx := context.Background()

	pw, passErr := encryptPassword(user.Password)
	if passErr != nil {
		return 0, passErr
	}
	user.Password = pw

	query := `INSERT INTO users (username, password, league_name, league_tag, discord_name) 
  VALUES ($1, $2, $3, $4, $5)
  RETURNING id`

	var userID int
	err := db.QueryRowContext(ctx, query, user.Username, user.Password, user.LeagueName, user.LeagueTag, user.DiscordName).Scan(&userID)
	return userID, err
}

func UpdateUser(db *sql.DB, user User) error {
	ctx := context.Background()

	pw, passErr := encryptPassword(user.Password)
	if passErr != nil {
		return passErr
	}
	user.Password = pw

	query := `UPDATE users SET username = $1, password = $2, league_name = $3, league_tag = $4, discord_name = $5 
  WHERE id = $6
  RETURNING true`

	var isUpdated bool
	err := db.QueryRowContext(ctx, query, user.Username, user.Password, user.LeagueName, user.LeagueTag, user.DiscordName, user.ID).Scan(&isUpdated)
	return err
}

func GetUser(db *sql.DB, id int) (User, error) {
	ctx := context.Background()

	query := `SELECT id, username, password, league_name, league_tag, discord_name 
  FROM users 
  WHERE id = $1`

	var user User
	err := db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Password, &user.LeagueName, &user.LeagueTag, &user.DiscordName)
	return user, err
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	ctx := context.Background()

	query := `SELECT id, username, password, league_name, league_tag, discord_name 
  FROM users`

	var users []User
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.LeagueName, &user.LeagueTag, &user.DiscordName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByUsername(db *sql.DB, username string) (User, error) {
	ctx := context.Background()

	query := `SELECT id, username, password, league_name, league_tag, discord_name 
  FROM users 
  WHERE username = $1`

	var user User
	err := db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.LeagueName, &user.LeagueTag, &user.DiscordName)
	return user, err
}

type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	ID       int    `json:"id"`
}

func GetUserByUsernameLogin(db *sql.DB, username string) (LoginResponse, error) {
	ctx := context.Background()

	query := `SELECT id, username, password
  FROM users 
  WHERE username = $1`

	var user LoginResponse
	err := db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, user.Password)
	return user, err
}
