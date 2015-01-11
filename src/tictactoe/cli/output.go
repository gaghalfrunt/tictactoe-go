package cli

import (
	"bufio"
	"fmt"
	"strconv"

	"tictactoe"
)

type Output struct {
	Writer *bufio.Writer
}

var _ tictactoe.Output = new(Output)

const boardTemplate = "\n" +
	" %s | %s | %s\n" +
	"----------\n" +
	" %s | %s | %s\n" +
	"----------\n" +
	" %s | %s | %s\n\n"

func (output *Output) ShowBoard(board tictactoe.Board) {
	var linearizedBoard []interface{}
	content := board.Content()

	for location := 1; location <= len(board.Content()); location++ {
		mark := content[location]

		if mark == "" {
			linearizedBoard = append(linearizedBoard, strconv.Itoa(location))
		} else {
			linearizedBoard = append(linearizedBoard, string(mark))
		}
	}

	output.write(boardTemplate, linearizedBoard...)
}

func (output *Output) ShowDrawMessage() {
	output.write("Game ended in a draw.")
}

func (output *Output) ShowWinnerMessage(winner tictactoe.Mark) {
	output.write("Game ended. Winner is: %s\n", winner)
}

func (output *Output) ShowNextMoveMessage(nextPlayer tictactoe.Mark) {
	output.write("Next move for %s: ", nextPlayer)
}

func (output *Output) ShowInvalidMoveMessage() {
	output.write("Invalid move\n")
}

func (output *Output) write(content string, args ...interface{}) {
	fmt.Fprintf(output.Writer, content, args...)
	output.Writer.Flush()
}
