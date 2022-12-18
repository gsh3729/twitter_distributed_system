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
	following := make(map[string][]string)
	followers := make(map[string][]string)

	following_resp := helpers.GetValueForKey("following")
	for _, ev := range following_resp.Kvs {
		json.Unmarshal(ev.Value, &following)
	}

	follower_resp := helpers.GetValueForKey("following")
	for _, ev := range following_resp.Kvs {
		json.Unmarshal(ev.Value, &followers)
	}

	if !helpers.StringInSlice(in.User2, following[in.User1]) {
		following[in.User1] = append(globals.Following[in.User1], in.User2)
		followers[in.User2] = append(globals.Followers[in.User2], in.User1)
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

func (s *Server) GetUserFollowers(ctx context.Context, in *GetFollowersRequest) (*GetFollowersResponse, error) {
	userFollowers := globals.Followers[in.Username]
	return &GetFollowersResponse{Users: userFollowers, Success: true}, nil
}

func (s *Server) GetUserFollowing(ctx context.Context, in *GetFollowingRequest) (*GetFollowingResponse, error) {
	userFollowing := globals.Following[in.Username]
	return &GetFollowingResponse{Users: userFollowing, Success: true}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *GetUsersRequest) (*GetUsersResponse, error) {
	people := []string{}
	for key := range globals.UserPass {
		if key != in.Username {
			people = append(people, key)
		}
	}
	return &GetUsersResponse{
		Users:   people,
		Success: true,
	}, nil
}
