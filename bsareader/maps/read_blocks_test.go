package maps_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/rowanjacobs/bsa-reader/bsareader/maps"
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/gomega"
)

var _ = Describe("ReadBlocks", func() {
	var bsa []byte
	BeforeEach(func() {
		var err error
		// TODO: use a fixture instead
		bsa, err = ioutil.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "MAPS.BSA"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns blocks!", func() {
		blocks := maps.ReadBlocks(bsa, "Ruins of Cosh Hall", "061")
		Expect(blocks).To(HaveLen(12))
	})
})
