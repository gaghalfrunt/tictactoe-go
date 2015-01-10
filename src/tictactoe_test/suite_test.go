package tictactoe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTicTacToeSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tic Tac Toe Suite")
}
