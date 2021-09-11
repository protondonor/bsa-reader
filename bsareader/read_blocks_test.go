package bsareader_test

import (
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

// idk if it's cool for me to redistribute MAPS.BSA with this repo
// so this is what we're going to do instead
const DAGGERFALL_PATH = "/home/rowan/games/abandon/dfall/arena2"

var _ = Describe("ReadBlocks", func() {
	var maps []byte
	BeforeEach(func() {
		var err error
		maps, err = ioutil.ReadFile(filepath.Join(DAGGERFALL_PATH, "MAPS.BSA"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns blocks!", func() {
		blocks := bsareader.ReadBlocks(maps, "Ruins of Cosh Hall", "061")
		Expect(blocks).To(HaveLen(12))
	})
})
