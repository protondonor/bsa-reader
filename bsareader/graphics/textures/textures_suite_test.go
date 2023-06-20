package textures_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Textures Suite")
}
