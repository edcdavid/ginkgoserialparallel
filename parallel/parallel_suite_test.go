package learning1_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParallel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parallel Suite")
}
