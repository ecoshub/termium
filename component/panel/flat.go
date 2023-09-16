package panel

import (
	"github.com/ecoshub/termium/utils/ansi"
)

type Flat struct {
	*Basic
	text string
}

func NewFlatPanel(conf *Config) *Flat {
	b := NewBasicPanel(conf)
	return &Flat{
		Basic: b,
	}
}

func (f *Flat) Append(text string) {
	f.text += text
	f.render()
}

func (f *Flat) Appendln(text string) {
	f.text += text + "\n"
	f.render()
}

func (f *Flat) Clear() {
	f.text = ""
	f.render()
	f.hasChanged()
}

func (f *Flat) GetBuffer() []string {
	return f.lines
}

func (f *Flat) GetSize() (int, int) {
	return f.Config.Width, f.Config.Height
}

func (f *Flat) render() {
	f.text = ansi.Strip(f.text)

	lines := make([]string, f.height)
	index := 0
	s := ""
	for _, r := range f.text {
		if r == rune('\n') {
			lines[index] = s
			s = ""
			index++
			continue
		}
		s += string(r)
		if len(s) == f.Config.Width {
			s = s[:f.Config.Width]
			lines[index] = s
			s = ""
			index++
		}
		if index >= f.height {
			break
		}
	}
	f.lines = lines
}
