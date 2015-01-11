package tictactoe

type Input interface {
	CanProvideNextMove() bool
	NextMove() int
}
