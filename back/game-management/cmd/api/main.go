package main

import (
	"database/sql"
	"fmt"
	"game-management-micro/data"
	proto "game-management-micro/riot"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"

var counts int64

type Config struct {
	DB      *data.DB
	Models  data.Models
	RiotAPI proto.RiotAPIServiceClient
}

func main() {
	fmt.Println("Starting game service")

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		fmt.Println("Can't connect to Postgres!")
		return
	}

	// wrap the sql.DB connection
	db := data.NewDB(conn)

	grpcConn := ConnectGRPC()
	if grpcConn == nil {
		fmt.Println("Can't connect to grpc server")
		return
	}

	riotAPIClient := proto.NewRiotAPIServiceClient(grpcConn)
	fmt.Println("Connected to gRPC server successfully")

	// set up config
	app := Config{
		DB:      db,
		Models:  data.NewModels(db),
		RiotAPI: riotAPIClient,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func connectToDB() *sql.DB {
	for {
		connection, err := openDB("host=localhost port=5433 user=postgres dbname=postgres sslmode=disable connect_timeout=5 password=admin")
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
