package cli_test

import (
	"bufio"
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "tictactoe/cli"
)

var _ = Describe("CLI Runner", func() {
	It("sets up and runs a new game", func() {
		var buffer bytes.Buffer
		output := Output{Writer: bufio.NewWriter(&buffer)}
		reader := strings.NewReader("4")
		input := Input{Reader: reader}

		runner := Runner{
			Input:  input,
			Output: output,
		}

		runner.Play()

		Expect(buffer.String()).To(ContainSubstring("Game ended in a draw"))
	})
})
