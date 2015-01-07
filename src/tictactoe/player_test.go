package tictactoe

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	var (
		inputSpy *InputSpy
		player   *Player
	)

	BeforeEach(func() {
		inputSpy = new(InputSpy)
		player = NewPlayer(X, inputSpy)
	})

	It("is a Participant", func() {
		var player Participant = NewPlayer(X, inputSpy)

		Expect(player).NotTo(BeNil())
	})

	It("has a mark", func() {
		player := NewPlayer(X, nil)

		Expect(player.Mark()).To(Equal(X))
	})

	It("has a way to provide a new move", func() {
		Expect(inputSpy.CanProvideNextMove()).To(Equal(true))
	})

	It("is ready when the input can provide a new move", func() {
		player.IsReady()

		Expect(inputSpy.CanProvideNextMoveHasBeenCalled).To(Equal(true))
	})

	It("provides the move from its input", func() {
		inputSpy.Move = 42

		theMove := player.NextMove()

		Expect(inputSpy.NextMoveHasBeenCalled).To(Equal(true))
		Expect(theMove).To(Equal(42))
	})
})

type InputSpy struct {
	Move                            int
	CanProvideNextMoveHasBeenCalled bool
	NextMoveHasBeenCalled           bool
}

var _ Input = new(InputSpy)

func (input *InputSpy) CanProvideNextMove() bool {
	input.CanProvideNextMoveHasBeenCalled = true

	return true
}

func (input *InputSpy) NextMove() int {
	input.NextMoveHasBeenCalled = true

	return input.Move
}
