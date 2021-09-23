package maps_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
)

var _ = Describe("tables", func() {
	bsa := []byte{
		0x8F, 0x3D, 0x05, 0x01, 0x00, 0xA8, 0xDB, 0x14, 0x01, 0x28, 0x4E, 0x33,
		0x02, 0x00, 0x00, 0x00, 0x00,
	}

	It("reads the map ID", func() {
		mapTable := maps.ReadTable(bsa)

		mapId := mapTable.Rows[0].MapId
		Expect(mapId.LocationExterior).To(BeEquivalentTo(0x53D8F))
		Expect(mapId.LocationIndex).To(BeEquivalentTo(0x10))
	})

	It("reads the latitude and location type", func() {
		mapTable := maps.ReadTable(bsa)

		latType := mapTable.Rows[0].LatitudeType
		Expect(latType.Latitude).To(BeEquivalentTo(0xdba800))
		Expect(latType.Type).To(BeEquivalentTo(0xa))
		Expect(latType.Discovered).To(BeFalse())
		Expect(latType.Hidden).To(BeFalse())
	})

	It("reads the longitude, width, and height", func() {
		mapTable := maps.ReadTable(bsa)

		longType := mapTable.Rows[0].LongitudeType
		Expect(longType.Longitude).To(BeEquivalentTo(0x4e2801))
		Expect(longType.Width).To(BeEquivalentTo(3))
		Expect(longType.Height).To(BeEquivalentTo(3))
	})

	It("reads dungeon flavor text index", func() {
		mapTable := maps.ReadTable(bsa)

		flavor := mapTable.Rows[0].Flavor
		Expect(flavor).To(BeEquivalentTo(2))
	})

	It("reads services bitfield", func() {
		mapTable := maps.ReadTable([]byte{0x8A, 0xE4, 0x13, 0x00, 0x00,
			0x20, 0x39, 0x44, 0x01, 0x28, 0x7A, 0x34, 0xFF, 0x00, 0x01, 0x00, 0x01})

		services := mapTable.Rows[0].Services
		Expect(services).To(BeEquivalentTo(0x1000100))
	})

	It("deals with multiple rows in the table", func() {
		bsa2 := []byte{
			0xAD, 0xE0, 0x03, 0x00, 0x00, 0xA8, 0x3E, 0x52, 0x01, 0xA8, 0x7A,
			0x33, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x8A, 0xE4, 0x13, 0x00, 0x00,
			0x20, 0x39, 0x44, 0x01, 0x28, 0x7A, 0x34, 0xFF, 0x00, 0x01, 0x00,
			0x01, 0x9C, 0xE4, 0x23, 0x00, 0x00, 0x28, 0x42, 0x08, 0x01, 0x28,
			0x7A, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00, 0x7D, 0xE8, 0x33, 0x00,
			0x00, 0xA8, 0x3E, 0x46, 0x01, 0xA8, 0x79, 0x33, 0xFF, 0x00, 0x00,
			0x00, 0x00,
		}

		mapTable := maps.ReadTable(bsa2)
		Expect(mapTable.Rows).To(HaveLen(4))
	})
})
