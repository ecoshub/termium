package panel

import (
	"fmt"

	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
)

type Config struct {
	Width        int
	Height       int
	Title        string
	RenderTitle  bool
	TitleStyle   *style.Style
	ContentStyle *style.Style
}

type Base struct {
	Config     *Config
	width      int
	height     int
	lines      []*Line
	hasChanged func()
}

func NewBasicPanel(conf *Config) *Base {
	if conf.Height < 1 {
		panic("panels height can not be less than 1 row")
	}
	height := conf.Height
	if conf.RenderTitle {
		height = conf.Height - 1
	}
	if conf.ContentStyle == nil {
		conf.ContentStyle = &style.Style{}
	}
	if conf.TitleStyle == nil {
		conf.TitleStyle = &style.Style{}
	}
	return &Base{
		width:      conf.Width,
		height:     height,
		Config:     conf,
		lines:      NewLines(height),
		hasChanged: func() {},
	}
}

func (bp *Base) Write(index int, line string, optionalStyle ...*style.Style) error {
	sty := bp.Config.ContentStyle
	if len(optionalStyle) > 0 {
		sty = optionalStyle[0]
	}
	if index >= (bp.height) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.height)
	}
	bp.lines[index] = &Line{Line: "", Style: sty}
	bp.renderLine(index)
	bp.hasChanged()
	return nil
}

func (bp *Base) Clear() {
	bp.lines = NewLines(bp.height)
	bp.render()
	bp.hasChanged()
}

func (bp *Base) ClearLine(index int) {
	bp.lines[index] = &Line{Line: "", Style: &style.Style{}}
}

func (bp *Base) GetSize() (int, int) {
	return bp.width, bp.height
}

func (bp *Base) GetBuffer() []*Line {
	return bp.lines
}

func (bp *Base) ChangeHandler(f func()) {
	if f == nil {
		return
	}
	bp.hasChanged = f
}

func (bp *Base) GetConfig() *Config {
	return bp.Config
}

func (bp *Base) render() {
	for i := range bp.lines {
		bp.renderLine(i)
	}
	bp.hasChanged()
}

func (bp *Base) renderLine(index int) {
	line := bp.lines[index]
	line.Line = ansi.Strip(line.Line)
	bp.lines[index] = line
}
