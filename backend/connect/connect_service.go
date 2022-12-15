package connect

import (
	context "context"

	globals "backend/globals"
	helpers "backend/helpers"
)

type Server struct {
	ConnectServiceServer
}

func (s *Server) Follow(ctx context.Context, in *FollowRequest) (*FollowResponse, error) {

	if !helpers.StringInSlice(in.User2, globals.Following[in.User1]) {
		globals.Following[in.User1] = append(globals.Following[in.User1], in.User2)
		globals.Followers[in.User2] = append(globals.Followers[in.User2], in.User1)
	}

	return &FollowResponse{Success: true}, nil
}

func (s *Server) Unfollow(ctx context.Context, in *UnfollowRequest) (*UnfollowResponse, error) {

	i := helpers.IndexOf(in.User2, globals.Following[in.User1])
	globals.Following[in.User1] = helpers.RemoveFromSlice(globals.Following[in.User1], i)

	j := helpers.IndexOf(in.User1, globals.Followers[in.User2])
	globals.Followers[in.User2] = helpers.RemoveFromSlice(globals.Followers[in.User2], j)

	return &UnfollowResponse{Success: true}, nil
}
