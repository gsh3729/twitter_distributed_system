package main

import (
	"fmt"
	"log"
	"net"

	"backend/authbackend"
	"backend/follow"
	"backend/tweet"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("Failed to Listen to TCP: %v", err)
	}

	auth_server := authbackend.Server{}
	follow_server := follow.Server{}
	tweet_server := tweet.Server{}

	Server := grpc.NewServer()

	authbackend.RegisterAuthServiceServer(Server, &auth_server)
	follow.RegisterFollowServiceServer(Server, &follow_server)
	tweet.RegisterTweetServiceServer(Server, &tweet_server)

	if err := Server.Serve(lis); err != nil {
		log.Fatalf("Failed to Listen to TCP: %s", err)
	}
}
