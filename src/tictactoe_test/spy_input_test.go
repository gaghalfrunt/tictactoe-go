package tictactoe_test

import . "tictactoe"

var _ Input = new(SpyInput)

type SpyInput struct {
	CanProvideNextMoveHasBeenCalled bool
	NextMoveHasBeenCalled           bool
}

func (input *SpyInput) CanProvideNextMove() bool {
	input.CanProvideNextMoveHasBeenCalled = true

	return true
}

func (input *SpyInput) NextMove(board Board, mark Mark) int {
	input.NextMoveHasBeenCalled = true

	return 1
}
