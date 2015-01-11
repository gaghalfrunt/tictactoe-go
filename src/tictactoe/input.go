package tictactoe

type Input interface {
	CanProvideNextMove() bool
	NextMove(board Board) int
}
