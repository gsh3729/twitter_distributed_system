package cart_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "."
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {