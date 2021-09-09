package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("location types", func() {
	Context("GetType", func() {
		It("returns the location type as a string", func() {
			lattype := bsareader.LatitudeType{Type: 12}

			Expect(lattype.GetType()).To(Equal("Graveyard"))
		})
	})
})
