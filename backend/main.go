package main

import (
	"fmt"
	"log"
	"net"

	"backend/auth"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("Failed to Listen to TCP: %v", err)
	}

	auth_server := auth.Server{}

	Server := grpc.NewServer()

	auth.RegisterAuthServiceServer(Server, &auth_server)

	if err := Server.Serve(lis); err != nil {
		log.Fatalf("Failed to Listen to TCP: %s", err)
	}
}
