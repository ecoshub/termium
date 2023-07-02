package component

import (
	"term/internal/ansi"
	"term/internal/utils"
)

func ResizeLine(line string, limit int) []rune {
	line = ansi.Strip(line)
	if utils.PrintableLen(line) >= limit {
		line := utils.CutUnicode(line, limit)
		runes := []rune(line)
		return runes
	}
	delta := limit - utils.PrintableLen(line)
	runes := []rune(line)
	runes = append(runes, utils.InitRuneArray(delta, ' ')...)
	return runes
}
