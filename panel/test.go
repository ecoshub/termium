package panel

import (
	"github.com/ecoshub/termium/utils"
)

type StackPanel struct {
	*BasePanel
	index int
}

func NewStackPanel(conf *Config) *StackPanel {
	ep := &StackPanel{
		BasePanel: &BasePanel{
			Config: conf,
			buffer: utils.InitRuneMatrix(conf.SizeX, conf.SizeY, ' '),
			lines:  make([]string, conf.SizeY),
		},
	}
	return ep
}

func (ep *StackPanel) Append(line string) {
	if ep.index >= ep.Config.SizeY {
		ep.lines = ep.lines[1:]
		ep.lines = append(ep.lines, line)
	} else {
		ep.lines[ep.index] = line
		ep.index++
	}
	ep.renderList()
}

func (ep *StackPanel) Clear() {
	ep.BasePanel.Clear()
	ep.index = 0
}
