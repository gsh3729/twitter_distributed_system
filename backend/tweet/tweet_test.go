package tweet

import (
	"testing"
	"log"
	"context"
	"google.golang.org/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
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

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:12380", "localhost:22380", "localhost:32380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Get(ctx, "tweets")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)

	// convert response to list of struct 

	// rep, err := tweet_server.GetTweets(context.Background(), &GetTweetsRequest{
	// 	Username: username,
	// })

	log.Printf("Posted a new tweet successfully")
}