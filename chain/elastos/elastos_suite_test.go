package elastos_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestElastos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Elastos Suite")
}
