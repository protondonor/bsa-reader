package palettes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/rowanjacobs/bsa-reader/bsareader/palettes"
	"os"
	"path/filepath"
)

var _ = Describe("ReadPal", func() {

	It("reads RGB values from a palette", func() {
		palBytes, err := os.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "PAL.RAW"))
		Expect(err).NotTo(HaveOccurred())

		palette := palettes.ReadPalette(palBytes)
		Expect(palette).To(HaveLen(256))
		r, g, b, _ := palette[0].RGBA()
		Expect(r).To(BeEquivalentTo(0))
		Expect(g).To(BeEquivalentTo(0))
		Expect(b).To(BeEquivalentTo(0))

		r, g, b, _ = palette[1].RGBA()
		// multiply by 0x101 because RGBA() returns alpha-premultiplied values
		// see https://go.dev/blog/image
		Expect(r).To(BeEquivalentTo(0x3F*0x101))
		Expect(g).To(BeEquivalentTo(0x39*0x101))
		Expect(b).To(BeEquivalentTo(0x20*0x101))
	})

	It("sets the alpha of the first color in the palette to 0 and all others to full", func() {
		palBytes, err := os.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "PAL.RAW"))
		Expect(err).NotTo(HaveOccurred())

		palette := palettes.ReadPalette(palBytes)
		Expect(palette).To(HaveLen(256))
		_, _, _, a := palette[0].RGBA()
		Expect(a).To(BeEquivalentTo(0))

		for i := 1; i < 256; i++ {
			_, _, _, a = palette[i].RGBA()
			Expect(a).To(BeEquivalentTo(65535))
		}
	})

	It("reads RGB values from a .COL file", func() {
		colBytes, err := os.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "PAL.PAL"))
		Expect(err).NotTo(HaveOccurred())

		palette := palettes.ReadCol(colBytes)
		Expect(palette).To(HaveLen(256))
		r, g, b, a := palette[0].RGBA()
		Expect(r).To(BeEquivalentTo(0))
		Expect(g).To(BeEquivalentTo(0))
		Expect(b).To(BeEquivalentTo(0))
		Expect(a).To(BeEquivalentTo(0))

		r, g, b, a = palette[1].RGBA()
		// multiply by 0x101 because RGBA() returns alpha-premultiplied values
		// see https://go.dev/blog/image
		Expect(r).To(BeEquivalentTo(0xFF*0x101))
		Expect(g).To(BeEquivalentTo(0))
		Expect(b).To(BeEquivalentTo(0xFF*0x101))
		Expect(a).To(BeEquivalentTo(0xFF*0x101))
	})
})
