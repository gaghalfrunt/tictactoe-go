package tictactoe

type Game struct {
	players []Player
	board   Board
	output  Output
}

type Output interface {
	ShowBoard(board Board)
}

func NewGame(playerA Player, playerB Player, output Output) Game {
	return Game{
		players: []Player{playerA, playerB},
		board:   Board{},
		output:  output,
	}
}

func (game *Game) PlayRound() {
	mark := game.players[0].Mark()
	move := game.players[0].NextMove()

	game.board.Place(mark, move)

	game.output.ShowBoard(game.board)

	game.players[0], game.players[1] = game.players[1], game.players[0]
}
