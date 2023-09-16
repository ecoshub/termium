package screen

import (
	"strings"
	"time"

	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
)

func (s *Screen) Start() {
	if s.started {
		return
	}

	if len(s.components) == 0 {
		panic("no component added to screen")
	}

	print(ansi.ClearScreen)
	s.drawCommandPallet()
	s.started = true
	s.Render()
	utils.WaitInterrupt(nil)
}

func (s *Screen) Render() {
	defer func() {
		s.lastRender = time.Now()
	}()

	print(ansi.MakeCursorInvisible)
	defer print(ansi.MakeCursorVisible)

	if s.CommandPalette == nil {
		s.render()
		s.drawCommandPallet()
		return
	}

	s.render()
	s.drawCommandPallet()
}

func (s *Screen) render() {
	print(ansi.SaveCursorPos)
	defer print(ansi.RestoreCursorPos)

	print(ansi.GoToFirstBlock)

	s.readComponents()
	print(s.String())

}

func (s *Screen) readComponents() {
	for _, c := range s.components {
		s.readComponent(c)
	}
}

func (s *Screen) readComponent(c *Component) {
	sizeX, sizeY := c.p.GetSize()
	buffer := c.p.GetBuffer()
	title := utils.FixedSizeLine(c.conf.Title, sizeX)

	offset := 0
	if c.conf.RenderTitle {
		// render title
		copy(s.buffer[c.conf.PosY][c.conf.PosX:c.conf.PosX+sizeX], title[:sizeX])
		offset += 1
	}

	// render component
	for i := 0; i < sizeY; i++ {
		copy(s.buffer[i+c.conf.PosY+offset][c.conf.PosX:c.conf.PosX+sizeX], buffer[i][:sizeX])
	}

}

func (s *Screen) String() string {
	builder := strings.Builder{}
	for i := range s.buffer {
		for j := range s.buffer[i] {
			builder.WriteRune(s.buffer[i][j])
		}
	}
	return builder.String()
}

func (s *Screen) drawCommandPallet() {
	ansi.GotoRowAndColumn(s.defaultCursorPosY, s.defaultCursorPosX)
	if s.CommandPalette == nil {
		return
	}

	print(ansi.EraseLine)
	print(s.CommandPalette.String())
	ansi.GotoRowAndColumn(s.defaultCursorPosY, s.CommandPalette.GetCursorIndex()+1)
}
