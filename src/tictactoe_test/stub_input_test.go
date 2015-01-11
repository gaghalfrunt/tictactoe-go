package tictactoe_test

import . "tictactoe"

var _ Input = new(StubInput)

type StubInput struct {
	moves []int
}

func (input *StubInput) CanProvideNextMove() bool {
	return len(input.moves) > 0
}

func (input *StubInput) NextMove(board Board, mark Mark) int {
	move := input.moves[0]
	input.moves = append(input.moves[:0], input.moves[1:]...)

	return move
}
