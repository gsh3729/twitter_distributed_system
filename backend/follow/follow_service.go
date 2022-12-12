package follow

import (
	globals "backend/globals"
)

func GetUserFollowers(username string) []string {
	userFollowers := globals.Followers[username]
	return userFollowers
}

func GetUserFollowing(username string) []string {
	userFollowing := globals.Following[username]
	return userFollowing
}
