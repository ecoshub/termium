package panel

import (
	"errors"
	"os"
	"strings"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/line"
	"github.com/ecoshub/termium/component/renderable"
	"github.com/ecoshub/termium/component/style"
)

var _ renderable.Renderable = &Stack{}

type Stack struct {
	*Base
	content []string
	index   int
}

func NewStackPanel(conf *config.Config) *Stack {
	b := NewBasicPanel(conf)
	return &Stack{
		Base: b,
	}
}

func (sp *Stack) Push(input string, optionalStyle ...*style.Style) {
	sty := sp.Config.ContentStyle
	if len(optionalStyle) > 0 {
		sty = optionalStyle[0]
	}
	if sp.index < sp.Config.Height {
		sp.lines[sp.index] = &line.Line{Line: input, Style: sty}
		sp.index++
	} else {
		// shift all up and add the last one on the end
		for i := 0; i < sp.Config.Height-1; i++ {
			sp.lines[i] = sp.lines[i+1]
		}
		sp.lines[sp.Config.Height-1] = &line.Line{Line: input, Style: sty}
	}
	sp.content = append(sp.content, input)
	sp.clearAllLines()
	sp.changedEvent()
}

func (sp *Stack) Flush() {
	sp.Clear()
	sp.content = make([]string, 0, 16)
	sp.index = 0
}

func (sp *Stack) Dump(path string) (int, error) {
	if len(sp.content) == 0 {
		return 0, errors.New("panel content is empty")
	}
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	b := strings.Builder{}
	b.Grow(20 * len(sp.content))
	for _, s := range sp.content {
		b.WriteString(s + "\n")
	}
	n, err := f.WriteString(b.String())
	if err != nil {
		return 0, err
	}
	return n, nil
}
