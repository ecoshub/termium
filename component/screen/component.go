package screen

import (
	"github.com/ecoshub/termium/component/panel"
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
func (s *Screen) ConstantText(posX, posY int, line string, foregroundColor, backgroundColor int) {
	p := panel.ConstantText(line, foregroundColor, backgroundColor)
	s.Add(p, posX, posY)
}

func (s *Screen) Add(p panel.Panel, posX, posY int) {
	pSizeX, pSizeY := p.GetSize()
	if posX+pSizeX > s.sizeX {
		panic("panel width exceeds current windows")
	}
	if posY+pSizeY > s.sizeY {
		panic("panel height exceeds current windows")
	}
	s.components = append(s.components, &Component{p: p, posX: posX, posY: posY})
	p.ChangeHandler(func() { s.Render() })
}
