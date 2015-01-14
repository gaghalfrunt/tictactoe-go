package cli_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"tictactoe"
	. "tictactoe/cli"
)

var _ = Describe("CLI Input", func() {
	It("reads the next move", func() {
		reader := strings.NewReader("3")
		input := Input{Reader: reader}

		move := input.NextMove(tictactoe.Board{}, tictactoe.X)

		Expect(move).To(Equal(3))
	})

	It("reads integer input", func() {
		reader := strings.NewReader("42")
		input := Input{Reader: reader}

		value := input.ReadInt()

		Expect(value).To(Equal(42))
	})

	It("defaults to 0 on invalid input", func() {
		reader := strings.NewReader("no integer")
		input := Input{Reader: reader}

		value := input.ReadInt()

		Expect(value).To(Equal(0))
	})
})
