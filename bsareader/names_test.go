package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("names", func() {
	bsa := []byte{
		1, 0, 0, 0, 0x52, 0x75, 0x69, 0x6E, 0x73, 0x20, 0x6F, 0x66,
		0x20, 0x43, 0x6F, 0x73, 0x68, 0x20, 0x48, 0x61, 0x6C, 0x6C,
		0, 0x4C, 0x6F, 0x64, 0x67, 0x65, 0, 0x6E, 0, 0x72, 0x6E,
		0, 0, 0,
	}

	It("reads the count of name records", func() {
		mapNames := bsareader.ReadNames(bsa)
		Expect(mapNames.LocationCount).To(BeEquivalentTo(1))
	})

	It("reads the name records", func() {
		mapNames := bsareader.ReadNames(bsa)
		Expect(mapNames.Names).To(ConsistOf("Ruins of Cosh Hall"))
	})
})
