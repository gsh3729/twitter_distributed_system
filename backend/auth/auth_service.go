package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	context "context"
)

type Server struct {
	AuthServiceServer
}

type User struct {
	Username string
	Password string
}

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	var users = make(map[string]User)

	// Get data from raft
	resp, err := http.Get("http://127.0.0.1:12380/users")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &users)
	if _, exists := users[in.Username]; exists {
		fmt.Println("User already exists")
		return &UserSignUpResponse{Success: false}, nil
	}

	newuser := users[in.Username]
	newuser.Username = in.Username
	newuser.Password = in.Password
	users[in.Username] = newuser

	usersjson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &users)
	if val, exists := users[in.Username]; !exists {
		fmt.Println("User does not exist")
		return &UserSignInResponse{Success: false}, nil
	} else {
		is_valid := val.Password == in.Password
		return &UserSignInResponse{Success: is_valid, Username: in.Username}, nil
	}
}
