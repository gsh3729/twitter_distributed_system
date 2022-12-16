package follow

import (
	"testing"
	"log"
	"context"
	"google.golang.org/grpc"

	// . "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
	// . "."
)

// func TestFollowers(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Testing followers")
// }

// var _ = Describe("Tweet", func() {
// 	Context("when a user connects to another user", func() {

// 	})

// })

// func TestFollowing(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Testing following")
// }

// var _ = Describe("Tweet", func() {
// 	Context("when a user connects to another user", func() {

// 	})

// })

func TestFollow(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username1 string = "user1"
	var username2 string = "user2"

	follow_server := NewFollowServiceClient(conn)
	response, err := follow_server.Follow(context.Background(), &FollowRequest{
		User1: username1,
		User2: username2,
	})

	if err != nil {
		t.Fatalf("TestFollow failed: %v", err)
	}

	if !response.Success {
		t.Error("TestFollow Failed")
	}

	// check the contents of followers
	resp, err := follow_server.GetUserFollowers(context.Background(), &GetFollowersRequest{
		Username: username2,
	})

	if err != nil {
		t.Fatalf("TestFollow failed: %v", err)
	}

	if !resp.Success {
		t.Error("TestFollow Failed")
	}

	for _, v := range resp.Users {
		if v == username1 {
			log.Printf("Follow working successfully")
		}
	}

	log.Printf("Follow not working successfully")
}

func TestUnfollow(t *testing.T) {

}




