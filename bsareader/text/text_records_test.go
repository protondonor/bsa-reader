package text_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/rowanjacobs/bsa-reader/bsareader/text"
	"os"
	"path/filepath"
)

var _ = Describe("ReadClass", func() {
	var textRsc []byte
	BeforeEach(func() {
		var err error
		// TODO: use a fixture instead
		textRsc, err = os.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "TEXT.RSC"))
		Expect(err).NotTo(HaveOccurred())
		Expect(len(textRsc)).To(Equal(353393))
	})

	It("reads text", func() {
		textRecordDatabase := text.ReadTextRecord(textRsc)

		Expect(textRecordDatabase.Length).To(Equal(uint16(1569)))
		Expect(len(textRecordDatabase.Headers)).To(Equal(1569))
		Expect(textRecordDatabase.Headers[0].Text).To(ContainSubstring("STRENGTH"))
		Expect(textRecordDatabase.Headers[0].Text).To(ContainSubstring("Strength governs encumbrance, weapon damage"))
		Expect(textRecordDatabase.Headers[0].Text).To(ContainSubstring("and the ease of increasing strength-related skills."))
		Expect(textRecordDatabase.Headers[0].Text).To(ContainSubstring("With your strength of %str, you are considered %ark"))
		Expect(textRecordDatabase.Headers[0].Text).To(ContainSubstring("(modifier is factored into your"))
		Expect(textRecordDatabase.Headers[0].Text).To(ContainSubstring("(hand-to-hand / weapon damage."))

		println(textRecordDatabase.Headers[1407].Text)
	})
})
