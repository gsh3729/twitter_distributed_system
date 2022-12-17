package authbackend

import (
	"testing"
	"log"
	"context"
	"google.golang.org/grpc"
	// . "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	// . "."
)


func TestAuth(t *testing.T) {
	// ctx := context.Background()
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username string = "harshaG"
	var password string = "proj123"

	auth_server := NewAuthServiceClient(conn)
	response, err := tweet_server.SignUp(context.Background(), &UserSignUpRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Fatalf("TestAuth failed: %v", err)
	}

	if !response.Success {
		t.Error("TestAuth Failed")
	}


	// check the content is posted or not

	// rep, err := tweet_server.GetTweets(context.Background(), &GetTweetsRequest{
	// 	Username: username,
	// })

	log.Printf("Posted a new tweet successfully")
}