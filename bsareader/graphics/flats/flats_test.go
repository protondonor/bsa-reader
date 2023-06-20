package flats_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/rowanjacobs/bsa-reader/bsareader/graphics/flats"
	"os"
	"path/filepath"
)

var _ = Describe("ReadFlats", func() {
	var flatsCfg []byte
	BeforeEach(func() {
		var err error
		// TODO: use a fixture instead
		flatsCfg, err = os.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "FLATS.CFG"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns flats", func() {
		flats := flats.ReadFlats(flatsCfg)
		Expect(flats).To(HaveLen(214))
		Expect(flats[0].Texture.File).To(Equal(175))
		Expect(flats[0].Texture.Index).To(Equal(0))
		Expect(flats[0].Description).To(Equal("beautiful maiden"))
		Expect(flats[0].Gender.Gender).To(Equal(2))
		Expect(flats[0].Gender.Obscene).To(BeTrue())
		Expect(flats[0].FaceIndex).To(Equal(410))
	})
})