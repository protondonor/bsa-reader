package textures_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader/textures"
)

var _ = Describe("Textures", func() {

	Describe("CompressionType", func() {
		Context("value is 0", func() {
			It("returns Uncompressed", func() {
				tr := textures.TextureRecord{Compression: 0}
				Expect(tr.CompressionType()).To(Equal("Uncompressed"))
			})
		})

		Context("value is 2", func() {
			It("returns RleCompressed", func() {
				tr := textures.TextureRecord{Compression: 2}
				Expect(tr.CompressionType()).To(Equal("RleCompressed"))
			})
		})

		Context("value is 0x0108", func() {
			It("returns ImageRle", func() {
				tr := textures.TextureRecord{Compression: 0x0108}
				Expect(tr.CompressionType()).To(Equal("ImageRle"))
			})
		})

		Context("value is 0x1108", func() {
			It("returns RecordRle", func() {
				tr := textures.TextureRecord{Compression: 0x1108}
				Expect(tr.CompressionType()).To(Equal("RecordRle"))
			})
		})
	})

	Describe("ReadTextures", func() {
		// var flatsCfg []byte
		// BeforeEach(func() {
		// var err error
		// TODO: use a fixture instead
		// flatsCfg, err = ioutil.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "FLATS.CFG"))
		// Expect(err).NotTo(HaveOccurred())
		// })
		// It("returns flats", func() {
		// 	flats := flats.ReadFlats(flatsCfg)
		// 	Expect(flats).To(HaveLen(214))
		// 	Expect(flats[0].Texture.File).To(Equal(175))
		// 	Expect(flats[0].Texture.Index).To(Equal(0))
		// 	Expect(flats[0].Description).To(Equal("beautiful maiden"))
		// 	Expect(flats[0].Gender.Gender).To(Equal(2))
		// 	Expect(flats[0].Gender.Obscene).To(BeTrue())
		// 	Expect(flats[0].FaceIndex).To(Equal(410))
		// })
	})
})
