package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("bsareader", func() {
	// TODO: property testing for headers
	It("reads BSA headers", func() {
		bsa := []byte{0, 8, 1, 0, 5}

		header := bsareader.ReadHeader(bsa)

		Expect(header.RecordCount).To(BeEquivalentTo(8))
		Expect(header.Type).To(BeEquivalentTo(bsareader.NameRecord))
	})

	Context("getting the offset for the BSA footer", func() {
		It("works with NameRecord", func() {
			offset := bsareader.GetFooterOffset(uint16(23), bsareader.NameRecord)

			Expect(offset).To(BeEquivalentTo(414))
		})

		It("works with NumberRecord", func() {
			offset := bsareader.GetFooterOffset(uint16(17), bsareader.NumberRecord)

			Expect(offset).To(BeEquivalentTo(136))
		})
	})

	Context("reading BSA footer", func() {
		Context("NameRecords", func() {
			It("reads multiple records", func() {
				bsa := []byte{77, 65, 80, 78, 65, 77, 69, 83, 46, 48, 48, 53, 0, 0, 1, 0, 0, 0,
					77, 65, 80, 80, 73, 84, 69, 77, 46, 48, 49, 55, 0, 0, 0, 0, 0, 1,
				}

				footer := bsareader.ReadFooter(bsa, bsareader.NameRecord)
				Expect(len(footer.Records)).To(Equal(2))
				Expect(footer.Records[0].Name).To(Equal("MAPNAMES.005"))
				Expect(footer.Records[0].Size).To(BeEquivalentTo(16777216))
				Expect(footer.Records[1].Name).To(Equal("MAPPITEM.017"))
				Expect(footer.Records[1].Size).To(BeEquivalentTo(1))
			})
		})

		Context("NumberRecords", func() {
			It("reads multiple records", func() {
				bsa := []byte{1, 2, 0, 0, 1, 0, 0, 0,
					3, 4, 0, 0, 0, 0, 0, 1,
				}

				footer := bsareader.ReadFooter(bsa, bsareader.NumberRecord)
				Expect(len(footer.Records)).To(Equal(2))
				Expect(footer.Records[0].Name).To(Equal("258"))
				Expect(footer.Records[0].Size).To(BeEquivalentTo(16777216))
				Expect(footer.Records[1].Name).To(Equal("772"))
				Expect(footer.Records[1].Size).To(BeEquivalentTo(1))
			})
		})
	})
})
