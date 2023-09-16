package main

import (
	"fmt"

	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/screen"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New(&screen.CommandPaletteConfig{Prompt: "eco$ ", Enable: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	sp := panel.NewStackPanel(10, 10)
	s.Add(sp, &screen.ComponentConfig{})

	sp1 := panel.NewStackPanel(10, 10)
	s.Add(sp1, &screen.ComponentConfig{PosX: 11})

	selector := 0

	// command handler
	s.CommandPalette.ListenActions(func(a *screen.KeyAction) {
		switch a.Action {
		// enter event handler
		case screen.KeyActionEnter:

			// also add command pallets own history module to select with up/down arrow keys later
			s.CommandPalette.AddToHistory(a.Input)

			switch a.Input {
			case "clear":
				s.CommandPalette.ClearHistory()
				sp.Clear()
			default:
				msg := fmt.Sprintf("\x1b[38;5;197m%s\x1b[0m ", a.Input)
				if selector == 0 {
					sp.Push(msg)
				} else {
					sp1.Push(msg)
				}
			}
			selector = (selector + 1) % 2
		}
	})

	s.Start()

	select {}
}
