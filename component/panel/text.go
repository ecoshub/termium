package panel

import (
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
)

type Text struct {
	*Base
	text string
}

func NewTextPanel(conf *Config) *Text {
	b := NewBasicPanel(conf)
	return &Text{
		Base: b,
	}
}

func (t *Text) Append(text string) {
	t.text += text
	t.render()
}

func (t *Text) Appendln(text string) {
	t.Append(text + "\n")
}

func (t *Text) Clear() {
	t.text = ""
	t.render()
}

func (t *Text) GetBuffer() []*Line {
	return t.lines
}

func (t *Text) GetSize() (int, int) {
	return t.width, t.height
}

func (t *Text) render() {
	defer t.hasChanged()

	t.text = ansi.Strip(t.text)

	lines := NewLines(t.height)
	index := 0
	s := ""
	for _, r := range t.text {
		if r == rune('\n') {
			lines[index] = &Line{Line: s, Style: &style.Style{}}
			s = ""
			index++
			continue
		}
		s += string(r)
		if len(s) == t.Config.Width {
			s = s[:t.Config.Width]
			lines[index].Line = s
			s = ""
			index++
		}
		if index >= t.height {
			break
		}
	}
	if s != "" {
		lines[index].Line += s
	}
	t.lines = lines
}
