package follow

import (
	context "context"
	"encoding/json"
	"log"

	globals "backend/globals"
	"backend/helpers"
)

type Server struct {
	FollowServiceServer
}

func GetMap(key string) map[string][]string {
	return_map := make(map[string][]string)

	following_resp := helpers.GetValueForKey(key)
	for _, ev := range following_resp.Kvs {
		json.Unmarshal(ev.Value, &return_map)
	}

	return return_map
}

func PutMap(map_to_put map[string][]string) {
	updated_map, err := json.Marshal(map_to_put)
	if err != nil {
		log.Println(err)
	}
	helpers.PutValueForKeys("following", string(updated_map))

}

func (s *Server) Follow(ctx context.Context, in *FollowRequest) (*FollowResponse, error) {
	followers, following := GetMap("followers")

	if !helpers.StringInSlice(in.User2, following[in.User1]) {
		following[in.User1] = append(following[in.User1], in.User2)
		followers[in.User2] = append(followers[in.User2], in.User1)
	}

	PutFollowersAndFollowingMaps(followers, following)

	return &FollowResponse{Success: true}, nil
}

func (s *Server) Unfollow(ctx context.Context, in *UnfollowRequest) (*UnfollowResponse, error) {
	followers, following := GetFollowersAndFollowingMaps()

	i := helpers.IndexOf(in.User2, following[in.User1])
	following[in.User1] = helpers.RemoveFromSlice(following[in.User1], i)

	j := helpers.IndexOf(in.User1, followers[in.User2])
	followers[in.User2] = helpers.RemoveFromSlice(followers[in.User2], j)

	PutFollowersAndFollowingMaps(followers, following)

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
