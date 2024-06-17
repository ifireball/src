package clone_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clone Suite")
}
