package panel

import (
	"errors"

	"github.com/ecoshub/termium/component/style"
)

type Stack struct {
	*Base
	index int
}

func NewStackPanel(conf *Config) *Stack {
	b := NewBasicPanel(conf)
	return &Stack{
		Base: b,
	}
}

func (sp *Stack) Write(index int, line string) error {
	if index > sp.index {
		return errors.New("index out of range")
	}
	sp.lines[index] = &Line{Line: "", Style: &style.Style{}}
	sp.render()
	return nil
}

func (sp *Stack) Push(line string, optionalStyle ...*style.Style) {
	sty := sp.Config.ContentStyle
	if len(optionalStyle) > 0 {
		sty = optionalStyle[0]
	}
	if sp.index >= sp.height {
		sp.lines = sp.lines[1:]
		sp.lines = append(sp.lines, &Line{Line: line, Style: sty})
	} else {
		sp.lines[sp.index] = &Line{Line: line, Style: sty}
		sp.index++
	}
	sp.render()
}

func (sp *Stack) Clear() {
	sp.Base.Clear()
	sp.index = 0
}
