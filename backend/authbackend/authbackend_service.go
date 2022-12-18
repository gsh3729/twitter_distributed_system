package authbackend

import (
	"encoding/json"
	"log"

	globals "backend/globals"
	helpers "backend/helpers"
	context "context"
)

type Server struct {
	AuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	users := make(map[string]globals.User)

	resp := helpers.GetValueForKey("users")

	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &users)

		if _, exists := users[in.Username]; exists {
			log.Print("User already exists")
			return &UserSignUpResponse{Success: false}, nil
		}
	}

	newuser := users[in.Username]
	newuser.Username = in.Username
	newuser.Password = in.Password
	users[in.Username] = newuser

	updatedusers, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	helpers.PutValueForKeys("users", string(updatedusers))

	return &UserSignUpResponse{Success: true}, nil
}

func (s *Server) SignIn(ctx context.Context, in *UserSignInRequest) (*UserSignInResponse, error) {
	var users = make(map[string]globals.User)

	// Get data from raft
	resp := helpers.GetValueForKey("users")

	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &users)
	}

	if val, exists := users[in.Username]; !exists {
		log.Print("User does not exist")
		return &UserSignInResponse{Success: false, Username: "User does not exist"}, nil
	} else {
		is_valid := val.Password == in.Password
		return &UserSignInResponse{Success: is_valid, Username: in.Username}, nil
	}
}
