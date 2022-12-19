package tweet

import (
	"testing"
	"log"
	"context"
	"google.golang.org/grpc"
	"backend/helpers"
	"encoding/json"
)

func TestTweeting(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username string = "sri"
	var tweet_text string = "Hi, this tweet is from sri"

	tweet_server := NewTweetServiceClient(conn)
	response, err := tweet_server.PostTweet(context.Background(), &PostTweetRequest{
		Username: username,
		Text:     tweet_text,
	})

	if err != nil || !response.Success{
		t.Error("TestTweeting failed: ", err)
	}

	tweets := make(map[string][]Tweet)
	resp := helpers.GetValueForKey("tweets")
	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &tweets)
	}

	var flag bool = true
	for _, v := range tweets[username] {
		if v.User == username && v.Text == tweet_text {
			flag = false
			log.Printf("Tweet tests passed successfully")
		}
	}
	if flag {
		t.Error("TestTweeting failed")
	}
}