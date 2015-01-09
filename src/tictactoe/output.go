package tictactoe

type Output interface {
	ShowBoard(board Board)
	ShowDrawMessage()
	ShowWinnerMessage(winner Mark)
}
