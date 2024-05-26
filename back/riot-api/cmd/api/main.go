package main

import (
	"database/sql"
	"fmt"
	"log"
	"riot-micro/cmd/data"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const grpcPort = "50001"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	// connect to DB
	conn := connectToDB()
	if conn == nil {
		fmt.Println("Can't connect to Postgres!")
		return
	}

	// wrap the sql.DB connection
	db := data.NewDB(conn)

	// Set up the models
	models := data.NewModels(db.DB)

	// Set up application configuration
	app := Config{
		Models: models,
		DB:     db.DB,
	}

	// Start the gRPC server
	app.gRPCListen()
}

func connectToDB() *sql.DB {
	for {
		connection, err := openDB("host=riot-postgres port=5432 user=postgres dbname=postgres sslmode=disable connect_timeout=5 password=admin")
		if err != nil {
			fmt.Println("Postgres not yet ready ...")
			counts++
		} else {
			fmt.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			fmt.Println(err)
			return nil
		}

		fmt.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := data.InitializeDatabase(db); err != nil {
		log.Fatal("error initializing database:", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
