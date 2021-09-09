package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"io/ioutil"
)

var _ = Describe("location record elements", func() {
	var bsa []byte
	BeforeEach(func() {
		var err error
		bsa, err = ioutil.ReadFile("fixtures/cybiades.pitem")
		Expect(err).NotTo(HaveOccurred())
		/* the first $locationCount dwords in a pItem record are a list
		 * of offsets to LocationExterior records.
		 */
		bsa = bsa[4:]
	})

	Describe("reading a single element", func() {
		It("reads object headers", func() {
			lre := bsareader.ReadLocationRecordElem(bsa)

			Expect(lre.ObjectHeader.Latitude).To(BeEquivalentTo(0x00dbb800))
			Expect(lre.ObjectHeader.Longitude).To(BeEquivalentTo(0x004E3801))
			Expect(lre.ObjectHeader.IsExterior).To(BeTrue())
			Expect(lre.ObjectHeader.ParentId & 0x000fffff).To(BeEquivalentTo(0))
			Expect(lre.ObjectHeader.ObjectId & 0x000fffff).To(BeEquivalentTo(0x40001))
		})

		It("reads LRE headers", func() {
			lre := bsareader.ReadLocationRecordElem(bsa)

			Expect(lre.Header.Name).To(Equal("Ruins of Cosh Hall"))
			Expect(lre.Header.Width).To(BeEquivalentTo(1))
			Expect(lre.Header.Height).To(BeEquivalentTo(1))
			Expect(lre.Header.LocationType).To(BeEquivalentTo(10))
			Expect(lre.Header.BuildingCount).To(BeEquivalentTo(0))
		})
	})

	//PIt("reads doors", func() {
	//	var err error
	//	bsa, err = ioutil.ReadFile("fixtures/betony.pitem")
	//	Expect(err).NotTo(HaveOccurred())
	//
	//	lres := bsareader.ReadLocationRecords(bsa)
	//
	//	Expect(len(lres[0].Doors)).To(BeEquivalentTo(0))
	//	Expect(len(lres[1].Doors)).To(BeEquivalentTo(0))
	//	Expect(len(lres[2].Doors)).To(BeEquivalentTo(1))
	//	Expect(lres[2].Doors[0].BuildingDataIndex).To(BeEquivalentTo(17))
	//})
})