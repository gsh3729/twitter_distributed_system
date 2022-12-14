package auth

import (
	"context"
	"log"
	globals "backend/globals"
	// helpers "backend/helpers"
	pb "backend/auth/proto"
	"github.com/Jille/raft-grpc-leader-rpc/rafterrors"
	"github.com/hashicorp/raft"
)

type rpcInterface struct {
	// wordTracker *wordTracker
	raft        *raft.Raft
}

func (r rpcInterface) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	f := r.raft.Apply([]byte(req.GetUsername()), time.Second)
	if err := f.Error(); err != nil {
		return nil, rafterrors.MarkRetriable(err)
	}
	globals.UserPass[req.GetUsername] = req.GetPassword
	return &pb.SignUpResponse{
		Username: req.GetUsername(),
	}, nil
}


func SignUp(username string, password string) string {
	globals.UserPass[username] = password
	return username
}

func SignIn(username string, password string) bool {
	return CheckUserPass(username, password)
}

func CheckUserPass(username, password string) bool {
	userpass := globals.UserPass

	log.Println("checkUserPass", username, password, userpass)

	if val, ok := userpass[username]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}