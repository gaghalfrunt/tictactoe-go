package main

import (
	"bufio"
	"os"

	"tictactoe/cli"
)

func main() {
	runner := cli.Runner{
		Input:  cli.Input{Reader: os.Stdin},
		Output: cli.Output{Writer: bufio.NewWriter(os.Stdout)},
	}

	runner.Play()
}
