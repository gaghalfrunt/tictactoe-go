package cli_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"tictactoe"
	. "tictactoe/cli"
)

var _ = Describe("CLI Input", func() {
	It("uses Reader to get next move", func() {
		reader := strings.NewReader("4")
		input := Input{Reader: reader}

		move := input.NextMove(tictactoe.Board{}, tictactoe.X)

		Expect(move).To(Equal(4))
	})
})
