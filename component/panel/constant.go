package panel

import (
	"strings"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
)

func NewTextLine(line string, sty *style.Style) *Base {
	line = ansi.Strip(line)
	line = strings.TrimSpace(line)
	conf := &config.Config{
		Width:        len(line),
		Height:       1,
		RenderTitle:  false,
		ContentStyle: sty,
	}
	bp := NewBasicPanel(conf)
	bp.Write(0, line)
	return bp
}
