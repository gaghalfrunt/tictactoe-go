package tictactoe_test

import . "tictactoe"

var _ Input = new(InputSpy)

type InputSpy struct {
	CanProvideNextMoveHasBeenCalled bool
	NextMoveHasBeenCalled           bool
}

func (input *InputSpy) CanProvideNextMove() bool {
	input.CanProvideNextMoveHasBeenCalled = true

	return true
}

func (input *InputSpy) NextMove(board Board) int {
	input.NextMoveHasBeenCalled = true

	return 1
}
