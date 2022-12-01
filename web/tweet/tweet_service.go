package tweet

import (
	// "log"
	// "strings"
	"time"

	globals "proj/web/globals"
)

func PostTweet(username string, tweetMsg string) bool {
	tweet := globals.Tweet{
		Time: time.Now(),
		Text: tweetMsg,
		User: username,
	}

	globals.Tweets[username] = append(globals.Tweets[username], tweet)

	return true
} 