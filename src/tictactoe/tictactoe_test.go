package tictactoe

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTttGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tic Tac Toe Suite")
}
