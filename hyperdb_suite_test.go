package hyperdb_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHyperdb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hyperdb Suite")
}
