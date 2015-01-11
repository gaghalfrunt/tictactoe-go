package tictactoe

type NegamaxInput struct{}

func (negamax *NegamaxInput) CanProvideNextMove() bool {
	return true
}

func (negamax *NegamaxInput) NextMove(board Board, mark Mark) int {
	bestMove := 0
	bestScore := -10

	for _, location := range board.AvailableLocations() {
		var newBoard Board = board
		newBoard.Place(mark, location)

		score := -negamax.scoreBoard(newBoard, Opponent(mark), bestScore, 10)

		if score > bestScore {
			bestScore = score
			bestMove = location
		}
	}

	return bestMove
}

func (negamax *NegamaxInput) scoreBoard(board Board, mark Mark, alpha int, beta int) int {
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

func (negamax *NegamaxInput) score(board Board, mark Mark) int {
	score := len(board.AvailableLocations())

	if board.Winner() == mark {
		return score
	} else {
		return -score
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
