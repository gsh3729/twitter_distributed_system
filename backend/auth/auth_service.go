package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	context "context"

	clientv3 "go.etcd.io/etcd/client/v2"
)

type Server struct {
	AuthServiceServer
}

type User struct {
	Username string
	Password string
}

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	// var users = make(map[string]User)

	// Get data from raft
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:12380", "localhost:22380", "localhost:32380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Get(ctx, "users")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
	// if _, exists := resp.Kvs[in.Username]; exists {
	// 	log.Print("User already exists")
	// 	return &UserSignUpResponse{Success: false}, nil
	// }

	// newuser := users[in.Username]
	// newuser.Username = in.Username
	// newuser.Password = in.Password
	// users[in.Username] = newuser

	// usersjson, err := json.Marshal(users)
	// if err != nil {
	// 	log.Print(err)
	// }
	// cmd := exec.Command("curl", "-L", "http://127.0.0.1:12380/users", "-XPUT", "-d "+string(usersjson))

	// cmd.Run()

	return &UserSignUpResponse{Success: true}, nil
}

func (s *Server) SignIn(ctx context.Context, in *UserSignInRequest) (*UserSignInResponse, error) {
	var users = make(map[string]User)

	// Get data from raft
	resp, err := http.Get("http://127.0.0.1:12380/users")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	json.Unmarshal(body, &users)
	if val, exists := users[in.Username]; !exists {
		log.Print("User does not exist")
		return &UserSignInResponse{Success: false}, nil
	} else {
		is_valid := val.Password == in.Password
		return &UserSignInResponse{Success: is_valid, Username: in.Username}, nil
	}
}
