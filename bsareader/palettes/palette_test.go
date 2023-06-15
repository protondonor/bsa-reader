package palettes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/rowanjacobs/bsa-reader/bsareader/palettes"
	"io/ioutil"
	"path/filepath"
)

var _ = Describe("ReadPal", func() {

	It("reads RGB values from a palette", func() {
		palBytes, err := ioutil.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "PAL.RAW"))
		Expect(err).NotTo(HaveOccurred())

		palette := palettes.ReadPalette(palBytes)
		Expect(palette).To(HaveLen(256))
		Expect(palette[0].Red).To(BeEquivalentTo(0))
		Expect(palette[0].Green).To(BeEquivalentTo(0))
		Expect(palette[0].Blue).To(BeEquivalentTo(0))
		Expect(palette[1].Red).To(BeEquivalentTo(0x3F))
		Expect(palette[1].Green).To(BeEquivalentTo(0x39))
		Expect(palette[1].Blue).To(BeEquivalentTo(0x20))
	})

	It("reads RGB values from a .COL file", func() {
		colBytes, err := ioutil.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "PAL.PAL"))
		Expect(err).NotTo(HaveOccurred())

		palette := palettes.ReadCol(colBytes)
		Expect(palette).To(HaveLen(256))
		Expect(palette[0].Red).To(BeEquivalentTo(0))
		Expect(palette[0].Green).To(BeEquivalentTo(0))
		Expect(palette[0].Blue).To(BeEquivalentTo(0))
		Expect(palette[1].Red).To(BeEquivalentTo(0xFF))
		Expect(palette[1].Green).To(BeEquivalentTo(0))
		Expect(palette[1].Blue).To(BeEquivalentTo(0xFF))
	})
})
