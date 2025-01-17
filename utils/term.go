package utils

import "golang.org/x/term"

func GetTerminalSize() (int, int, error) {
	return term.GetSize(0)
}
