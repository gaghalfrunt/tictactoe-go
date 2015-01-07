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

			game.PlayRound()

			Expect(output.ShowBoardHasBeenCalled).To(BeTrue())
		})
	})

	Context("Play Policy", func() {
		It("switches players after each round", func() {
			playerA := NewStubPlayer(X, 1)
			playerB := NewStubPlayer(O, 9)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.PlayRound()
			game.PlayRound()

			Expect(output.board.Content()[1]).To(Equal(X))
			Expect(output.board.Content()[9]).To(Equal(O))
		})
	})
})
