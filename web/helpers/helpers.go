package helpers

import (
	"log"
	"sort"
	"strings"

	globals "proj/web/globals"
)

func CheckUserPass(username, password string) bool {
	userpass := globals.UserPass

	log.Println("checkUserPass", username, password, userpass)

	if val, ok := userpass[username]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func RemoveFromSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

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
