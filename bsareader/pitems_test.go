package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	"io/ioutil"

	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("pItems", func() {
	var betony, cybiades []byte
	BeforeEach(func() {
		var err error
		betony, err = ioutil.ReadFile("fixtures/betony.pitem")
		Expect(err).NotTo(HaveOccurred())
		cybiades, err = ioutil.ReadFile("fixtures/cybiades.pitem")
		Expect(err).NotTo(HaveOccurred())
	})

	It("reads offsets", func() {
		/*
			Each MapPItem record begins with a list of UInt32 values with
			MapNames.LocationCount elements. This is a list of offsets to the
			LocationExterior structures for each location. These values are
			relative to the end of the list, so adding (LocationCount * 4) to
			the offset locates the actual record. One offset may be greater
			than a subsequent one; no data order is presumed.

			At each offset indicated, a LocationExterior structure is present,
			which includes all fields from the LocationRecordElement and some
			data of its own.
		*/
		pItems := bsareader.ReadPItems(cybiades, 1)

		Expect(pItems.Offsets).To(HaveLen(1))
		Expect(pItems.Offsets[0]).To(BeEquivalentTo(0))
	})

	It("reads location record elements", func() {
		pItems := bsareader.ReadPItems(cybiades, 1)

		Expect(pItems.Items[0].LRE.Doors).To(HaveLen(0))
		Expect(pItems.Items[0].LRE.Header.Name).To(Equal("Ruins of Cosh Hall"))
	})

	It("reads multiple elements", func() {
		pItems := bsareader.ReadPItems(betony, 25)

		Expect(pItems.Items[2].LRE.Header.Name).To(Equal("Tristore Laboratory"))
	})
})
