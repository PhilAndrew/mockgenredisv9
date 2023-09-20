package mock_v9

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCaches(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golang Library - Caches - Test Suite")
}
