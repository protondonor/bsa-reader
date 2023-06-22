package classes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rowanjacobs/bsa-reader/bsareader"
	"github.com/rowanjacobs/bsa-reader/bsareader/classes"
	"os"
	"path/filepath"
)

var _ = Describe("ReadClass", func() {
	var classCfg []byte
	BeforeEach(func() {
		var err error
		// TODO: use a fixture instead
		classCfg, err = os.ReadFile(filepath.Join(bsareader.GetDaggerfallPath(), "CLASS00.CFG"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns a class", func() {
		class := classes.ReadClass(classCfg)

		Expect(class.Name).To(Equal("Mage"))
		Expect(class.HP).To(BeEquivalentTo(6))
		Expect(class.Attributes.Strength).To(BeEquivalentTo(42))
		Expect(class.Attributes.Intelligence).To(BeEquivalentTo(60))
		Expect(class.Attributes.Willpower).To(BeEquivalentTo(65))
		Expect(class.Attributes.Agility).To(BeEquivalentTo(45))
		Expect(class.Attributes.Endurance).To(BeEquivalentTo(45))
		Expect(class.Attributes.Personality).To(BeEquivalentTo(50))
		Expect(class.Attributes.Speed).To(BeEquivalentTo(43))
		Expect(class.Attributes.Luck).To(BeEquivalentTo(50))
		Expect(class.Skills.Primary).To(ConsistOf("Mysticism", "Alteration", "Thaumaturgy"))
		Expect(class.Skills.Major).To(ConsistOf("Illusion", "Destruction", "Restoration"))
		Expect(class.Skills.Minor).To(ConsistOf("Medical", "Short Blade", "Blunt Weapon", "Dragonish", "Daedric", "Dodging"))
		Expect(class.SPMultiplier / 4).To(BeEquivalentTo(2))
	})

	Context("spell points", func() {
		It("reads spell point values as multipliers", func() {
			// storing all multipliers n as 4*n
			// so that they can all be precise values and not floats
			classCfg[5] = 0
			Expect(classes.ReadClass(classCfg).SPMultiplier).To(BeEquivalentTo(4 * 3))

			classCfg[5] = 4
			Expect(classes.ReadClass(classCfg).SPMultiplier).To(BeEquivalentTo(4 * 2))

			classCfg[5] = 8
			Expect(classes.ReadClass(classCfg).SPMultiplier).To(BeEquivalentTo(7))

			classCfg[5] = 12
			Expect(classes.ReadClass(classCfg).SPMultiplier).To(BeEquivalentTo(6))

			classCfg[5] = 16
			Expect(classes.ReadClass(classCfg).SPMultiplier).To(BeEquivalentTo(4 * 1))

			classCfg[5] = 20
			Expect(classes.ReadClass(classCfg).SPMultiplier).To(BeEquivalentTo(2))
		})
	})
})
