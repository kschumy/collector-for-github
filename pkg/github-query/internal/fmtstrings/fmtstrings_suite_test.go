package fmtstrings_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFmtstrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fmtstrings Suite")
}
