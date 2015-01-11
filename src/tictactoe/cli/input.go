package cli

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"tictactoe"
)

type Input struct {
	Reader io.Reader
}

var _ tictactoe.Input = new(Input)

func (input *Input) CanProvideNextMove() bool {
	return true
}

func (input *Input) NextMove(board tictactoe.Board, mark tictactoe.Mark) int {
	reader := bufio.NewReader(input.Reader)

	userInput, _ := reader.ReadString('\n')
	move, _ := strconv.Atoi(strings.TrimSpace(userInput))

	return move
}
