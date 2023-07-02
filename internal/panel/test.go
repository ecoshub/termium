package panel

import (
	"fmt"
	"time"

	"github.com/ecoshub/termium/internal/utils"
)

type StackPanel struct {
	*BasePanel
	index int
}

func NewStackPanel(conf *Config) *StackPanel {
	ep := &StackPanel{
		BasePanel: &BasePanel{
			Config: conf,
			buffer: utils.InitRuneMatrix(conf.Size.X, conf.Size.Y, ' '),
			lines:  make([]string, conf.Size.Y),
		},
	}
	if conf.AutoDummyInput {
		go func() {
			n := 0
			for range time.NewTicker(time.Millisecond * 250).C {
				ep.Append(fmt.Sprintf("%s_%d", "\x1b[1;31mHello\x1b[0m my name is eco", n))
				n++
			}
		}()
	}
	return ep
}

func (ep *StackPanel) Append(line string) {
	if ep.index >= ep.Config.Size.Y {
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
