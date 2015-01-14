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

	Context("Game Output", func() {
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

		It("shows a request message for the next move", func() {
			output.ShowNextMoveMessage(tictactoe.O)

			Expect(buffer.String()).To(ContainSubstring("Next move for o"))
		})

		It("shows an error message for invalid moves", func() {
			output.ShowInvalidMoveMessage()

			Expect(buffer.String()).To(ContainSubstring("Invalid move"))
		})
	})

	Context("Game mode selection", func() {
		It("prints a list of possible options", func() {
			modes := []string{
				"Human vs. Human",
				"Human vs. Computer",
				"Computer vs. Human",
				"Computer vs. Computer",
			}

			output.ShowGameModes(modes)

			Expect(buffer.String()).To(ContainSubstring("Available Game Modes"))
			Expect(buffer.String()).To(ContainSubstring("1. Human vs. Human"))
			Expect(buffer.String()).To(ContainSubstring("2. Human vs. Computer"))
			Expect(buffer.String()).To(ContainSubstring("3. Computer vs. Human"))
			Expect(buffer.String()).To(ContainSubstring("4. Computer vs. Computer"))
		})

		It("asks for a mode", func() {
			output.ShowGameModePrompt()

			Expect(buffer.String()).To(ContainSubstring("Game Mode to play: "))
		})

		It("shows an error message for invalid Game Mode selection", func() {
			output.ShowInvalidGameModeMessage()

			Expect(buffer.String()).To(ContainSubstring("Invalid game mode"))
		})
	})
})
