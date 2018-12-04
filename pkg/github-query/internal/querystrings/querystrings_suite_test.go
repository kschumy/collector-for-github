package querystrings_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQuerystrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Querystrings Suite")
}
