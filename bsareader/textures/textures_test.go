package textures_test

import (
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
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
		var texturesFile []byte
		BeforeEach(func() {
			var err error
			// TODO: use a fixture instead
			texturesFile, err = ioutil.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "TEXTURE.003"))
			Expect(err).NotTo(HaveOccurred())
		})
		It("returns textures", func() {
			txts := textures.ReadTextures(texturesFile)
			Expect(txts.Header.Count).To(BeEquivalentTo(56))
			Expect(txts.Header.Name).To(Equal(" Desert Terrain Set Wint"))
			Expect(txts.RecordPointers).To(HaveLen(56))
			Expect(txts.TextureRecords).To(HaveLen(56))

			tr := txts.TextureRecords[0]
			Expect(tr.Width).To(BeEquivalentTo(64))
			Expect(tr.Height).To(BeEquivalentTo(64))
			Expect(tr.CompressionType()).To(Equal("Uncompressed"))
			Expect(tr.FrameCount).To(BeEquivalentTo(1))
		})

		Context("textures", func() {
			It("contain image data", func() {
				txts := textures.ReadTextures(texturesFile)
				tr := txts.TextureRecords[0]
				img := tr.Uncompress(texturesFile)
				Expect(img).To(HaveLen(64))
				Expect(img[0]).To(HaveLen(64))
			})
		})
	})
})
