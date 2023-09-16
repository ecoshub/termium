package ansi

import (
	"fmt"
)

const (
	EscapeChar          string = "\x1b"
	EraseLine           string = EscapeChar + "[2K"
	ClearScreen         string = EscapeChar + "[2J"
	MakeCursorInvisible string = EscapeChar + "[?25l"
	MakeCursorVisible   string = EscapeChar + "[?25h"
	SaveCursorPos       string = EscapeChar + "7"
	RestoreCursorPos    string = EscapeChar + "8"
	GoLeftOneChar       string = EscapeChar + "[1D"
	GoRightOneChar      string = EscapeChar + "[1C"
	GoDownOneChar       string = EscapeChar + "[1B"
	GoUpOneChar         string = EscapeChar + "[1A"
	GoToFirstBlock      string = EscapeChar + "[1;1H"
	ResetAllModes       string = EscapeChar + "[0m"
	SetForegroundColor  string = EscapeChar + "[38;5;%dm"
	SetBackgroundColor  string = EscapeChar + "[48;5;%dm"
	SetBold             string = EscapeChar + "[1m"
	ResetBold           string = EscapeChar + "[22m"
	SetItalic           string = EscapeChar + "[3m"
	ResetItalic         string = EscapeChar + "[23m"
	SetUnderline        string = EscapeChar + "[4m"
	ResetUnderline      string = EscapeChar + "[24m"
	SetBlink            string = EscapeChar + "[5m"
	ResetBlink          string = EscapeChar + "[25m"
)

func GoLeftNChar(n int) {
	fmt.Printf(EscapeChar+"[%dD", n)
}

func GoRightNChar(n int) {
	fmt.Printf(EscapeChar+"[%dC", n)
}

func GoUpNChar(n int) {
	fmt.Printf(EscapeChar+"[%dA", n)
}

func GoDownNChar(n int) {
	fmt.Printf(EscapeChar+"[%dB", n)
}

func GotoRowAndColumn(l, c int) {
	fmt.Printf(EscapeChar+"[%d;%dH", l, c)
}
