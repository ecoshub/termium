package panel

import (
	"github.com/ecoshub/termium/utils"
)

type Stack struct {
	*Basic
	index int
}

func NewStack(width, height int) *Stack {
	ep := &Stack{
		Basic: &Basic{
			Config: &Config{Width: width, Height: height},
			buffer: utils.InitRuneMatrix(width, height, ' '),
			lines:  make([]string, height),
		},
	}
	return ep
}

func (ep *Stack) Push(line string) {
	if ep.index >= ep.Config.Height {
		ep.lines = ep.lines[1:]
		ep.lines = append(ep.lines, line)
	} else {
		ep.lines[ep.index] = line
		ep.index++
	}
	ep.renderList()
}

func (ep *Stack) Clear() {
	ep.Basic.Clear()
	ep.index = 0
}
