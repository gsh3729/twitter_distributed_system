package follow

import (
	"testing"
	"log"
	"context"
	"google.golang.org/grpc"

)

func TestFollow(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username1 string = "sri"
	var username2 string = "har"

	follow_server := NewFollowServiceClient(conn)
	response, err := follow_server.Follow(context.Background(), &FollowRequest{
		User1: username1,
		User2: username2,
	})

	if err != nil && !response.Success {
		t.Error("TestFollow failed: ", err)
	} 

	resp, err := follow_server.GetUserFollowers(context.Background(), &GetFollowersRequest{
		Username: username2,
	})

	if err != nil && !resp.Success {
		t.Error("TestFollow failed: ", err)
	}


	for _, v := range resp.Users {
		if v == username1 {
			log.Printf("Follow working successfully")
		}
	}

	log.Printf("Follow not working successfully")
}

func TestUnfollow(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()

	var username1 string = "user1"
	var username2 string = "user2"

	follow_server := NewFollowServiceClient(conn)
	response, err := follow_server.Unfollow(context.Background(), &UnfollowRequest{
		User1: username1,
		User2: username2,
	})

	if err != nil && !response.Success {
		t.Error("TestUnfollow failed: ", err)
	} 

	resp, err := follow_server.GetUserFollowers(context.Background(), &GetFollowersRequest{
		Username: username2,
	})

	if err != nil && !resp.Success {
		t.Error("TestUnfollow failed: ", err)
	}

	for _, v := range resp.Users {
		if v == username1 {
			t.Error("Unfollow not working successfully")
		}
	}

	log.Printf("Follow working successfully")
}




