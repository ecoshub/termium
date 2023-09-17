package screen

import (
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
)

const (
	DefaultCommandPalettePrompt string = ">> "
)

type ComponentConfig struct {
	PosX int
	PosY int
}

type Component struct {
	posX int
	posY int
	p    panel.Panel
}

// ConstantText add constant text to string if you don't want to change colors just enter zeros (0)
func (s *Screen) ConstantText(posX, posY int, line string, sty *style.Style) {
	p := panel.ConstantText(line, sty)
	s.Add(p, posX, posY)
}

func (s *Screen) Add(p panel.Panel, posX, posY int) {
	pSizeX, pSizeY := p.GetSize()
	if posX+pSizeX > utils.TerminalWith {
		panic("panel width exceeds current windows")
	}
	if posY+pSizeY > utils.TerminalHeight {
		panic("panel height exceeds current windows")
	}

	s.renderer.components = append(s.renderer.components, &Component{p: p, posX: posX, posY: posY})

	p.ChangeHandler(func() { s.renderer.Render() })
}
