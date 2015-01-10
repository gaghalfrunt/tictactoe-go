package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tictactoe"
)

type Input struct {
}

var _ tictactoe.Input = new(Input)

func (input *Input) CanProvideNextMove() bool {
	return true
}

func (input *Input) NextMove() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Next move: ")
	text, _ := reader.ReadString('\n')
	move, _ := strconv.ParseInt(strings.TrimSpace(text), 10, 0)
	theMove := int(move)

	return theMove
}
