package tictactoe

type Player struct {
	mark  Mark
	input Input
}

func NewPlayer(mark Mark, input Input) Player {
	return Player{mark: mark, input: input}
}

func (player *Player) Mark() Mark {
	return player.mark
}

func (player *Player) IsReady() bool {
	return player.input.CanProvideNextMove()
}

func (player *Player) NextMove(board Board) int {
	return player.input.NextMove(board)
}
