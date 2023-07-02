package utils

import (
	"strings"
	"unicode"
)

func InitRuneMatrix(x, y int, char rune) [][]rune {
	sb := make([][]rune, y)
	for i := 0; i < y; i++ {
		sb[i] = InitRuneArray(x, char)
	}
	return sb
}

func InitRuneArray(size int, char rune) []rune {
	sl := make([]rune, size)
	for i := range sl {
		sl[i] = char
	}
	return sl
}

func PrintableLen(line string) int {
	c := 0
	for _, r := range line {
		if unicode.IsPrint(r) {
			c++
		}
	}
	return c
}

func CutUnicode(line string, limit int) string {
	s := strings.Builder{}
	c := 0
	for _, r := range line {
		if c > limit {
			break
		}
		s.WriteRune(r)
		c++
	}
	return s.String()
}
