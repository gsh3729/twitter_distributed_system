package tweet

import (
	context "context"
	"encoding/json"
	"sort"
	"time"

	globals "backend/globals"
	"backend/helpers"
)

type Tweet struct {
	Time string
	Text string
	User string
}

type Server struct {
	TweetServiceServer
}

func (s *Server) GetTweets(ctx context.Context, in *GetTweetsRequest) (*GetTweetsResponse, error) {
	var feed []Tweet
	var tweettexts []string
	var tweetowners []string
	var tweettimestamp []string

	tweets := make(map[string][]Tweet)
	resp := helpers.GetValueForKey("tweets")
	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &tweets)
	}

	for _, tweet := range tweets[in.Username] {
		feed = append(feed, tweet)
		tweettexts = append(tweettexts, tweet.Text)
		tweetowners = append(tweetowners, tweet.User)
		tweettimestamp = append(tweettimestamp, tweet.Time.Format("2006-01-02 15:04:05"))
	}
	following := helpers.GetMap("following")
	for _, element := range following[in.Username] {
		for _, tweet := range globals.Tweets[element] {
			feed = append(feed, tweet)
			tweettexts = append(tweettexts, tweet.Text)
			tweetowners = append(tweetowners, tweet.User)
			tweettimestamp = append(tweettimestamp, tweet.Time.Format("2006-01-02 15:04:05"))
		}
	}

	sort.SliceStable(feed[:], func(i, j int) bool {
		return feed[i].Time.Before(feed[j].Time)
	})

	return &GetTweetsResponse{Time: tweettimestamp, Text: tweettexts, User: tweetowners, Success: true}, nil
}

func (s *Server) PostTweet(ctx context.Context, in *PostTweetRequest) (*PostTweetResponse, error) {
	tweet := globals.Tweet{
		Time: time.Now(),
		Text: in.Text,
		User: in.Username,
	}

	globals.Tweets[in.Username] = append(globals.Tweets[in.Username], tweet)

	return &PostTweetResponse{Success: true}, nil
}
