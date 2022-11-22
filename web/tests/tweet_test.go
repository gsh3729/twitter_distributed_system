package cart_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Posting a tweet")
}

var _ = Describe("Tweet", func() {
	Context("initially", func() {
		cart := Cart{}

		It("has 0 items", func() {
			Expect(cart.TotalUniqueItems()).Should(BeZero())
		})

		It("has 0 units", func() {
			Expect(cart.TotalUnits()).Should(BeZero())
		})

		Specify("the total amount is 0.00", func() {
			Expect(cart.TotalAmount()).Should(BeZero())
		})
	})
})