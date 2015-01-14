package tictactoe

type GameModes map[string][]Player

func AvailableGameModes(input Input) GameModes {
	modes := make(GameModes)

	modes["Human vs. Human"] = human_vs_human(input)
	modes["Human vs. Computer"] = human_vs_computer(input)
	modes["Computer vs. Human"] = computer_vs_human(input)
	modes["Computer vs. Computer"] = computer_vs_computer(input)

	return modes
}

func human_vs_human(input Input) []Player {
	return []Player{
		NewPlayer(X, input),
		NewPlayer(O, input),
	}
}

func human_vs_computer(input Input) []Player {
	return []Player{
		NewPlayer(X, input),
		NewPlayer(O, new(NegamaxInput)),
	}
}

func computer_vs_human(input Input) []Player {
	return []Player{
		NewPlayer(X, new(NegamaxInput)),
		NewPlayer(O, input),
	}
}

func computer_vs_computer(input Input) []Player {
	return []Player{
		NewPlayer(X, new(NegamaxInput)),
		NewPlayer(O, new(NegamaxInput)),
	}
}
