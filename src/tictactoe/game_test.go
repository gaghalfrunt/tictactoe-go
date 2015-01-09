package tictactoe

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Context("Display", func() {
		It("shows the board after a round", func() {
			playerA := NewStubPlayer(X, 1)
			playerB := NewStubPlayer(O)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.Play()

			Expect(output.ShowBoardHasBeenCalled).To(BeTrue())
		})

		It("announces a draw", func() {
			playerA := NewStubPlayer(X, 3, 4, 5, 8, 9)
			playerB := NewStubPlayer(O, 1, 2, 6, 7)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.Play()

			Expect(output.ShowDrawMessageHasBeenCalled).To(BeTrue())
		})

		It("announces a winner", func() {
			playerA := NewStubPlayer(X, 1, 2, 3)
			playerB := NewStubPlayer(O, 4, 5)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.Play()

			Expect(output.ShowWinnerMessageHasBeenCalled).To(BeTrue())
			Expect(output.AnnouncedWinner).To(Equal(X))
		})
	})

	Context("Play Policy", func() {
		It("switches players after each round", func() {
			playerA := NewStubPlayer(X, 1)
			playerB := NewStubPlayer(O, 9)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.Play()

			content := output.board.Content()

			Expect(content[1]).To(Equal(X))
			Expect(content[9]).To(Equal(O))
		})

		It("plays until finished", func() {
			playerA := NewStubPlayer(X, 3, 4, 5, 8, 9)
			playerB := NewStubPlayer(O, 1, 2, 6, 7)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.Play()

			Expect(game.IsFinished()).To(BeTrue())
		})
	})
})
