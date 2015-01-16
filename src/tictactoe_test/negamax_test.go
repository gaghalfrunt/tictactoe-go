package tictactoe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "tictactoe"
)

var _ = Describe("AI", func() {
	It("is an Input object", func() {
		var negamax Input = new(NegamaxInput)

		Expect(negamax).NotTo(BeNil())
	})

	It("can provide the next move", func() {
		negamax := new(NegamaxInput)

		Expect(negamax.CanProvideNextMove()).To(BeTrue())
	})

	It("goes for wins", func() {
		Expect(boardWith("-xx------").next(X)).To(Equal(1))
		Expect(boardWith("--x---x--").next(X)).To(Equal(5))
		Expect(boardWith("-------xx").next(X)).To(Equal(7))
	})

	It("blocks winning states", func() {
		Expect(boardWith("xx-------").next(O)).To(Equal(3))
		Expect(boardWith("x-------x").next(O)).To(Equal(5))
		Expect(boardWith("xox-o-oxx").next(O)).To(Equal(6))
	})

	It("blocks a possible fork", func() {
		Expect(boardWith("x--------").next(O)).To(Equal(5))
		Expect(boardWith("--x------").next(O)).To(Equal(5))
		Expect(boardWith("------x--").next(O)).To(Equal(5))
		Expect(boardWith("--------x").next(O)).To(Equal(5))
	})

	It("opens the game with a corner location", func() {
		Expect([]int{1, 3, 7, 9}).To(ContainElement(boardWith("---------").next(X)))
	})
})

type boardBuilder struct {
	board Board
}

func (builder boardBuilder) next(nextPlayer Mark) int {
	negamax := new(NegamaxInput)
	return negamax.NextMove(builder.board, nextPlayer)
}

func boardWith(content string) boardBuilder {
	board := Board{}
	marks := map[string]Mark{
		"x": X,
		"o": O,
	}

	for index, character := range content {
		mark, ok := marks[string(character)]

		if ok {
			board.Place(mark, index+1)
		}
	}

	return boardBuilder{board: board}
}
