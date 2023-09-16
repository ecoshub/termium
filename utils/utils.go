package utils

import (
	"log"
	"os"
	"os/signal"
	"strings"
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

func FixedSizeLine(line string, limit int) []rune {
	pll := len(line)
	if pll >= limit {
		line := CutUnicode(line, limit)
		runes := []rune(line)
		return runes
	}
	delta := limit - pll
	runes := []rune(line)
	runes = append(runes, InitRuneArray(delta, ' ')...)
	return runes
}

func WaitInterrupt(interruptFunc func()) {
	chanInterrupt := make(chan os.Signal, 1)
	signal.Notify(chanInterrupt, os.Interrupt)
	for {
		select {
		case <-chanInterrupt:
			log.Print("Interrupted. Exiting...")
			if interruptFunc != nil {
				interruptFunc()
			}
			os.Exit(1)
		}
	}
}
