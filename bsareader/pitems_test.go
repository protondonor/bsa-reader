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

	It("reads exterior data", func() {
		pItems := bsareader.ReadPItems(cybiades, 1)

		Expect(pItems.Items[0].Exterior.Name).To(Equal("Ruins of Cosh Hall"))
		Expect(pItems.Items[0].Exterior.MapId.LocationExterior).To(BeEquivalentTo(0x53D8F))
		Expect(pItems.Items[0].Exterior.Width).To(BeEquivalentTo(1))
		Expect(pItems.Items[0].Exterior.Height).To(BeEquivalentTo(1))
		Expect(pItems.Items[0].Exterior.BlockIndex[0]).To(BeEquivalentTo(21)) // 21: RUIN
		Expect(pItems.Items[0].Exterior.BlockNumber[0]).To(BeEquivalentTo(11))
		Expect(pItems.Items[0].Exterior.BlockChar[0]).To(BeEquivalentTo(15))
		Expect(pItems.Items[0].Exterior.Encounters).To(BeEquivalentTo(2))
		Expect(pItems.Items[0].Exterior.DungeonName).To(Equal("Ruins of Cosh Hall"))
		Expect(pItems.Items[0].Exterior.BlockCount).To(BeEquivalentTo(12))
		Expect(pItems.Items[0].Exterior.Services).To(BeEquivalentTo(0))
	})

	It("reads building data", func() {
		pItems := bsareader.ReadPItems(betony, 25)

		Expect(pItems.Items[3].LRE.Header.Name).To(Equal("Mastersley Orchard"))
		Expect(pItems.Items[3].Buildings).To(HaveLen(1))
		Expect(pItems.Items[3].Buildings[0].NameSeed).To(BeEquivalentTo(0x3222))
		Expect(pItems.Items[3].Buildings[0].FactionId).To(BeEquivalentTo(0))
		Expect(pItems.Items[3].Buildings[0].ObjectId).To(BeEquivalentTo(0x583a0004))
		Expect(pItems.Items[3].Buildings[0].Type).To(BeEquivalentTo(0x12))
		Expect(pItems.Items[3].Buildings[0].Quality).To(BeEquivalentTo(3))
	})

	It("reads multiple elements", func() {
		pItems := bsareader.ReadPItems(betony, 25)

		Expect(pItems.Items[2].LRE.Header.Name).To(Equal("Tristore Laboratory"))
	})
})
