package tictactoe

type NegamaxInput struct{}

func (negamax *NegamaxInput) CanProvideNextMove() bool {
	return true
}

func (negamax *NegamaxInput) NextMove(board Board, mark Mark) int {
	bestMove := 0
	bestScore := -1.0

	for _, location := range board.AvailableLocations() {
		var newBoard Board = board
		newBoard.Place(mark, location)

		score := -negamax.scoreBoard(newBoard, Opponent(mark), bestScore, 1.0)

		if score > bestScore {
			bestScore = score
			bestMove = location
		}
	}

	return bestMove
}

func (negamax *NegamaxInput) scoreBoard(board Board, mark Mark, alpha float64, beta float64) float64 {
	if board.IsFinished() {
		return negamax.score(board, mark)
	}

	bestScore := alpha

	for _, location := range board.AvailableLocations() {
		var newBoard = board
		newBoard.Place(mark, location)

		score := -negamax.scoreBoard(newBoard, Opponent(mark), -beta, -alpha)

		bestScore = max(bestScore, score)
		alpha = max(alpha, score)

		if alpha >= beta {
			break
		}
	}

	return bestScore
}

func (negamax *NegamaxInput) score(board Board, mark Mark) float64 {
	if board.Winner() == mark {
		return board.Value()
	} else {
		return -board.Value()
	}
}

func max(a float64, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
