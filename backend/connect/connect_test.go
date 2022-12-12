package connect_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "backend/connect"
)

func TestConnect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Connecting one user to another user")
}

var _ = Describe("Connect", func() {
	BeforeEach(func() {
		// book = &books.Book{
		//   Title: "Les Miserables",
		//   Author: "Victor Hugo",
		//   Pages: 2783,
		// }
		// Expect(book.IsValid()).To(BeTrue())
	})

	Context("when follow req is issued", func() {
		// It("interprets the single author name as a last name", func() {
		// 	Expect(book.AuthorLastName()).To(Equal("Hugo"))
		// })
		// Follow()
	})

	Context("when unfollow req is issued", func() {
		// Unfollow()
	})

})
