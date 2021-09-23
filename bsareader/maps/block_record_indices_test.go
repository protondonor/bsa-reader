package maps_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
)

var _ = Describe("BlockInfo", func() {
	Context("Exterior blocks", func() {
		Context("Default logic", func() {
			It("returns block string for block at count", func() {
				extData := maps.ExteriorData{
					BlockNumber: []byte{17},
					BlockIndex:  []byte{1},
					BlockChar:   []byte{1},
				}

				Expect(extData.BlockInfo(0)).To(Equal("GENRBA17"))
			})
		})

		Context("Temple", func() {
			It("uses special values for number and char", func() {
				extData := maps.ExteriorData{
					BlockNumber: []byte{17},
					BlockIndex:  []byte{13},
					BlockChar:   []byte{1},
				}

				Expect(extData.BlockInfo(0)).To(Equal("TEMPAAB0"))
			})
		})

		PContext("Wayrest", func() {})

		PContext("Sentinel", func() {})
	})

	Context("Dungeon blocks", func() {
		It("returns block string", func() {
			dBlock := maps.DungeonBlock{
				BlockNumber: 19,
				BlockIndex:  3,
			}

			Expect(dBlock.BlockInfo()).To(Equal("S0000019.RDB"))
		})
	})
})
