package bsareader_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBsareader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bsareader Suite")
}
