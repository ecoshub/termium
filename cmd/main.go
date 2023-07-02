package main

import (
	"fmt"
	"term/internal/models/dimension"
	"term/internal/palette"
	"term/internal/panel"
	"term/internal/screen"
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

	scc := &panel.Config{
		Size: &dimension.D2{
			X: screen.TerminalWith,
			Y: 5,
		},
	}

	c1 := panel.NewStackPanel(scc)
	sc := &screen.SectionConfig{
		Title: "output:",
		Pos: &dimension.D2{
			X: 0,
			Y: screen.TerminalHeight - 5 - 1 - 1 - 1,
		},
		RenderTitle: true,
	}

	s.AddNewComponent(c1, sc)

	scc2 := &panel.Config{
		Size: &dimension.D2{
			X: 100,
			Y: 5,
		},
		AutoDummyInput: true,
	}

	c2 := panel.NewStackPanel(scc2)
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

	s.AttachCommandPalletFunc(func(a *palette.Action) {
		switch a.Action {
		case palette.ActionEnter:
			c1.Append(" > " + a.Input)
			cp.AddToHistory(a.Input)
			if a.Input == "clear" {
				c1.Clear()
			}
		}
		s.Render()
	})

	s.Start()
	s.ShowFPS = true
	s.RenderPeriodically(screen.DefaultRefreshDelay)

	select {}
}

func defaultScreen() {

	s, err := screen.NewDefaultScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	c1 := panel.NewBasePanel(&panel.Config{Size: &dimension.D2{X: 32, Y: 5}})
	s.AddNewComponent(c1, &screen.SectionConfig{Title: "--- FIRST ----", Pos: &dimension.D2{X: 1, Y: 1}, RenderTitle: true})

	c2 := panel.NewBasePanel(&panel.Config{Size: &dimension.D2{X: 32, Y: 3}})
	s.AddNewComponent(c2, &screen.SectionConfig{Title: "--- SECOND ---", Pos: &dimension.D2{X: 1, Y: 7}, RenderTitle: true})

	c3 := panel.NewBasePanel(&panel.Config{Size: &dimension.D2{X: 64, Y: 20}})
	s.AddNewComponent(c3, &screen.SectionConfig{Title: "--- THIRD ---", Pos: &dimension.D2{X: 34, Y: 1}, RenderTitle: true})

	c4 := panel.NewBasePanel(&panel.Config{Size: &dimension.D2{X: 32, Y: 10}})
	s.AddNewComponent(c4, &screen.SectionConfig{Title: "--- FORTH ---", Pos: &dimension.D2{X: 1, Y: 11}, RenderTitle: true})

	s.AttachCommandPalletFunc(func(a *palette.Action) {})

	s.Start()
	s.RenderPeriodically(screen.DefaultRefreshDelay)
}
