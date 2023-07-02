package screen

import "golang.org/x/crypto/ssh/terminal"

var (
	TerminalWith   int = 0
	TerminalHeight int = 0
)

func init() {
	var err error
	TerminalWith, TerminalHeight, err = terminal.GetSize(0)
	if err != nil {
		panic(err)
	}
}
