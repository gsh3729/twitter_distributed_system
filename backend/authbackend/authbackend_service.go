package authbackend

import (
	"encoding/json"
	"log"

	globals "backend/globals"
	context "context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	AuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	users := make(map[string]globals.User)

	

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

	_, err = cli.Put(context.TODO(), "users", string(updatedusers))
	if err != nil {
		log.Fatal(err)
	}

	return &UserSignUpResponse{Success: true}, nil
}

func (s *Server) SignIn(ctx context.Context, in *UserSignInRequest) (*UserSignInResponse, error) {
	var users = make(map[string]globals.User)

	// Get data from raft
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   globals.Endpoints,
		DialTimeout: globals.Timeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx2, cancel := context.WithTimeout(context.Background(), globals.Timeout)
	resp, err := cli.Get(ctx2, "users")
	cancel()
	if err != nil {
		log.Fatal(err)
	}

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
