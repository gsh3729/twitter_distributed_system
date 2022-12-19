package authbackend

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func cleanup_function()

func TestAuth1(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username string = "harshaG"
	var password string = "proj123"

	auth_server := NewAuthServiceClient(conn)
	response, err := auth_server.SignUp(context.Background(), &UserSignUpRequest{
		Username: username,
		Password: password,
	})

	if err != nil || !response.Success {
		t.Error("TestAuth signup failed: ", err)
	}

	resp, err := auth_server.SignIn(context.Background(), &UserSignInRequest{
		Username: username,
		Password: password,
	})

	if err != nil || !resp.Success {
		t.Error("TestAuth signin failed: ", err)
	}

	log.Printf("Auth test1 passed successfully")
}

func TestAuth2(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username string = "harshaG"
	var password string = "proj123"

	auth_server := NewAuthServiceClient(conn)
	response, err := auth_server.SignUp(context.Background(), &UserSignUpRequest{
		Username: username,
		Password: password,
	})

	if err != nil || response.Success {
		t.Error("TestAuth signup failed: ", err)
	}

	password = "p123"
	resp, err := auth_server.SignIn(context.Background(), &UserSignInRequest{
		Username: username,
		Password: password,
	})

	if err != nil || resp.Success {
		t.Error("TestAuth signin failed: ", err)
	}
	
	cleanup := cleanup_function()

	log.Printf("Auth test2 passed successfully")
}
