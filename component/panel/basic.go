package panel

import (
	"fmt"

	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
)

type Config struct {
	Width  int
	Height int
}

type Panel interface {
	GetSize() (int, int)
	GetBuffer() [][]rune
	ChangeHandler(h func())
}

type Basic struct {
	Config     *Config
	buffer     [][]rune
	lines      []string
	hasChanged func()
}

func NewBasicPanel(width, height int) *Basic {
	return &Basic{
		Config: &Config{Width: width, Height: height},
		buffer: utils.InitRuneMatrix(width, height, ' '),
		lines:  make([]string, height),
	}
}

func (bp *Basic) Write(index int, line string) error {
	if index >= (bp.Config.Height) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.Config.Height)
	}
	bp.lines[index] = line
	bp.renderLine(index)
	bp.hasChanged()
	return nil
}

func (bp *Basic) Clear() {
	bp.lines = make([]string, bp.Config.Height)
	bp.buffer = utils.InitRuneMatrix(bp.Config.Width, bp.Config.Height, ' ')
	bp.render()
	bp.hasChanged()
}

func (bp *Basic) ClearLine(index int) {
	bp.buffer[index] = utils.InitRuneArray(bp.Config.Width, ' ')
}

func (bp *Basic) GetSize() (int, int) {
	return bp.Config.Width, bp.Config.Height
}

func (bp *Basic) GetBuffer() [][]rune {
	return bp.buffer
}

func (bp *Basic) ChangeHandler(f func()) {
	bp.hasChanged = f
}

func (bp *Basic) render() {
	for i := range bp.lines {
		bp.renderLine(i)
	}
}

func (bp *Basic) renderLine(index int) {
	line := bp.lines[index]
	line = ansi.Strip(line)
	r := utils.FixedSizeLine(line, bp.Config.Width)
	bp.buffer[index] = []rune(r)
}
