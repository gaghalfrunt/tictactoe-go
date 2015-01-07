package tictactoe

type Input interface {
	CanProvideNextMove() bool
	NextMove() int
}

type Player struct {
	mark  string
	input Input
}

func NewPlayer(mark string, input Input) *Player {
	return &Player{mark: mark, input: input}
}

func (player *Player) Mark() string {
	return player.mark
}

func (player *Player) IsReady() bool {
	return player.input.CanProvideNextMove()
}

func (player *Player) NextMove() int {
	return player.input.NextMove()
}
