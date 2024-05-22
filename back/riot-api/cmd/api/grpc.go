package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"riot-micro/cmd/data"
	"riot-micro/riot"

	"google.golang.org/grpc"
)

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	riot.RegisterRiotAPIServiceServer(s, &RiotAPIServer{
		Models: app.Models,
		db:     app.DB,
	})

	log.Printf("gRPC Server started on port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

type RiotAPIServer struct {
	db *sql.DB
	riot.UnimplementedRiotAPIServiceServer
	Models data.Models
}
