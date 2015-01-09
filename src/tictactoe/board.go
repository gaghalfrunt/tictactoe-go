package tictactoe

type Mark string

const (
	X Mark = "x"
	O Mark = "o"
)

func Opponent(mark Mark) Mark {
	if mark == X {
		return O
	} else {
		return X
	}
}

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

func (board Board) IsFinished() bool {
	if board.HasWinner() || board.HasDraw() {
		return true
	}

	return false
}

func (board Board) HasWinner() bool {
	for _, line := range board.allLines() {
		if line.hasWinner() {
			return true
		}
	}

	return false
}

func (board Board) HasDraw() bool {
	availableMoves := board.availableLocations()

	return len(availableMoves) == 0 && !board.HasWinner()
}

func (board Board) Winner() Mark {
	for _, line := range board.allLines() {
		if line.hasWinner() {
			return *line.winner()
		}
	}

	return ""
}

func (board Board) allLines() []line {
	var lines []line

	lines = append(lines, board.rows()...)
	lines = append(lines, board.columns()...)
	lines = append(lines, board.diagonals()...)

	return lines
}

func (board Board) rows() []line {
	content := board.Content()

	return []line{
		newLine(content[1], content[2], content[3]),
		newLine(content[4], content[5], content[6]),
		newLine(content[7], content[8], content[9]),
	}
}

func (board Board) columns() []line {
	content := board.Content()

	return []line{
		newLine(content[1], content[4], content[7]),
		newLine(content[2], content[5], content[8]),
		newLine(content[3], content[6], content[9]),
	}
}

func (board Board) diagonals() []line {
	content := board.Content()

	return []line{
		newLine(content[1], content[5], content[9]),
		newLine(content[3], content[5], content[7]),
	}
}

type line struct {
	elements []Mark
}

func newLine(marks ...Mark) line {
	return line{marks}
}

func (line line) hasWinner() bool {
	if line.hasUniqueMark(X) || line.hasUniqueMark(O) {
		return true
	}

	return false
}

func (line line) hasUniqueMark(mark Mark) bool {
	for _, element := range line.elements {
		if element != mark {
			return false
		}
	}

	return true
}

func (line line) winner() *Mark {
	if line.hasWinner() {
		return &line.elements[0]
	}

	return nil
}
