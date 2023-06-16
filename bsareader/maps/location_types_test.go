package maps_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
)

var _ = Describe("location types", func() {
	Context("GetType", func() {
		It("returns the location type as a string", func() {
			latType := maps.LatitudeType{Type: 12}

			Expect(latType.GetType()).To(Equal("Graveyard"))
		})
	})
})
