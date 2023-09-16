package main

import (
	"fmt"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/utils"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New(&palette.CommandPaletteConfig{Prompt: "$$>> ", ForegroundColor: 227})
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a stack panel to use as a command history
	historyPanel := panel.NewStackPanel(utils.TerminalWith, 5)

	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {

		// also add command pallets own history module to select with up/down arrow keys later
		s.CommandPalette.AddToHistory(input)

		// lets add a command
		// if command is clear. clear the history pallet
		if input == "clear" {
			historyPanel.Clear()
			s.CommandPalette.ClearHistory()
			return
		}

		// append input in to history panel
		historyPanel.Push(input)
	})

	// lets add this panel to top left corner (0,0)
	s.Add(historyPanel, &screen.ComponentConfig{
		Title:       "History:",
		RenderTitle: true,
		PosX:        0,
		// 7 is panel height (5) + terminal height(1) + history panel title(1)
		PosY: utils.TerminalHeight - 7,
	})

	s.Start()
}
