package tweet

import (
	"testing"

	"context"
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
	ctx := context.Background()
	
	defer conn.Close()

	tweet_server := NewTweetServiceClient(conn)
	response, err := tweet_server.PostTweet(context.Background(), &PostTweetRequest{
		Username: user.(string),
		Text:     tweetMsg,
	})

}