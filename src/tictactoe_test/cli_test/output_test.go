package cli_test

import (
	"bufio"
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"tictactoe"
	. "tictactoe/cli"
)

var _ = Describe("CLI Output", func() {
	var (
		buffer bytes.Buffer
		output Output
	)

	BeforeEach(func() {
		buffer.Reset()
		output = Output{Writer: bufio.NewWriter(&buffer)}
	})

	It("announces a draw", func() {
		output.ShowDrawMessage()

		Expect(buffer.String()).To(ContainSubstring("Game ended in a draw"))
	})

	It("announces a winner", func() {
		output.ShowWinnerMessage(tictactoe.X)

		Expect(buffer.String()).To(ContainSubstring("Winner is: x"))
	})

	It("prints the board content", func() {
		board := tictactoe.Board{}

		board.Place(tictactoe.X, 1)
		board.Place(tictactoe.O, 5)
		board.Place(tictactoe.X, 9)

		output.ShowBoard(board)

		Expect(buffer.String()).To(ContainSubstring("x | 2 | 3"))
		Expect(buffer.String()).To(ContainSubstring("4 | o | 6"))
		Expect(buffer.String()).To(ContainSubstring("7 | 8 | x"))
	})

	It("show a request message for the next move", func() {
		output.ShowNextMoveMessage(tictactoe.O)

		Expect(buffer.String()).To(ContainSubstring("Next move for o"))
	})
})
