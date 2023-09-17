package screen

import (
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
)

func (s *Screen) Start() {
	if s.started {
		return
	}

	if len(s.renderer.components) == 0 {
		panic("no component added to screen")
	}

	print(ansi.ClearScreen)
	s.renderer.RenderCommandPalette()
	s.started = true
	s.renderer.Render()

	utils.WaitInterrupt(nil)
}

func (r *Renderer) Render() {
	r.Lock()
	defer r.Unlock()

	r.renderCore()
}

func (r *Renderer) renderCore() {
	print(ansi.MakeCursorInvisible)
	defer print(ansi.MakeCursorVisible)

	r.render()
}

func (r *Renderer) render() {
	print(ansi.SaveCursorPos)
	defer print(ansi.RestoreCursorPos)

	r.readComponents()
}

func (r *Renderer) readComponents() {
	for i := range r.components {
		r.readComponent(i)
	}
}

func (r *Renderer) readComponent(index int) {
	c := r.components[index]

	sizeX, sizeY := c.p.GetSize()
	buffer := c.p.GetBuffer()
	panelConfig := c.p.GetConfig()

	offset := 0
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
		offset = 1
	}

	for i := 0; i < sizeY; i++ {
		ansi.GotoRowAndColumn(c.posY+i+offset+1, c.posX)
		line := ansi.ClearLine(string(buffer[i].Line), sizeX)
		line = string(utils.FixedSizeLine(line, sizeX))
		line = style.SetStyle(line, buffer[i].Style)
		print(line)
	}

}

func (r *Renderer) RenderCommandPalette() {
	r.Lock()
	defer r.Unlock()

	ansi.GotoRowAndColumn(utils.TerminalHeight, 0)
	if r.commandPalette == nil {
		return
	}

	print(ansi.EraseLine)
	print(r.commandPalette.String())
	ansi.GotoRowAndColumn(utils.TerminalHeight, len(r.commandPalette.Config.Prompt)+r.commandPalette.PromptLine.GetCursorIndex()+1)
}
