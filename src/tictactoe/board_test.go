package tictactoe

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	Context("new", func() {
		It("has all locations available", func() {
			board := Board{}

			Expect(board.availableLocations()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		})
	})

	Context("with moves", func() {
		It("occupies the location", func() {
			board := Board{}
			board.place("x", 1)

			Expect(board.availableLocations()).To(Equal([]int{2, 3, 4, 5, 6, 7, 8, 9}))
		})

		It("returns the state of the board", func() {
			board := Board{}
			board.place("x", 1)
			board.place("o", 5)
			board.place("x", 9)

			expectedContent := map[int]Mark{
				1: "x", 2: "", 3: "",
				4: "", 5: "o", 6: "",
				7: "", 8: "", 9: "x",
			}

			Expect(board.content()).To(Equal(expectedContent))
		})
	})
})
