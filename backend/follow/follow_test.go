package follow

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "."
)

func TestFollowers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing followers")
}

var _ = Describe("Tweet", func() {
	Context("when a user connects to another user", func() {

	})

})

func TestFollowing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing following")
}

var _ = Describe("Tweet", func() {
	Context("when a user connects to another user", func() {

	})

})