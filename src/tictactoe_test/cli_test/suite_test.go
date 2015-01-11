package cli_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTicTacToeCLISuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tic Tac Toe CLI Suite")
}
