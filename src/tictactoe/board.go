package tictactoe

type Mark string

const (
	X Mark = "x"
	O Mark = "o"
)

type Board struct {
	cells [9]Mark
}

func (board *Board) availableLocations() []int {
	var freeLocations []int

	for index, element := range board.cells {
		if element == "" {
			freeLocations = append(freeLocations, index+1)
		}
	}

	return freeLocations
}

func (board *Board) Place(mark Mark, location int) {
	board.cells[location-1] = mark
}

func (board *Board) Content() map[int]Mark {
	content := make(map[int]Mark)

	for index, value := range board.cells {
		content[index+1] = value
	}

	return content
}
