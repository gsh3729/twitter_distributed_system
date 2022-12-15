package auth

import (
	context "context"

	globals "backend/globals"
	helpers "backend/helpers"
)

type Server struct {
	AuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, in *UserSignUpRequest) (*UserSignUpResponse, error) {
	globals.UserPass[in.Username] = in.Password
	return &UserSignUpResponse{Success: true}, nil
}

func (s *Server) SignIn(ctx context.Context, in *UserSignInRequest) (*UserSignInResponse, error) {
	is_valid := helpers.CheckUserPass(in.Username, in.Password)
	return &UserSignInResponse{Success: is_valid, Username: in.Username}, nil
}
