package homepage

import (
	"sort"

	globals "backend/globals"
)

func GetTweetsForHomepage(username string) []globals.Tweet {
	var feed []globals.Tweet

	for _, tweet := range globals.Tweets[username] {
		feed = append(feed, tweet)
	}

	for _, element := range globals.Following[username] {
		for _, tweet := range globals.Tweets[element] {
			feed = append(feed, tweet)
		}
	}

	sort.SliceStable(feed[:], func(i, j int) bool {
		return feed[i].Time.Before(feed[j].Time)
	})
	return feed
}
