package printergomocks_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCrudGoMock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CrudGoMock Suite")
}
