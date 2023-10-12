package style

import (
	"fmt"

	"github.com/ecoshub/termium/utils/ansi"
)

var (
	// DefaultStyleEvent a yellowish color style
	DefaultStyleInfo *Style = &Style{
		ForegroundColor: 228,
		Italic:          true,
	}
	// DefaultStyleEvent a orangish color style
	DefaultStyleWarning *Style = &Style{
		ForegroundColor: 209,
	}
	// DefaultStyleEvent a reddish color style
	DefaultStyleError *Style = &Style{
		ForegroundColor: 197,
	}
	// DefaultStyleEvent a greenish color style
	DefaultStyleSuccess *Style = &Style{
		ForegroundColor: 83,
	}
	// DefaultStyleEvent a cyan like color style
	DefaultStyleEvent *Style = &Style{
		ForegroundColor: 45,
	}
)

type Style struct {
	ForegroundColor int
	BackgroundColor int
	Bold            bool
	Italic          bool
	Underline       bool
	Blink           bool
}

func SetStyle(line string, sty *Style) string {
	l := ""
	if sty.Bold {
		l += ansi.SetBold
	}
	if sty.Italic {
		l += ansi.SetItalic
	}
	if sty.Underline {
		l += ansi.SetUnderline
	}
	if sty.Blink {
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
