package tictactoe

type Input interface {
	CanProvideNextMove() bool
	NextMove(board Board, mark Mark) int
}
