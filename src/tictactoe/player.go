package tictactoe

type Input interface {
	CanProvideNextMove() bool
	NextMove() int
}

type Participant interface {
	Mark() Mark
	IsReady() bool
	NextMove() int
}

type Player struct {
	mark  Mark
	input Input
}

func NewPlayer(mark Mark, input Input) *Player {
	return &Player{mark: mark, input: input}
}

func (player *Player) Mark() Mark {
	return player.mark
}

func (player *Player) IsReady() bool {
	return player.input.CanProvideNextMove()
}

func (player *Player) NextMove() int {
	return player.input.NextMove()
}
