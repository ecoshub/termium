package panel

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ecoshub/termium/internal/models/dimension"
	"github.com/ecoshub/termium/internal/utils"
)

type BasePanel struct {
	Config *Config
	buffer [][]rune
	lines  []string
}

func NewBasePanel(conf *Config) *BasePanel {
	bp := &BasePanel{
		Config: conf,
		buffer: utils.InitRuneMatrix(conf.Size.X, conf.Size.Y, ' '),
		lines:  make([]string, conf.Size.Y),
	}
	go func() {
		n := rand.Intn(26)
		count := 0
		for range time.NewTicker(time.Millisecond * 500).C {
			bp.buffer = utils.InitRuneMatrix(bp.Config.Size.X, bp.Config.Size.Y, rune(n+97))
			bp.Insert(count, "\x1b[1;31mHello\x1b[0m my name is eco")
			count = (count + 1) % (bp.Config.Size.Y)
		}
	}()
	return bp
}

func (bp *BasePanel) Insert(index int, line string) error {
	if index >= (bp.Config.Size.Y) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, bp.Config.Size.Y)
	}
	bp.lines[index] = line
	bp.renderLine(index)
	return nil
}

func (bp *BasePanel) Clear() {
	bp.lines = make([]string, bp.Config.Size.Y)
	bp.buffer = utils.InitRuneMatrix(bp.Config.Size.X, bp.Config.Size.Y, ' ')
	bp.renderList()
}

func (bp *BasePanel) ClearLine(index int) {
	bp.buffer[index] = utils.InitRuneArray(bp.Config.Size.X, ' ')
}

func (bp *BasePanel) GetSize() *dimension.Vector {
	return bp.Config.Size
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
	r := FixedSizeLine(line, bp.Config.Size.X)
	bp.buffer[index] = []rune(r)
}
