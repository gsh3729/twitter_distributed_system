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

func TestFollowers(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err2 := grpc.Dial(":9000", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Couldn't connect: %s", err2)
	}
	defer conn.Close()
}

func TestFollowing(t *testing.T) {

}


