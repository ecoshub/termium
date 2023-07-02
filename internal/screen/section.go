package screen

import (
	"term/internal/component"
	"term/internal/models/dimension"
)

type SectionConfig struct {
	Title       string
	Pos         *dimension.D2
	RenderTitle bool
}

type Section struct {
	c    component.Component
	conf *SectionConfig
}

func (s *Screen) AddNewComponent(c component.Component, sc *SectionConfig) {
	if sc.Pos.X+c.GetSize().X > s.size.X {
		panic("component is not fitting on current terminal size")
	}
	if sc.Pos.Y+c.GetSize().Y > s.size.Y {
		panic("component is not fitting on current terminal size")
	}
	s.sections = append(s.sections, &Section{c: c, conf: sc})
}
