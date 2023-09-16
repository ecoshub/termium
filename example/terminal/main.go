package main

import (
	"fmt"

	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/screen"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a stack panel to use as a command history
	historyPanel := panel.NewStackPanel(screen.TerminalWith, 5)

	// creating  command pallet
	s.CreateCommandPallet(&screen.CommandPaletteConfig{
		Prompt: "ecoshub$ ",
	})

	// command handler
	s.CommandPalette.ListenActions(func(a *screen.KeyAction) {
		switch a.Action {
		// enter event handler
		case screen.KeyActionEnter:

			// append input in to history panel
			historyPanel.Push(a.Input)

			// also add command pallets own history module to select with up/down arrow keys later
			s.CommandPalette.AddToHistory(a.Input)

			// lets add a command
			// if command is clear. clear the history pallet
			if a.Input == "clear" {
				historyPanel.Clear()
				s.CommandPalette.ClearHistory()
			}
			s.Render()
		case screen.KeyActionInnerEvent:
			s.Render()
		}
	})

	// lets add this panel to top left corner (0,0)
	s.Add(historyPanel, &screen.ComponentConfig{
		Title:       "History:",
		RenderTitle: true,
		PosX:        0,
		// 7 is panel height (5) + terminal height(1) + history panel title(1)
		PosY: screen.TerminalHeight - 7,
	})

	s.Start()

	select {}
}
