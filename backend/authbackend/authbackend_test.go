package authbackend

import (
	"backend/globals"
	"backend/helpers"
	"context"
	"encoding/json"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func cleanup() {
	users := make(map[string]globals.User)

	resp := helpers.GetValueForKey("users")

	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &users)
	}

	delete(users, "harshaG")

	updatedusers, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	helpers.PutValueForKeys("users", string(updatedusers))
}

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

	defer cleanup()

	log.Printf("Auth test2 passed successfully")
}
