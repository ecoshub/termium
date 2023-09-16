package screen

import (
	"time"

	"github.com/ecoshub/termium/component/style"
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
	s.RenderCommandPalette()
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
		s.RenderCommandPalette()
		return
	}

	s.render()
	s.RenderCommandPalette()
}

func (s *Screen) render() {
	print(ansi.SaveCursorPos)
	defer print(ansi.RestoreCursorPos)

	print(ansi.GoToFirstBlock)

	s.readComponents()
}

func (s *Screen) readComponents() {
	for _, c := range s.components {
		s.readComponent(c)
	}
}

func (s *Screen) readComponent(c *Component) {
	sizeX, sizeY := c.p.GetSize()
	buffer := c.p.GetBuffer()
	panelConfig := c.p.GetConfig()

	offset := 0
	// render the title
	if panelConfig.RenderTitle {
		offset = 1
	}

	// Clear before write component buffer
	for i := 0; i < sizeY+offset; i++ {
		ansi.GotoRowAndColumn(c.posY+i+1, c.posX)
		blank := utils.FixedSizeLine("", sizeX)
		print(string(blank))
	}

	// render the title
	if panelConfig.RenderTitle {
		// go to title position and clear
		ansi.GotoRowAndColumn(c.posY+1, c.posX)
		blank := utils.FixedSizeLine("", sizeX)
		print(string(blank))

		// go to title position again to write title
		ansi.GotoRowAndColumn(c.posY+1, c.posX)
		line := ansi.ClearLine(panelConfig.Title, sizeX)
		line = style.SetStyle(line, panelConfig.TitleStyle)
		print(line)
	}

	for i := 0; i < sizeY; i++ {
		ansi.GotoRowAndColumn(c.posY+i+offset+1, c.posX)
		line := ansi.ClearLine(string(buffer[i]), sizeX)
		line = style.SetStyle(line, panelConfig.ContentStyle)
		print(line)
	}

}

func (s *Screen) RenderCommandPalette() {
	ansi.GotoRowAndColumn(s.defaultCursorPosY, s.defaultCursorPosX)
	if s.CommandPalette == nil {
		return
	}

	print(ansi.EraseLine)
	print(s.CommandPalette.String())
	ansi.GotoRowAndColumn(s.defaultCursorPosY, s.CommandPalette.GetCursorIndex()+1)
}
