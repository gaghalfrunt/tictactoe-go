package tictactoe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "tictactoe"
)

var _ = Describe("Game Modes", func() {
	It("returns a list of available modes", func() {
		var dummyInput Input
		modes := AvailableGameModes(dummyInput)

		Expect(modes["Human vs. Human"]).NotTo(BeNil())
		Expect(modes["Human vs. Computer"]).NotTo(BeNil())
		Expect(modes["Computer vs. Human"]).NotTo(BeNil())
		Expect(modes["Computer vs. Computer"]).NotTo(BeNil())
	})
})
