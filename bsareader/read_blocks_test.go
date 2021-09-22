package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

var _ = Describe("ReadBlocks", func() {
	var maps []byte
	BeforeEach(func() {
		var err error
		// TODO: use a fixture instead
		maps, err = ioutil.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "MAPS.BSA"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns blocks!", func() {
		blocks := bsareader.ReadBlocks(maps, "Ruins of Cosh Hall", "061")
		Expect(blocks).To(HaveLen(12))
	})
})
