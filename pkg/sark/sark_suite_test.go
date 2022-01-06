package sark_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSark(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SARK Suite")
}
