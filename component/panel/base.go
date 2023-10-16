package panel

import (
	"fmt"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/line"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
)

type Base struct {
	Config     *config.Config
	lines      []*line.Line
	hasChanged func()
}

func NewBasicPanel(conf *config.Config) *Base {
	if conf.Height < 1 {
		panic("panels height can not be less than 1 row")
	}
	height := conf.Height
	if conf.RenderTitle {
		conf.Height--
	}
	if conf.ContentStyle == nil {
		conf.ContentStyle = &style.Style{}
	}
	if conf.TitleStyle == nil {
		conf.TitleStyle = &style.Style{}
	}
	return &Base{
		Config:     conf,
		lines:      line.NewLines(height),
		hasChanged: func() {},
	}
}

func (bp *Base) Write(index int, input string, optionalStyle ...*style.Style) error {
	sty := bp.Config.ContentStyle
	if len(optionalStyle) > 0 {
		sty = optionalStyle[0]
	}
	if index >= (bp.Config.Height) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.Config.Height)
	}
	bp.lines[index] = &line.Line{Line: input, Style: sty}
	bp.renderLine(index)
	bp.hasChanged()
	return nil
}

func (bp *Base) WriteMulti(lines []string) error {
	for i := 0; i < bp.Config.Height; i++ {
		if i >= len(lines) {
			break
		}
		bp.lines[i] = &line.Line{Line: lines[i], Style: &style.Style{}}
	}
	bp.render()
	bp.hasChanged()
	return nil
}

func (bp *Base) WriteMultiStyle(lines []string, sty *style.Style) error {
	for i := 0; i < bp.Config.Height; i++ {
		if i >= len(lines) {
			break
		}
		bp.lines[i] = &line.Line{Line: lines[i], Style: sty}
	}
	bp.render()
	bp.hasChanged()
	return nil
}

func (bp *Base) Clear() {
	bp.lines = line.NewLines(bp.Config.Height)
	bp.render()
	bp.hasChanged()
}

func (bp *Base) ClearLine(index int) {
	bp.lines[index] = line.Empty()
}

func (bp *Base) Buffer() []*line.Line {
	return bp.lines
}

func (bp *Base) ListenChangeHandler(f func()) {
	if f == nil {
		return
	}
	bp.hasChanged = f
}

func (bp *Base) Configuration() *config.Config {
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
