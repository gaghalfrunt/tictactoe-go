package tictactoe

type Game struct {
	players []Player
	board   Board
	output  Output
}

func NewGame(playerA Player, playerB Player, output Output) Game {
	return Game{
		players: []Player{playerA, playerB},
		board:   Board{},
		output:  output,
	}
}

func (game *Game) Play() {
	for {
		if game.isOngoing() {
			mark := game.players[0].Mark()
			move := game.players[0].NextMove()

			game.board.Place(mark, move)

			game.output.ShowBoard(game.board)

			if game.IsFinished() {
				game.showResult()
			}

			game.switchPlayers()
		} else {
			break
		}
	}
}

func (game *Game) IsFinished() bool {
	return game.board.HasWinner() || game.board.HasDraw()
}

func (game *Game) isOngoing() bool {
	return game.nextPlayerIsReady() && !game.IsFinished()
}

func (game *Game) nextPlayerIsReady() bool {
	return game.players[0].IsReady()
}

func (game *Game) switchPlayers() {
	game.players[0], game.players[1] = game.players[1], game.players[0]
}

func (game *Game) showResult() {
	if game.board.HasDraw() {
		game.output.ShowDrawMessage()
	} else {
		game.output.ShowWinnerMessage(game.board.Winner())
	}
}
