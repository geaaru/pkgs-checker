package pkglist_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPkglist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkglist Suite")
}
