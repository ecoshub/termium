package screen

import (
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/renderable"
	"github.com/ecoshub/termium/component/style"
)

const (
	DefaultCommandPalettePrompt string = ">> "
)

type Component struct {
	posX       int
	posY       int
	renderable renderable.Renderable
}

// FixedText add constant text to screen
func (s *Screen) FixedText(posX, posY int, line string, optionalStyle ...*style.Style) {
	var st *style.Style
	if len(optionalStyle) == 1 {
		st = optionalStyle[0]
	} else {
		st = &style.Style{}
	}
	p := panel.NewTextLine(line, st)
	s.Add(p, posX, posY)
}

// Add add renderable component to given screen position
func (s *Screen) Add(p renderable.Renderable, posX, posY int) {
	conf := p.Configuration()
	if posX+conf.Width > s.TerminalWidth {
		panic("panel width exceeds current windows")
	}
	if posY+conf.Height > s.TerminalHeight {
		panic("panel height exceeds current windows")
	}

	s.renderer.components = append(s.renderer.components, &Component{renderable: p, posX: posX, posY: posY})
	index := len(s.renderer.components) - 1
	p.ListenChangeHandler(func() {
		s.renderer.componentRendered[index] = false
		s.renderer.Render()
	})
}
