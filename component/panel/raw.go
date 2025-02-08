package panel

import (
	"fmt"
	"unicode"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/line"
	"github.com/ecoshub/termium/component/renderable"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
)

var _ renderable.Renderable = &Raw{}

const (
	DefaultDefaultChar rune = ' '
)

type Raw struct {
	*Base
	defaultChar rune
	index       int
}

func NewRawPanel(conf *config.Config) *Raw {
	if conf.Height < 1 {
		panic("panels height can not be less than 1 row")
	}
	if conf.RenderTitle {
		conf.Height--
	}
	if conf.ContentStyle == nil {
		conf.ContentStyle = &style.Style{}
	}
	if conf.TitleStyle == nil {
		conf.TitleStyle = &style.Style{}
	}
	b := NewBasicPanel(conf)
	return &Raw{
		Base:        b,
		defaultChar: DefaultDefaultChar,
	}
}

func (raw *Raw) Write(index int, input rune, optionalStyle ...*style.Style) error {
	sty := raw.Config.ContentStyle
	if len(optionalStyle) > 0 {
		sty = optionalStyle[0]
	}
	row := index / raw.Config.Width
	column := index - row*raw.Config.Width
	if row >= (raw.Config.Height) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, raw.Config.Height)
	}
	lineAsBytes := []rune(raw.lines[row].Line)
	buffer := makeEmptyByteLine(raw.Config.Width, raw.defaultChar)
	copy(buffer[:], lineAsBytes[:])
	buffer[column] = input
	raw.lines[row] = &line.Line{Line: string(buffer), Style: sty}
	raw.cleanLine(row)
	raw.changedEvent()
	return nil
}

func (raw *Raw) Put(input rune, optionalStyle ...*style.Style) error {
	// Carriage Return (CR)
	if input == 13 {
		raw.index += raw.Config.Width - raw.index%(raw.Config.Width)
		return nil
	}
	if !unicode.IsGraphic(input) {
		return nil
	}
	defer func() { raw.index++ }()
	return raw.Write(raw.index, input, optionalStyle...)
}

func (r *Raw) Configuration() *config.Config {
	return r.Config
}

func (raw *Raw) ListenChangeHandler(f func()) {
	if f == nil {
		return
	}
	raw.hasChanged = f
}

func (raw *Raw) Clear() {
	raw.lines = line.NewLines(raw.Config.Height, raw.Config.ContentStyle)
	raw.cleanAllLines()
	raw.changedEvent()
}

func (raw *Raw) ClearLine(index int) {
	raw.lines[index] = line.Empty()
}

func (raw *Raw) Buffer() []*line.Line {
	return raw.lines
}

func (raw *Raw) cleanAllLines() {
	for i := range raw.lines {
		raw.cleanLine(i)
	}
}

func (raw *Raw) cleanLine(index int) {
	line := raw.lines[index]
	line.Line = ansi.Strip(line.Line)
	raw.lines[index] = line
}

func makeEmptyByteLine(size int, char rune) []rune {
	buffer := make([]rune, size)
	for i := range buffer {
		buffer[i] = char
	}
	return buffer
}
