package screen

import (
	"os"

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
	s.Add(p, &ComponentConfig{PosX: posX, PosY: posY})
}

func (s *Screen) Add(p panel.Panel, sc *ComponentConfig) {
	pSizeX, pSizeY := p.GetSize()
	if sc.PosX+pSizeX > s.sizeX {
		os.Exit(1)
	}
	if sc.PosY+pSizeY > s.sizeY {
		os.Exit(1)
	}
	s.components = append(s.components, &Component{p: p, conf: sc})
}
