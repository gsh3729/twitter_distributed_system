package connect

import (
	// "log"
	// "strings"

	globals "proj/web/globals"
	helpers "proj/web/helpers"
)

func Follow(username1 string, username2 string) bool {

	if !helpers.StringInSlice(username2, globals.Following[username1]) {
		globals.Following[username1] = append(globals.Following[username1], username2)
		globals.Followers[username2] = append(globals.Followers[username2], username1)
	}

	return true
} 

func Unfollow(username1 string, username2 string) bool {

	i := helpers.IndexOf(username2, globals.Following[username1])
	globals.Following[username1] = helpers.RemoveFromSlice(globals.Following[username1], i)

	j := helpers.IndexOf(username1, globals.Followers[username2])
	globals.Followers[username2] = helpers.RemoveFromSlice(globals.Followers[username2], j)

	return true
}
