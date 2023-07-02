package main

import (
	"fmt"
	"math/rand"
	"strings"
	"term/internal/component/basic"
	"term/internal/component/stack"
	"term/internal/models/dimension"
	"term/internal/palette"
	"term/internal/screen"
	"time"
)

func main() {
	// defaultScreen()
	customScreen()
}

func customScreen() {

	s, err := screen.NewScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	scc := &stack.Config{
		Size: &dimension.D2{
			X: screen.TerminalWith,
			Y: 5,
		},
	}

	c1 := stack.New(scc)
	sc := &screen.SectionConfig{
		Title: "output:",
		Pos: &dimension.D2{
			X: 0,
			Y: screen.TerminalHeight - 5 - 1 - 1 - 1,
		},
		RenderTitle: true,
	}

	s.AddNewComponent(c1, sc)

	scc2 := &stack.Config{
		Size: &dimension.D2{
			X: 100,
			Y: 5,
		},
		Dummy: true,
	}

	c2 := stack.New(scc2)
	sc2 := &screen.SectionConfig{
		Title: "TEST PANEL 01:",
		Pos: &dimension.D2{
			X: 1,
			Y: 1,
		},
		RenderTitle: true,
	}

	s.AddNewComponent(c2, sc2)

	cp, err := palette.New(&palette.CommandPaletteConfig{
		Width:           screen.TerminalWith,
		PromptString:    "$$--> ",
		HistoryCapacity: palette.DefaultHistoryCapacity,
	})

	s.AttachCommandPallet(cp)

	// go useComponent(s, c1)

	s.AttachCommandPalletFunc(func(a *palette.Action) {
		switch a.Action {
		case palette.ActionEnter:
			element := strings.TrimPrefix(a.Input, "add ")
			c1.Append(" > " + element)
		}
	})

	s.Start()
	s.ShowFPS = true
	s.RenderPeriodically(screen.DefaultRefreshDelay)

	select {}
}

func useComponent(s *screen.Screen, c *basic.Component) {
	var r int
	var n int
	for {
		r = rand.Intn(300) + 20
		n = rand.Intn(c.GetSize().Y)

		time.Sleep(time.Duration(r) * time.Millisecond)
		c.Insert(n, fmt.Sprintf("hello_%d", r))
		time.Sleep(time.Duration(r) * time.Millisecond)
		s.Render()
	}
}

func defaultScreen() {

	s, err := screen.NewDefaultScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	c1 := basic.New(&basic.Config{Title: "--- FIRST ----", Size: &dimension.D2{X: 32, Y: 5}})
	s.AddNewComponent(c1, &screen.SectionConfig{Pos: &dimension.D2{X: 1, Y: 1}, RenderTitle: true})

	c2 := basic.New(&basic.Config{Title: "--- SECOND ---", Size: &dimension.D2{X: 32, Y: 3}})
	s.AddNewComponent(c2, &screen.SectionConfig{Pos: &dimension.D2{X: 1, Y: 7}, RenderTitle: true})

	c3 := basic.New(&basic.Config{Title: "--- THIRD ---", Size: &dimension.D2{X: 64, Y: 20}})
	s.AddNewComponent(c3, &screen.SectionConfig{Pos: &dimension.D2{X: 34, Y: 1}, RenderTitle: true})

	c4 := basic.New(&basic.Config{Title: "--- FORTH ---", Size: &dimension.D2{X: 32, Y: 10}})
	s.AddNewComponent(c4, &screen.SectionConfig{Pos: &dimension.D2{X: 1, Y: 11}, RenderTitle: true})

	s.AttachCommandPalletFunc(func(a *palette.Action) {})

	s.Start()
	s.RenderPeriodically(screen.DefaultRefreshDelay)
}
