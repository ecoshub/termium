package panel

import (
	"errors"
	"os"

	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
)

type ASCII struct {
	*Base
	content string
	text    string
}

func NewASCIIPanel(conf *Config) *ASCII {
	b := NewBasicPanel(conf)
	return &ASCII{
		Base: b,
	}
}

func (a *ASCII) Append(text string) {
	a.text += text
	a.content += text
	a.render()
}

func (a *ASCII) Appendln(text string) {
	a.Append(text + "\n")
}

func (a *ASCII) Clear() {
	a.text = ""
	a.render()
}

func (a *ASCII) GetBuffer() []*Line {
	return a.lines
}

func (a *ASCII) GetSize() (int, int) {
	return a.width, a.height
}

func (a *ASCII) Flush() {
	a.content = ""
}

func (a *ASCII) Dump(path string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	if a.content == "" {
		return 0, errors.New("panel content is empty")
	}
	n, err := f.WriteString(a.content)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (a *ASCII) render() {
	defer a.hasChanged()

	a.text = ansi.Strip(a.text)

	lines := NewLines(a.height)
	index := 0
	s := ""
	for _, r := range a.text {
		if r == rune('\n') {
			lines[index] = &Line{Line: s, Style: &style.Style{}}
			s = ""
			index++
			continue
		}
		s += string(r)
		if len(s) == a.Config.Width {
			s = s[:a.Config.Width]
			lines[index].Line = s
			s = ""
			index++
		}
		if index >= a.height {
			break
		}
	}
	if s != "" {
		lines[index].Line += s
	}
	a.lines = lines
}
