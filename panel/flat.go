package panel

import (
	"github.com/ecoshub/termium/ansi"
	"github.com/ecoshub/termium/utils"
)

type Flat struct {
	Config *Config

	buffer     string
	runeBuffer [][]rune
}

func NewFlatPanel(conf *Config) *Flat {
	return &Flat{
		Config:     conf,
		runeBuffer: utils.InitRuneMatrix(conf.Width, conf.Height, ' '),
	}
}

func (f *Flat) Append(text string) {
	f.buffer += text
	f.render()
}

func (f *Flat) Appendln(text string) {
	f.buffer += text + "\n"
	f.render()
}

func (f *Flat) Clear() {
	f.buffer = ""
	f.render()
}

func (f *Flat) GetBuffer() [][]rune {
	return f.runeBuffer
}

func (f *Flat) GetSize() (int, int) {
	return f.Config.Width, f.Config.Height
}

func (f *Flat) render() {
	mat := utils.InitRuneMatrix(f.Config.Width, f.Config.Height, ' ')
	buff := make([]rune, 0, f.Config.Width)
	index := 0
	f.buffer = ansi.Strip(f.buffer)
	for _, r := range f.buffer {
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
	f.runeBuffer = mat
}
