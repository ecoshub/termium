package utils

import "golang.org/x/term"

var (
	TerminalWith   int = 0
	TerminalHeight int = 0
)

func init() {
	var err error
	TerminalWith, TerminalHeight, err = term.GetSize(0)
	if err != nil {
		panic(err)
	}
}
