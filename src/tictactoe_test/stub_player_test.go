package tictactoe_test

import . "tictactoe"

func NewStubPlayer(mark Mark, moves ...int) Player {
	return NewPlayer(mark, &StubInput{moves})
}
