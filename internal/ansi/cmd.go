package ansi

import (
	"fmt"
)

const (
	EscapeChar string = "\x1b"
)

func EraseLine() {
	print(EscapeChar + "[2K")
}

func ClearScreen() {
	print(EscapeChar + "[2J")
}

func SaveCursorPos() {
	print(EscapeChar + "7")
}

func RestoreCursorPos() {
	print(EscapeChar + "8")
}

func ResetAllModes() {
	print(EscapeChar + "[0m")
}

func GoLeft(n int) {
	fmt.Printf(EscapeChar+"[%dD", n)
}

func GoRight(n int) {
	fmt.Printf(EscapeChar+"[%dC", n)
}

func GoUp(n int) {
	fmt.Printf(EscapeChar+"[%dA", n)
}

func GoDown(n int) {
	fmt.Printf(EscapeChar+"[%dB", n)
}

func GoToFirstBlock() {
	Goto(0, 0)
}

func Goto(l, c int) {
	fmt.Printf(EscapeChar+"[%d;%dH", l, c)
}
