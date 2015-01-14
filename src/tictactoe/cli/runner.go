package cli

import (
	"sort"

	"tictactoe"
)

type Runner struct {
	Input  Input
	Output Output
}

func (runner *Runner) Play() {
	game := runner.setup()
	game.Play()
}

func (runner *Runner) setup() tictactoe.Game {
	modes := tictactoe.AvailableGameModes(&runner.Input)
	modeIdentifiers := runner.orderedIdentifiers(modes)

	mode := runner.requestGameModeIndex(modeIdentifiers)
	players := modes[modeIdentifiers[mode]]

	return tictactoe.NewGame(players[0], players[1], &runner.Output)
}

func (runner *Runner) orderedIdentifiers(modes tictactoe.GameModes) []string {
	var orderedKeys []string

	for key := range modes {
		orderedKeys = append(orderedKeys, key)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(orderedKeys)))

	return orderedKeys
}

func (runner *Runner) requestGameModeIndex(modes []string) int {
	runner.Output.ShowGameModes(modes)

	for {
		runner.Output.ShowGameModePrompt()

		chosenIndex := runner.Input.ReadInt() - 1

		if runner.valueIsValid(chosenIndex, modes) {
			return chosenIndex
		}

		runner.Output.ShowInvalidGameModeMessage()
	}
}

func (runner *Runner) valueIsValid(value int, modes []string) bool {
	return value >= 0 && value < len(modes)
}
