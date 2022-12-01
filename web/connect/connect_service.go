package connect

import (
	// "log"
	// "strings"

	globals "proj/web/globals"
	helpers "proj/web/helpers"
)

func Follow(username1 string, username2 string) []string {

	if !helpers.StringInSlice(username2, globals.Following[username1]) {
		globals.Following[username1] = append(globals.Following[username1], user)
		globals.Followers[connectTo] = append(globals.Followers[connectTo], user.(string))
	}

	userFollowers := globals.Followers[username]
	return userFollowers
} 