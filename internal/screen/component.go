package screen

import (
	"term/internal/models/dimension"
	"term/internal/panel"
)

type ComponentConfig struct {
	Title       string
	RenderTitle bool
	Position    *dimension.Vector
}

type Component struct {
	p    panel.Panel
	conf *ComponentConfig
}

func (s *Screen) AddNewComponent(p panel.Panel, sc *ComponentConfig) {
	if sc.Position.X+p.GetSize().X > s.size.X {
		panic("component is not fitting on current terminal size")
	}
	if sc.Position.Y+p.GetSize().Y > s.size.Y {
		panic("component is not fitting on current terminal size")
	}
	s.components = append(s.components, &Component{p: p, conf: sc})
}
