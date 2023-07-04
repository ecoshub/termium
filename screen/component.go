package screen

import (
	"github.com/ecoshub/termium/panel"
)

type ComponentConfig struct {
	Title       string
	RenderTitle bool
	PosX        int
	PosY        int
}

type Component struct {
	p    panel.Panel
	conf *ComponentConfig
}

func (s *Screen) ConstantText(posX, posY int, line string) {
	p := panel.ConstantText(line)
	s.AddNewComponent(p, &ComponentConfig{PosX: posX, PosY: posY})
}

func (s *Screen) AddNewComponent(p panel.Panel, sc *ComponentConfig) {
	pSizeX, pSizeY := p.GetSize()
	if sc.PosX+pSizeX > s.sizeX {
		panic("component is not fitting on current terminal size")
	}
	if sc.PosY+pSizeY > s.sizeY {
		panic("component is not fitting on current terminal size")
	}
	s.components = append(s.components, &Component{p: p, conf: sc})
}
