package screen

import (
	"strings"
	"term/internal/ansi"
	"term/internal/panel"
	"term/internal/utils"
	"time"
)

func (s *Screen) RenderPeriodically(refreshRate time.Duration) {
	for range time.NewTicker(refreshRate).C {
		s.Render()
	}
}

func (s *Screen) Start() {
	if s.started {
		return
	}

	go ListenInterrupt()
	ansi.ClearScreen()
	s.drawCommandPallet()

	s.started = true
}

func (s *Screen) Render() {
	defer func() {
		s.lastRender = time.Now()
	}()

	if s.commandPalette == nil {
		s.render()
		s.drawCommandPallet()
		return
	}

	ansi.SaveCursorPos()
	s.render()
	ansi.RestoreCursorPos()
	s.drawCommandPallet()
}

func (s *Screen) render() {
	ansi.GoToFirstBlock()

	s.readComponents()
	print(s.String())

}

func (s *Screen) readComponents() {
	for _, c := range s.components {
		s.readComponent(c)
	}
	s.calculateFPS()
}

func (s *Screen) readComponent(c *Component) {
	size := c.p.GetSize()
	buffer := c.p.GetBuffer()
	title := panel.FixedSizeLine(c.conf.Title, size.X)

	offset := 0
	if c.conf.RenderTitle {
		// render title
		copy(s.buffer[c.conf.Position.Y][c.conf.Position.X:c.conf.Position.X+size.X], title[:size.X])
		offset += 1
	}

	// render component
	for j := 0; j < size.Y; j++ {
		copy(s.buffer[j+c.conf.Position.Y+offset][c.conf.Position.X:c.conf.Position.X+size.X], buffer[j][:size.X])
	}
}

func (s *Screen) String() string {
	builder := strings.Builder{}
	for i := range s.buffer {
		for j := range s.buffer[i] {
			builder.WriteRune(s.buffer[i][j])
		}
	}
	return builder.String()
}

func (s *Screen) drawCommandPallet() {
	ansi.Goto(s.defaultCursorPosY, s.defaultCursorPosX)
	if s.commandPalette == nil {
		return
	}
	ansi.EraseLine()
	pb := s.commandPalette.Buffer()
	print(pb)
	ansi.Goto(s.defaultCursorPosY, utils.PrintableLen(pb)+1)
}
