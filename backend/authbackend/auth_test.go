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

	if err != nil {
		t.Fatalf("TestAuth signup failed: %v", err)
	}

	if !response.Success {
		t.Error("TestAuth signup Failed")
	}

	resp, err := auth_server.SignIn(context.Background(), &UserSignInRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Fatalf("TestAuth signin failed: %v", err)
	}

	if !resp.Success {
		t.Error("TestAuth signin Failed")
	}

	log.Printf("Auth tests passed successfully")
}