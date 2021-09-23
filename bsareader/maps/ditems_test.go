package maps_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
	"io/ioutil"

	. "github.com/onsi/gomega"
)

var _ = Describe("dItems", func() {
	var cybiades []byte
	BeforeEach(func() {
		var err error
		cybiades, err = ioutil.ReadFile("fixtures/cybiades.ditem")
		Expect(err).NotTo(HaveOccurred())
	})

	It("reads offsets", func() {
		/*
		 * Each MapDItem record begins with an UInt32 count (DungeonCount)
		 * of DungeonInterior elements contained within the record.
		 */
		dItems := maps.ReadDItems(cybiades)

		Expect(dItems.Offsets).To(HaveLen(1))
		// I feel like these are in reverse order but idk
		Expect(dItems.Offsets[0].Offset).To(BeEquivalentTo(0))
		Expect(dItems.Offsets[0].DungeonObjId).To(BeEquivalentTo(1))
	})

	It("reads dungeon records", func() {
		dItems := maps.ReadDItems(cybiades)

		Expect(dItems.Items).To(HaveLen(1))
		Expect(dItems.Items[0].LRE.Doors).To(HaveLen(8))
		Expect(dItems.Items[0].Header.BlockCount).To(BeEquivalentTo(12))
		Expect(dItems.Items[0].Blocks[0].BlockInfo()).To(Equal("B0000007.RDB"))
		Expect(dItems.Items[0].Blocks[1].BlockInfo()).To(Equal("B0000008.RDB"))
		Expect(dItems.Items[0].Blocks[2].BlockInfo()).To(Equal("B0000003.RDB"))
		Expect(dItems.Items[0].Blocks[3].BlockInfo()).To(Equal("N0000042.RDB"))
	})

	Context("reading multiple records", func() {
		var betony []byte
		BeforeEach(func() {
			var err error
			betony, err = ioutil.ReadFile("fixtures/betony.ditem")
			Expect(err).NotTo(HaveOccurred())
		})

		It("reads dungeon records", func() {
			dItems := maps.ReadDItems(betony)

			Expect(dItems.Items).To(HaveLen(8))
			Expect(dItems.Items[0].LRE.Header.Name).To(Equal("Tristore Laboratory"))
			Expect(dItems.Items[0].Header.BlockCount).To(BeEquivalentTo(12))
			Expect(dItems.Items[1].LRE.Header.Name).To(Equal("The Stronghold of Hearthhouse"))
			Expect(dItems.Items[1].Header.BlockCount).To(BeEquivalentTo(10))
			Expect(dItems.Items[2].LRE.Header.Name).To(Equal("The Cabal of Gwynona"))
			Expect(dItems.Items[2].Header.BlockCount).To(BeEquivalentTo(11))
			Expect(dItems.Items[3].LRE.Header.Name).To(Equal("Yeomcroft's Hold"))
			Expect(dItems.Items[3].Header.BlockCount).To(BeEquivalentTo(8))
		})

		Context("GetBlocks", func() {
			It("can retrieve blocks given a dungeon name", func() {
				dItems := maps.ReadDItems(betony)

				blocks := dItems.GetBlocks("The Yeomhouse Crypts")
				Expect(blocks).To(HaveLen(5))
				Expect(blocks).To(ConsistOf("B0000009.RDB", "B0000006.RDB", "M0000008.RDB", "B0000013.RDB", "B0000001.RDB"))
			})
		})
	})
})
