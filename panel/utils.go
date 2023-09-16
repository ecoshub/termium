package panel

import (
	"github.com/ecoshub/termium/utils"
)

func FixedSizeLine(line string, limit int) []rune {
	// line = ansi.Strip(line)
	pll := utils.PrintableLen(line)
	if pll >= limit {
		line := utils.CutUnicode(line, limit)
		runes := []rune(line)
		return runes
	}
	delta := limit - pll
	runes := []rune(line)
	runes = append(runes, utils.InitRuneArray(delta, '.')...)
	return runes
}
