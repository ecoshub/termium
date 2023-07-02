package ansi

import (
	"fmt"
)

func EraseLine() {
	print("\x1b[2K")
}

func GoLeft(n int) {
	fmt.Printf("\x1b[%dD", n)
}

func GoRight(n int) {
	fmt.Printf("\x1b[%dC", n)
}

func GoUp(n int) {
	fmt.Printf("\x1b[%dA", n)
}

func GoDown(n int) {
	fmt.Printf("\x1b[%dB", n)
}

func GoToFirstBlock() {
	Goto(0, 0)
}

func ClearScreen() {
	print("\x1b[2J")
}

func Goto(l, c int) {
	fmt.Printf("\x1b[%d;%dH", l, c)
}

func SaveCursorPos() {
	print("\x1b7")
}

func RestoreCursorPos() {
	print("\x1b8")
}

func ResetAllModes() {
	print("\x1b[0m")
}
