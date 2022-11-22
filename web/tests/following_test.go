package tests

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "."
)

func TestFollowing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing following")
}

var _ = Describe("Tweet", func() {
	Context("when a tweet is posted", func() {
		
	})

})