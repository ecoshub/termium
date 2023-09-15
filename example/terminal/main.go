package main

import (
	"fmt"

	"github.com/ecoshub/termium/palette"
	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/screen"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.NewScreen()
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a stack panel to use as a command history
	historyPanel := panel.NewStack(screen.TerminalWith, 5)

	// creating  command pallet
	s.CreateCommandPallet(&palette.CommandPaletteConfig{
		Width:           screen.TerminalWith,
		PromptString:    "ecoshub$ ",
		HistoryCapacity: palette.DefaultHistoryCapacity,
	})

	// command handler
	s.CommandPalette.ListenActions(func(a *palette.Action) {
		switch a.Action {
		// enter event handler
		case palette.ActionEnter:

			// append input in to history panel
			historyPanel.Push(a.Input)

			// also add command pallets own history module to select with arrow keys later
			s.CommandPalette.AddToHistory(a.Input)

			// lets add a command
			// if command is clear. clear the history pallet
			if a.Input == "clear" {
				historyPanel.Clear()
				s.CommandPalette.ClearHistory()
			}
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

	s.Run()
}
