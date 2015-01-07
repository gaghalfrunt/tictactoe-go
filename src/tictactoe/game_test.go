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

type SpyOutput struct {
	board                  Board
	ShowBoardHasBeenCalled bool
}

var _ Output = new(SpyOutput)

func (output *SpyOutput) ShowBoard(board Board) {
	output.ShowBoardHasBeenCalled = true
	output.board = board
}

func NewStubPlayer(mark Mark, moves ...int) Player {
	return *NewPlayer(mark, StubInput{moves})
}

type StubInput struct {
	moves []int
}

var _ Input = new(StubInput)

func (input StubInput) CanProvideNextMove() bool {
	return len(input.moves) > 0
}

func (input StubInput) NextMove() int {
	move := input.moves[0]
	input.moves = append(input.moves[:0], input.moves[1:]...)

	return move
}
