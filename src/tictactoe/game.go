package tictactoe

type Game struct {
	participants []Participant
	board        Board
	output       Output
}

type Output interface {
	ShowBoard(board Board)
}

func NewGame(playerA Participant, playerB Participant, output Output) Game {
	return Game{
		participants: []Participant{playerA, playerB},
		board:        Board{},
		output:       output,
	}
}

func (game *Game) PlayRound() {
	mark := game.participants[0].Mark()
	move := game.participants[0].NextMove()

	game.board.Place(mark, move)

	game.output.ShowBoard(game.board)

	game.participants[0], game.participants[1] = game.participants[1], game.participants[0]
}
