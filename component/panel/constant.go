package panel

import (
	"strings"

	"github.com/ecoshub/termium/utils/ansi"
)

func ConstantText(line string, foregroundColor, backgroundColor int) *Basic {
	line = ansi.Strip(line)
	line = strings.TrimSpace(line)
	conf := &Config{
		Width:           len(line),
		Height:          1,
		RenderTitle:     false,
		ForegroundColor: foregroundColor,
		BackgroundColor: backgroundColor,
	}
	bp := NewBasicPanel(conf)
	bp.Write(0, line)
	bp.hasChanged()
	return bp
}
