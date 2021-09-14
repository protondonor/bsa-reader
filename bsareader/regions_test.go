package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("Regions", func() {
	Context("ParseRegions", func() {
		It("returns 3 digit region codes", func() {
			Expect(bsareader.ParseRegion("Cybiades")).To(Equal("061"))
			Expect(bsareader.ParseRegion("\tcYBIAdES   ")).To(Equal("061"))
			Expect(bsareader.ParseRegion("61")).To(Equal("061"))
		})
	})
})
