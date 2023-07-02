package screen

import (
	"term/internal/models/dimension"
	"term/internal/panel"
)

type SectionConfig struct {
	Title       string
	Pos         *dimension.D2
	RenderTitle bool
}

type Section struct {
	p    panel.Panel
	conf *SectionConfig
}

func (s *Screen) AddNewComponent(p panel.Panel, sc *SectionConfig) {
	if sc.Pos.X+p.GetSize().X > s.size.X {
		panic("component is not fitting on current terminal size")
	}
	if sc.Pos.Y+p.GetSize().Y > s.size.Y {
		panic("component is not fitting on current terminal size")
	}
	s.sections = append(s.sections, &Section{p: p, conf: sc})
}
