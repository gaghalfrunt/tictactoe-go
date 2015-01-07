package tictactoe

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Context("Display", func() {
		It("shows the board after a round", func() {
			playerA := NewStubPlayer(X).WithMoves(1)
			playerB := NewStubPlayer(O)
			output := new(SpyOutput)

			game := NewGame(playerA, playerB, output)

			game.PlayRound()

			Expect(output.ShowBoardHasBeenCalled).To(BeTrue())
		})
	})

	Context("Play Policy", func() {
		It("switches players after each round", func() {
			playerA := NewStubPlayer(X).WithMoves(1)
			playerB := NewStubPlayer(O).WithMoves(9)
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

type StubPlayer struct {
	*Player
	moves []int
}

var _ Participant = new(StubPlayer)

func NewStubPlayer(mark Mark) *StubPlayer {
	return &StubPlayer{
		Player: NewPlayer(mark, nil),
	}
}

func (player StubPlayer) NextMove() int {
	move := player.moves[0]
	player.moves = append(player.moves[:0], player.moves[1:]...)

	return move
}

func (player *StubPlayer) WithMoves(moves ...int) *StubPlayer {
	for _, move := range moves {
		player.moves = append(player.moves, move)
	}

	return player
}
