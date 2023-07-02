package ansi

import "strconv"

const (
	ColorCodeGray  int = 240
	ColorCodeRed   int = 1
	ColorCodeGreen int = 77
)

func ColorLine(colorCode int, line string) string {
	return "\x1b[38;5;" + strconv.Itoa(colorCode) + "m" + line + "\x1b[0m"
}
