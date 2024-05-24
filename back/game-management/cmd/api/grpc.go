package main

import (
	"fmt"

	"google.golang.org/grpc"
)

func ConnectGRPC() *grpc.ClientConn {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:50001", opts...)
	if err != nil {
		fmt.Println("Failed to connect to gRPC server:", err)
		return nil
	}
	return conn
}
