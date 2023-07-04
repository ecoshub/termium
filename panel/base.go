package panel

import (
	"fmt"

	"github.com/ecoshub/termium/utils"
)

type BasePanel struct {
	Config *Config
	buffer [][]rune
	lines  []string
}

func NewBasePanel(conf *Config) *BasePanel {
	bp := &BasePanel{
		Config: conf,
		buffer: utils.InitRuneMatrix(conf.SizeX, conf.SizeY, ' '),
		lines:  make([]string, conf.SizeY),
	}
	return bp
}

func (bp *BasePanel) Insert(index int, line string) error {
	if index >= (bp.Config.SizeY) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.Config.SizeY)
	}
	bp.lines[index] = line
	bp.renderLine(index)
	return nil
}

func (bp *BasePanel) Clear() {
	bp.lines = make([]string, bp.Config.SizeY)
	bp.buffer = utils.InitRuneMatrix(bp.Config.SizeX, bp.Config.SizeY, ' ')
	bp.renderList()
}

func (bp *BasePanel) ClearLine(index int) {
	bp.buffer[index] = utils.InitRuneArray(bp.Config.SizeX, ' ')
}

func (bp *BasePanel) GetSize() (int, int) {
	return bp.Config.SizeX, bp.Config.SizeY
}

func (bp *BasePanel) GetBuffer() [][]rune {
	return bp.buffer
}

func (bp *BasePanel) renderList() {
	for i := range bp.lines {
		bp.renderLine(i)
	}
}

func (bp *BasePanel) renderLine(index int) {
	line := bp.lines[index]
	r := FixedSizeLine(line, bp.Config.SizeX)
	bp.buffer[index] = []rune(r)
}
