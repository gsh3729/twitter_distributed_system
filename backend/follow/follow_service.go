package follow

import (
	context "context"
	"encoding/json"

	globals "backend/globals"
	"backend/helpers"
)

type Server struct {
	FollowServiceServer
}

func (s *Server) Follow(ctx context.Context, in *FollowRequest) (*FollowResponse, error) {
	followers := helpers.GetMap("followers")
	following := helpers.GetMap("following")

	if !helpers.StringInSlice(in.User2, following[in.User1]) {
		following[in.User1] = append(following[in.User1], in.User2)
		followers[in.User2] = append(followers[in.User2], in.User1)
	}

	helpers.PutMap("followers", followers)
	helpers.PutMap("following", following)

	return &FollowResponse{Success: true}, nil
}

func (s *Server) Unfollow(ctx context.Context, in *UnfollowRequest) (*UnfollowResponse, error) {
	followers := helpers.GetMap("followers")
	following := helpers.GetMap("following")

	i := helpers.IndexOf(in.User2, following[in.User1])
	following[in.User1] = helpers.RemoveFromSlice(following[in.User1], i)

	j := helpers.IndexOf(in.User1, followers[in.User2])
	followers[in.User2] = helpers.RemoveFromSlice(followers[in.User2], j)

	helpers.PutMap("followers", followers)
	helpers.PutMap("following", following)

	return &UnfollowResponse{Success: true}, nil
}

func (s *Server) GetUserFollowers(ctx context.Context, in *GetFollowersRequest) (*GetFollowersResponse, error) {
	followers := helpers.GetMap("followers")
	userFollowers := followers[in.Username]
	return &GetFollowersResponse{Users: userFollowers, Success: true}, nil
}

func (s *Server) GetUserFollowing(ctx context.Context, in *GetFollowingRequest) (*GetFollowingResponse, error) {
	following := helpers.GetMap("following")
	userFollowing := following[in.Username]
	return &GetFollowingResponse{Users: userFollowing, Success: true}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *GetUsersRequest) (*GetUsersResponse, error) {
	people := []string{}
	var users = make(map[string]globals.User)

	resp := helpers.GetValueForKey("users")
	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &users)
	}

	for user := range users {
		if user != in.Username {
			people = append(people, user)
		}
	}

	return &GetUsersResponse{
		Users:   people,
		Success: true,
	}, nil
}
