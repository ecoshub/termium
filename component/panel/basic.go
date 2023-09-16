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

type Panel interface {
	GetSize() (int, int)
	GetBuffer() []string
	GetConfig() *Config
	ChangeHandler(h func())
}

type Basic struct {
	Config *Config

	width      int
	height     int
	lines      []string
	hasChanged func()
}

func NewBasicPanel(conf *Config) *Basic {
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
	return &Basic{
		width:  conf.Width,
		height: height,
		Config: conf,
		lines:  make([]string, height),
	}
}

func (bp *Basic) Write(index int, line string) error {
	if index >= (bp.height) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.height)
	}
	bp.lines[index] = line
	bp.renderLine(index)
	bp.hasChanged()
	return nil
}

func (bp *Basic) Clear() {
	bp.lines = make([]string, bp.height)
	bp.render()
	bp.hasChanged()
}

func (bp *Basic) ClearLine(index int) {
	bp.lines[index] = ""
}

func (bp *Basic) GetSize() (int, int) {
	return bp.width, bp.height
}

func (bp *Basic) GetBuffer() []string {
	return bp.lines
}

func (bp *Basic) ChangeHandler(f func()) {
	bp.hasChanged = f
}

func (bp *Basic) GetConfig() *Config {
	return bp.Config
}

func (bp *Basic) render() {
	for i := range bp.lines {
		bp.renderLine(i)
	}
	bp.hasChanged()
}

func (bp *Basic) renderLine(index int) {
	line := bp.lines[index]
	line = ansi.Strip(line)
	bp.lines[index] = line
}
