package screen

import (
	"fmt"
	"strings"
	"time"

	"github.com/ecoshub/termium/ansi"
	"github.com/ecoshub/termium/panel"
)

// Run async run command starts and renders periodically
func (s *Screen) Run() {
	s.Start()
	go s.RenderPeriodically(DefaultRefreshDelay)
}

func (s *Screen) RenderPeriodically(refreshRate time.Duration) {
	for range time.NewTicker(refreshRate).C {
		s.Render()
	}
}

func (s *Screen) Start() {
	defer s.Render()

	if s.started {
		return
	}

	if len(s.components) == 0 {
		panic("no component added to screen")
	}

	go WaitInterrupt(func() {
		fmt.Println("Exiting...")
	})

	ansi.ClearScreen()

	s.drawCommandPallet()

	s.started = true
}

func (s *Screen) Render() {
	defer func() {
		s.lastRender = time.Now()
	}()

	ansi.MakeCursorInvisible()
	defer ansi.MakeCursorVisible()

	if s.CommandPalette == nil {
		s.render()
		s.drawCommandPallet()
		return
	}

	s.render()
	s.drawCommandPallet()
}

func (s *Screen) render() {
	ansi.SaveCursorPos()
	defer ansi.RestoreCursorPos()

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
	sizeX, sizeY := c.p.GetSize()
	buffer := c.p.GetBuffer()
	title := panel.FixedSizeLine(c.conf.Title, sizeX)

	offset := 0
	if c.conf.RenderTitle {
		// render title
		copy(s.buffer[c.conf.PosY][c.conf.PosX:c.conf.PosX+sizeX], title[:sizeX])
		offset += 1
	}

	// render component
	for i := 0; i < sizeY; i++ {
		copy(s.buffer[i+c.conf.PosY+offset][c.conf.PosX:c.conf.PosX+sizeX], buffer[i][:sizeX])
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
	if s.CommandPalette == nil || !s.CommandPalette.Config.Enable {
		return
	}

	ansi.EraseLine()
	pb := s.CommandPalette.Buffer()
	print(pb)
	ansi.Goto(s.defaultCursorPosY, s.CommandPalette.cursorIndex+1)
}
