package panel

import (
	"github.com/ecoshub/termium/utils"
)

type Stack struct {
	*Basic
	index int
}

func NewStackPanel(width, height int) *Stack {
	return &Stack{
		Basic: &Basic{
			Config: &Config{Width: width, Height: height},
			buffer: utils.InitRuneMatrix(width, height, ' '),
			lines:  make([]string, height),
		},
	}
}

func (sp *Stack) Push(line string) {
	if sp.index >= sp.Config.Height {
		sp.lines = sp.lines[1:]
		sp.lines = append(sp.lines, line)
	} else {
		sp.lines[sp.index] = line
		sp.index++
	}
	sp.render()
}

func (sp *Stack) Clear() {
	sp.Basic.Clear()
	sp.index = 0
}
