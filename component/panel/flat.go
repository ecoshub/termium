package panel

import (
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
)

type Flat struct {
	*Basic
	text string
}

func NewFlatPanel(conf *Config) *Flat {
	return &Flat{
		Basic: &Basic{
			Config: conf,
			buffer: utils.InitRuneMatrix(conf.Width, conf.Height, ' '),
		},
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

func (f *Flat) GetBuffer() [][]rune {
	return f.buffer
}

func (f *Flat) GetSize() (int, int) {
	return f.Config.Width, f.Config.Height
}

func (f *Flat) render() {
	mat := utils.InitRuneMatrix(f.Config.Width, f.Config.Height, ' ')
	buff := make([]rune, 0, f.Config.Width)
	index := 0
	f.text = ansi.Strip(f.text)
	for _, r := range f.text {
		if r == rune('\n') {
			copy(mat[index], buff[:])
			buff = make([]rune, 0, f.Config.Width)
			index++
			continue
		}
		buff = append(buff, r)
		if utils.PrintableLen(string(buff)) == f.Config.Width {
			copy(mat[index][:f.Config.Width], buff[:f.Config.Width])
			buff = make([]rune, 0, f.Config.Width)
			index++
		}
		if index >= f.Config.Height {
			break
		}
	}
	f.buffer = mat
}
