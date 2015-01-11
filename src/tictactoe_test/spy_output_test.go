package tictactoe_test

import . "tictactoe"

var _ Output = new(SpyOutput)

type SpyOutput struct {
	board                            Board
	AnnouncedWinner                  Mark
	ShowBoardHasBeenCalled           bool
	ShowDrawMessageHasBeenCalled     bool
	ShowWinnerMessageHasBeenCalled   bool
	ShowNextMoveMessageHasBeenCalled bool
}

func (output *SpyOutput) ShowBoard(board Board) {
	output.ShowBoardHasBeenCalled = true
	output.board = board
}

func (output *SpyOutput) ShowDrawMessage() {
	output.ShowDrawMessageHasBeenCalled = true
}

func (output *SpyOutput) ShowWinnerMessage(winner Mark) {
	output.ShowWinnerMessageHasBeenCalled = true
	output.AnnouncedWinner = winner
}

func (output *SpyOutput) ShowNextMoveMessage(nextPlayer Mark) {
	output.ShowNextMoveMessageHasBeenCalled = true
}
