package screen

import (
	"sync"
	"time"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
)

type Renderer struct {
	sync.Mutex
	terminalWidth          int
	terminalHeight         int
	components             []*Component
	componentRendered      map[int]bool
	componentTitleRenderer map[int]bool
	commandPalette         *palette.Palette
	renderCommandPallet    bool
	minRenderTimeGap       time.Duration
	queue                  chan struct{}
	fl                     *utils.FileLogger
}

func (r *Renderer) Routine() {
	t := time.NewTicker(r.minRenderTimeGap)
	for range t.C {
		<-r.queue
		r.render()
	}
}

func (r *Renderer) Render() {
	r.queue <- struct{}{}
}

func (r *Renderer) render() {
	r.Lock()
	defer r.Unlock()

	print(ansi.MakeCursorInvisible)
	defer print(ansi.MakeCursorVisible)

	print(ansi.SaveCursorPos)
	defer print(ansi.RestoreCursorPos)

	r.renderComponents()
}

func (r *Renderer) renderComponents() {
	for i := range r.components {
		r.RenderComponent(i)
	}
}

func (r *Renderer) RenderComponent(index int) {

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
			line := ansi.ClearLine(panelConfig.Title)
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
		line := ansi.ClearLine(string(buffer[i].Line))
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

func (r *Renderer) setRenderedLock(index int, enable bool) {
	r.Lock()
	defer r.Unlock()

	r.componentRendered[index] = enable
}
