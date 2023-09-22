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

func (sp *Stack) Append(line string) {
	if sp.index == 0 {
		sp.Push(line)
		return
	}
	l := sp.lines[sp.index-1]
	l.Line = l.Line + line
	sp.render()
}

func (sp *Stack) Appendln(line string) {
	sp.Append(line)
	sp.index++
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
	if sp.index > 1 {
		sp.content = append(sp.content, sp.lines[sp.index-2].Line)
	}
	sp.render()
}

func (sp *Stack) Flush() {
	sp.content = make([]string, 0, 16)
}

func (sp *Stack) Dump(path string) (int, error) {
	if sp.index == 0 {
		return 0, errors.New("panel content is empty")
	}
	dump := append(sp.content, sp.lines[sp.index-1].Line)
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	b := strings.Builder{}
	b.Grow(20 * len(dump))
	for _, s := range dump {
		b.WriteString(s + "\n")
	}
	n, err := f.WriteString(b.String())
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (sp *Stack) Clear() {
	sp.Base.Clear()
	sp.index = 0
}
