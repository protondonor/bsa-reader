package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("bsareader", func() {
	It("reads BSA headers", func() {
		bsa := []byte{0, 8, 1, 0, 5}

		header := bsareader.ReadHeader(bsa)

		Expect(header.RecordCount).To(BeEquivalentTo(8))
		Expect(header.BsaType).To(BeEquivalentTo(bsareader.NameRecord))
	})
})
