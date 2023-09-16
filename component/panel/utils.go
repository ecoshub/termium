package panel

import (
	"fmt"

	"github.com/ecoshub/termium/utils/ansi"
)

func SetStyle(line string, sty *Style) string {
	l := ""
	if sty.SetBold {
		l += ansi.SetBold
	}
	if sty.SetItalic {
		l += ansi.SetItalic
	}
	if sty.SetUnderline {
		l += ansi.SetUnderline
	}
	if sty.SetBlink {
		l += ansi.SetBlink
	}
	if sty.ForegroundColor != 0 {
		l += fmt.Sprintf(ansi.SetForegroundColor, sty.ForegroundColor)
	}
	if sty.BackgroundColor != 0 {
		l += fmt.Sprintf(ansi.SetBackgroundColor, sty.BackgroundColor)
	}
	l += line
	l += ansi.ResetAllModes
	return l
}
