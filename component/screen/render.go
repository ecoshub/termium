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

	if len(s.renderer.components) != 0 {
		print(ansi.ClearScreen)
		s.renderer.Render()
	}

	s.renderer.RenderCommandPalette()
	s.started = true

	utils.WaitInterrupt(func() {
		print(ansi.MakeCursorVisible)
	})
}

func (r *Renderer) Render() {
	if r.maxRenderTimeGap != 0 {
		if r.sinceLastRender() < r.maxRenderTimeGap {
			return
		}
		defer r.setLastRender()
	}

	print(ansi.MakeCursorInvisible)

	r.Lock()
	defer r.Unlock()

	r.renderCore()
}

func (s *Screen) Print(input string) {
	ansi.GotoRowAndColumn(s.TerminalHeight-1, 0)
	println()
	print(ansi.EraseLine)
	println(input)
	s.lineBuffer = input
	s.renderer.RenderCommandPalette()
}

func (s *Screen) AppendToLastLine(input string) {
	ansi.GotoRowAndColumn(s.TerminalHeight-1, len(ansi.Strip(s.lineBuffer))+1)
	s.lineBuffer += input
	println(input)
	s.renderer.RenderCommandPalette()
}

func (r *Renderer) renderCore() {
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

	if r.componentRendered[index] {
		return
	}

	defer func() { r.componentRendered[index] = true }()

	buffer := c.renderable.Buffer()
	panelConfig := c.renderable.Configuration()

	offset := 0
	// render the title
	if panelConfig.RenderTitle {
		if !r.componentTitleRenderer[index] {
			// go to title position and clear
			ansi.GotoRowAndColumn(c.posY+1, c.posX)
			blank := utils.FixedSizeLine("", panelConfig.Width)
			print(string(blank))

			// go to title position again to write title
			ansi.GotoRowAndColumn(c.posY+1, c.posX)
			line := ansi.ClearLine(panelConfig.Title, panelConfig.Width)
			missingChars := panelConfig.Width - len(line)
			for i := 0; i < missingChars; i++ {
				line += " "
			}
			line = style.SetStyle(line, panelConfig.TitleStyle)
			print(line)
			r.componentTitleRenderer[index] = true
		}
		offset = 1
	}

	for i := 0; i < panelConfig.Height; i++ {
		ansi.GotoRowAndColumn(c.posY+i+offset+1, c.posX)
		line := ansi.ClearLine(string(buffer[i].Line), panelConfig.Width)
		line = string(utils.FixedSizeLine(line, panelConfig.Width))
		line = style.SetStyle(line, buffer[i].Style)
		print(line)
	}

}

func (r *Renderer) RenderCommandPalette() {
	if !r.renderCommandPallet {
		return
	}

	if !r.commandPalette.PromptLine.IsDirty() {
		return
	}

	print(ansi.MakeCursorInvisible)
	defer print(ansi.MakeCursorVisible)

	r.Lock()
	defer r.Unlock()

	ansi.GotoRowAndColumn(r.terminalHeight, 0)
	if r.commandPalette == nil {
		return
	}

	print(ansi.EraseLine)
	print(r.commandPalette.Prompt())
	print(r.commandPalette.Input())
	ansi.GotoRowAndColumn(r.terminalHeight, len(r.commandPalette.Config.Prompt)+r.commandPalette.PromptLine.GetCursorIndex()+1)
	r.commandPalette.PromptLine.Rendered()

}

func (r *Renderer) setLastRender() {
	r.lastRender = time.Now()
}

func (r *Renderer) sinceLastRender() time.Duration {
	return time.Since(r.lastRender)
}
