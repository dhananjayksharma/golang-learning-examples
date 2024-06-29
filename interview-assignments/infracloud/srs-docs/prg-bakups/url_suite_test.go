package urls_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestShortURL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShortURL Suite")
}
