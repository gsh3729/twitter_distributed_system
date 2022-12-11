package tweet

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "."
)

func TestTweeting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Posting a tweet")
}

var _ = Describe("Tweet", func() {
	Context("when a tweet is posted", func() {
		
	})

})