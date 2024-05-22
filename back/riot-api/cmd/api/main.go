package main

import (
	"database/sql"
	"log"
	"riot-micro/cmd/data"

	_ "github.com/lib/pq"
)

const grpcPort = "50001"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	// Initialize the database connection
	db, err := sql.Open("postgres", "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("error opening database:", err)
	}
	defer db.Close()

	// Ensure the database is properly set up
	if err := data.InitializeDatabase(db); err != nil {
		log.Fatal("error initializing database:", err)
	}

	// Set up the models
	models := data.NewModels(db)

	// Set up application configuration
	app := Config{
		Models: models,
		DB:     db,
	}

	// Start the gRPC server
	app.gRPCListen()
}
