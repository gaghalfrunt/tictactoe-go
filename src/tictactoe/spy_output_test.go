package tictactoe

var _ Output = new(SpyOutput)

type SpyOutput struct {
	board                  Board
	ShowBoardHasBeenCalled bool
}

func (output *SpyOutput) ShowBoard(board Board) {
	output.ShowBoardHasBeenCalled = true
	output.board = board
}
