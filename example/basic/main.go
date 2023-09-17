package main

import (
	"fmt"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New(&screen.Config{
		CommandPaletteConfig: &palette.Config{
			Prompt: "~ root# ",
			Style: &style.Style{
				ForegroundColor: 227,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a stack panel to use as a command history
	historyPanel := panel.NewStackPanel(&panel.Config{
		Width:  100,
		Height: utils.TerminalHeight - 2,
		Title:  "History:",
		ContentStyle: &style.Style{
			ForegroundColor: 195,
		},
	})

	// lets add this panel to top left corner (0,0)
	s.Add(historyPanel, 0, 0)

	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {

		// also add command pallets own history module to select with up/down arrow keys later
		s.CommandPalette.AddToHistory(input)

		// lets add a command
		// if command is clear. clear the history pallet
		if input == "clear" {
			historyPanel.Clear()
			s.CommandPalette.ClearHistory()
			s.ResetScreen()
			return
		}

		// append input in to history panel
		historyPanel.Push(input)
	})

	s.Start()
}
