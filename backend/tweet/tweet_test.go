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

	var username string = "harshaG"
	var tweet_text string = "Hi, this tweet is from harsha"

	tweet_server := NewTweetServiceClient(conn)
	response, err := tweet_server.PostTweet(context.Background(), &PostTweetRequest{
		Username: username,
		Text:     tweet_text,
	})

	if err != nil {
		t.Fatalf("TestTweeting failed: %v", err)
	}

	if !response.Success {
		t.Error("TestTweeting Failed")
	}

	// check the content is posted or not

	tweets := make(map[string][]Tweet)
	resp := helpers.GetValueForKey("tweets")
	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &tweets)
	}

	// convert response to list of struct 
	// var feed []globals.Tweet

	// for _, v := range tweets {
	// 	if v == username1 {
	// 		log.Printf("Follow working successfully")
	// 	}
	// }

	log.Printf("Posted a new tweet successfully")
}