package main

import (
	"fmt"
	"time"

	"github.com/ecoshub/termium/palette"
	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/screen"
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
		SizeX: screen.TerminalWith,
		SizeY: 5,
	}

	c1 := panel.NewStackPanel(scc)
	sc := &screen.ComponentConfig{
		Title:       "output:",
		PosX:        0,
		PosY:        screen.TerminalHeight - 5 - 1 - 1 - 1,
		RenderTitle: true,
	}

	s.AddNewComponent(c1, sc)

	scc2 := &panel.Config{
		SizeX: 45,
		SizeY: 5,
	}

	c2 := panel.NewStackPanel(scc2)
	sc2 := &screen.ComponentConfig{
		RenderTitle: true,
		Title:       "TEST PANEL 01:",
		PosX:        1,
		PosY:        1,
	}

	c3 := panel.NewStackPanel(&panel.Config{SizeX: 45, SizeY: 5})
	scc3 := &screen.ComponentConfig{
		RenderTitle: true,
		Title:       "TEST PANEL 02:",
		PosX:        46,
		PosY:        1,
	}

	go DummyData(c2)
	go DummyData(c3)

	s.AddNewComponent(c2, sc2)
	s.AddNewComponent(c3, scc3)

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
}

func defaultScreen() {

	s, err := screen.NewDefaultScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	c1 := panel.NewBasePanel(&panel.Config{SizeX: 32, SizeY: 5})
	s.AddNewComponent(c1, &screen.ComponentConfig{Title: "--- FIRST ----", PosX: 1, PosY: 1, RenderTitle: true})

	c2 := panel.NewBasePanel(&panel.Config{SizeX: 32, SizeY: 3})
	s.AddNewComponent(c2, &screen.ComponentConfig{Title: "--- SECOND ---", PosX: 1, PosY: 7, RenderTitle: true})

	c3 := panel.NewBasePanel(&panel.Config{SizeX: 64, SizeY: 20})
	s.AddNewComponent(c3, &screen.ComponentConfig{Title: "--- THIRD ---", PosX: 34, PosY: 1, RenderTitle: true})

	c4 := panel.NewBasePanel(&panel.Config{SizeX: 32, SizeY: 10})
	s.AddNewComponent(c4, &screen.ComponentConfig{Title: "--- FORTH ---", PosX: 1, PosY: 1, RenderTitle: true})

	s.AttachCommandPalletFunc(func(a *palette.Action) {})

	s.Start()
	s.RenderPeriodically(screen.DefaultRefreshDelay)
}

func DummyData(cp *panel.StackPanel) {
	c := 0
	for range time.NewTicker(time.Millisecond * 250).C {
		str := fmt.Sprintf("Hello this is the %d. entry of this panel", c)
		cp.Append(str)
		c++
	}
}
