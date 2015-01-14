package tictactoe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "tictactoe"
)

var _ = Describe("Player", func() {
	var (
		spyInput *SpyInput
		player   Player
	)

	BeforeEach(func() {
		spyInput = new(SpyInput)
		player = NewPlayer(X, spyInput)
	})

	It("has a mark", func() {
		player := NewPlayer(X, nil)

		Expect(player.Mark()).To(Equal(X))
	})

	It("has a way to provide a new move", func() {
		Expect(spyInput.CanProvideNextMove()).To(Equal(true))
	})

	It("is ready when the input can provide a new move", func() {
		player.IsReady()

		Expect(spyInput.CanProvideNextMoveHasBeenCalled).To(Equal(true))
	})

	It("provides the move from its input", func() {
		player.NextMove(Board{})

		Expect(spyInput.NextMoveHasBeenCalled).To(Equal(true))
	})
})
