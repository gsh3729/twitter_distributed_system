package tweet

import (
	"testing"
	"log"
	"context"
	"google.golang.org/grpc"
	// . "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	// . "."
)

// func TestTweeting(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Posting a tweet")
// }

// var _ = Describe("Tweet", func() {
// 	Context("when a tweet is posted", func() {
		
// 	})

// })

func TestTweeting(t *testing.T) {
	// ctx := context.Background()
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	tweet_server := NewTweetServiceClient(conn)
	response, err := tweet_server.PostTweet(context.Background(), &PostTweetRequest{
		Username: "harsha",
		Text:     "Hi, this tweet is from harsha",
	})

	if err != nil {
		t.Fatalf("TestTweeting failed: %v", err)
	}

	if !response.Success {
		t.Error("TestTweeting Failed")
	}

	

	log.Printf("Posted a new tweet successfully")
}