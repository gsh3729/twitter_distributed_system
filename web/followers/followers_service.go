package followers

import (
	// "log"
	// "strings"

	globals "proj/web/globals"
)

func GetUserFollowers(username string) []string {
	userFollowers := globals.Followers[username]
	return userFollowers
} 