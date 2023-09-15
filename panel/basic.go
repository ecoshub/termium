package panel

import (
	"fmt"

	"github.com/ecoshub/termium/utils"
)

type Config struct {
	Width  int
	Height int
}

type Panel interface {
	GetSize() (int, int)
	GetBuffer() [][]rune
}

type Basic struct {
	Config *Config
	buffer [][]rune
	lines  []string
}

func NewBasic(width, height int) *Basic {
	bp := &Basic{
		Config: &Config{Width: width, Height: height},
		buffer: utils.InitRuneMatrix(width, height, ' '),
		lines:  make([]string, height),
	}
	return bp
}

func (bp *Basic) Write(index int, line string) error {
	if index >= (bp.Config.Height) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.Config.Height)
	}
	bp.lines[index] = line
	bp.renderLine(index)
	return nil
}

func (bp *Basic) Clear() {
	bp.lines = make([]string, bp.Config.Height)
	bp.buffer = utils.InitRuneMatrix(bp.Config.Width, bp.Config.Height, ' ')
	bp.renderList()
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

func (bp *Basic) renderList() {
	for i := range bp.lines {
		bp.renderLine(i)
	}
}

func (bp *Basic) renderLine(index int) {
	line := bp.lines[index]
	r := FixedSizeLine(line, bp.Config.Width)
	bp.buffer[index] = []rune(r)
}
