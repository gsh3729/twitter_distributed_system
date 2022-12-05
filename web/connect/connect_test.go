package connect

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// . "."
) 

func TestConnect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Connecting one user to another user")
}

var _ = Describe("Connect", func() {
	Context("when a connection is requested", func() {
		
	})

})