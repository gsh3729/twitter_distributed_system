package following


import (
	// "log"
	// "strings"

	globals "proj/web/globals"
)

func GetUserFollowing(username string) []string {
	userFollowing := globals.Following[username]
	return userFollowing
} 