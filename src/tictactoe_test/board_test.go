package tictactoe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "tictactoe"
)

var _ = Describe("Board", func() {
	Context("new", func() {
		It("has all locations available", func() {
			board := Board{}

			Expect(board.AvailableLocations()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})

		It("is not finished", func() {
			board := Board{}

			Expect(board.IsFinished()).To(BeFalse())
		})
	})

	Context("with moves", func() {
		It("occupies the location", func() {
			board := Board{}
			board.Place(X, 1)

			Expect(board.AvailableLocations()).To(Equal([]int{2, 3, 4, 5, 6, 7, 8, 9}))
		})

		It("returns the state of the board", func() {
			board := Board{}
			board.Place(X, 1)
			board.Place(O, 5)
			board.Place(X, 9)

			expectedContent := map[int]Mark{
				1: X, 2: "", 3: "",
				4: "", 5: O, 6: "",
				7: "", 8: "", 9: X,
			}

			Expect(board.Content()).To(Equal(expectedContent))
		})
	})

	Context("finished", func() {
		Context("winning", func() {
			It("row #1", func() {
				board := Board{}
				board.Place(X, 1)
				board.Place(X, 2)
				board.Place(X, 3)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("row #2", func() {
				board := Board{}
				board.Place(X, 4)
				board.Place(X, 5)
				board.Place(X, 6)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("row #3", func() {
				board := Board{}
				board.Place(X, 7)
				board.Place(X, 8)
				board.Place(X, 9)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("column #1", func() {
				board := Board{}
				board.Place(X, 1)
				board.Place(X, 4)
				board.Place(X, 7)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("column #2", func() {
				board := Board{}
				board.Place(X, 2)
				board.Place(X, 5)
				board.Place(X, 8)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("column #3", func() {
				board := Board{}
				board.Place(X, 3)
				board.Place(X, 6)
				board.Place(X, 9)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("diagonal #1", func() {
				board := Board{}
				board.Place(X, 1)
				board.Place(X, 5)
				board.Place(X, 9)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("diagonal #2", func() {
				board := Board{}
				board.Place(X, 3)
				board.Place(X, 5)
				board.Place(X, 7)

				Expect(board.IsFinished()).To(BeTrue())
			})

			It("knows the winner", func() {
				board := Board{}
				board.Place(X, 1)
				board.Place(X, 2)
				board.Place(X, 3)

				Expect(board.Winner()).To(Equal(X))
			})
		})

		Context("draw", func() {
			It("is finished", func() {
				board := Board{}

				board.Place(O, 1)
				board.Place(O, 2)
				board.Place(X, 3)

				board.Place(X, 4)
				board.Place(X, 5)
				board.Place(O, 6)

				board.Place(O, 7)
				board.Place(X, 8)
				board.Place(X, 9)

				Expect(board.IsFinished()).To(BeTrue())
			})
		})
	})
})
