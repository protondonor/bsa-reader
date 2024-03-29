package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("bsareader", func() {
	bsa := []byte{2, 0, 0, 1,
		72, 101, 108, 108, 111,
		119, 111, 114, 108, 100, 33,
		77, 65, 80, 78, 65, 77, 69, 83, 46, 48, 48, 53, 0, 0, 5, 0, 0, 0,
		77, 65, 80, 80, 73, 84, 69, 77, 46, 48, 49, 55, 0, 0, 6, 0, 0, 0,
	}
	numRecordBsa := []byte{2, 1, 0, 0, 1, 0, 0, 0,
		4, 3, 0, 0, 0, 0, 0, 1,
	}

	// TODO: property testing for headers
	It("reads BSA headers", func() {
		bsa := []byte{8, 0, 0, 1, 5}

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
				footer := bsareader.ReadFooter(bsa[15:], bsareader.NameRecord)
				Expect(len(footer)).To(Equal(2))
				Expect(footer[0].Name).To(Equal("MAPNAMES.005"))
				Expect(footer[0].Size).To(BeEquivalentTo(5))
				Expect(footer[1].Name).To(Equal("MAPPITEM.017"))
				Expect(footer[1].Size).To(BeEquivalentTo(6))
			})
		})

		Context("NumberRecords", func() {
			It("reads multiple records", func() {
				footer := bsareader.ReadFooter(numRecordBsa, bsareader.NumberRecord)
				Expect(len(footer)).To(Equal(2))
				Expect(footer[0].Name).To(Equal("258"))
				Expect(footer[0].Size).To(BeEquivalentTo(1))
				Expect(footer[1].Name).To(Equal("772"))
				Expect(footer[1].Size).To(BeEquivalentTo(16777216))
			})
		})
	})

	It("reads BSA files", func() {
		records := bsareader.Read(bsa)
		Expect(len(records)).To(Equal(2))

		Expect(records[0].Contents).To(Equal([]byte("Hello")))
		Expect(records[1].Contents).To(Equal([]byte("world!")))
	})

	It("lists the contents of BSA files", func() {
		records := bsareader.List(bsa)
		Expect(len(records)).To(Equal(2))

		Expect(records[0].Contents).To(Equal([]byte{}))
	})

	It("returns a specific record from a BSA file", func() {
		record := bsareader.ReadRecord(bsa, "MAPNAMES.005")
		Expect(record.Contents).To(Equal([]byte("Hello")))

		record = bsareader.ReadRecord(bsa, "MAPPITEM.017")
		Expect(record.Contents).To(Equal([]byte("world!")))
	})
})
