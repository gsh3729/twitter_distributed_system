package authbackend

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	context "context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	AuthServiceServer
}

type User struct {
	Username string
	Password string
}

var timeout = 5 * time.Second

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	var users = make(map[string]User)

	// Get data from raft
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:12380", "localhost:22380", "localhost:32380"},
		DialTimeout: timeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx2, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := cli.Get(ctx2, "users")
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	newuser := users[in.Username]
	newuser.Username = in.Username
	newuser.Password = in.Password
	users[in.Username] = newuser

	usersjson, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
	}
	cmd := exec.Command("curl", "-L", "http://127.0.0.1:12380/users", "-XPUT", "-d "+string(usersjson))

	cmd.Run()

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
