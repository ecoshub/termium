package panel

import (
	"errors"
	"os"
	"strings"

	"github.com/ecoshub/termium/component/style"
)

type Stack struct {
	*Base
	content []string
	index   int
}

func NewStackPanel(conf *Config) *Stack {
	b := NewBasicPanel(conf)
	return &Stack{
		Base: b,
	}
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
	sp.content = append(sp.content, line)
	sp.render()
}

func (sp *Stack) Flush() {
	sp.content = make([]string, 0, 16)
}

func (sp *Stack) Dump(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if len(sp.content) == 0 {
		return errors.New("panel content is empty")
	}
	b := strings.Builder{}
	b.Grow(20 * len(sp.content))
	for _, s := range sp.content {
		b.WriteString(s + "\n")
	}
	_, err = f.WriteString(b.String())
	if err != nil {
		return err
	}
	return nil
}

func (sp *Stack) Clear() {
	sp.Base.Clear()
	sp.index = 0
}
