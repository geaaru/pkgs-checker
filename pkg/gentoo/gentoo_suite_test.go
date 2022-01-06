package gentoo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGentoo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gentoo Suite")
}
