package auth

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"os/exec"

	context "context"

	globals "backend/globals"
	// helpers "backend/helpers"
)

type Server struct {
	AuthServiceServer
}

type User struct {
	Username string
	Password string
}

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	var users = 
	resp, err := http.Get("http://127.0.0.1:12380/users")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	globals.UserPass[in.Username] = in.Password
	return &UserSignUpResponse{Success: true}, nil
}

func (s *Server) SignIn(ctx context.Context, in *UserSignInRequest) (*UserSignInResponse, error) {
	is_valid := CheckUserPass(in.Username, in.Password)
	return &UserSignInResponse{Success: is_valid, Username: in.Username}, nil
}

func CheckUserPass(username, password string) bool { //validate
	userpass := globals.UserPass

	if val, ok := userpass[username]; ok {
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}