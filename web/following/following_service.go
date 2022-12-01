package following


import (
	// "log"
	// "strings"

	globals "proj/web/globals"
)

func GetUserFollowing(username string) []string {
	userFollowers := globals.Following[username]
	return userFollowers
} 