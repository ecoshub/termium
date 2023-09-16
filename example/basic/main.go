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
	s, err := screen.New(&palette.CommandPaletteConfig{Prompt: "~ root# ", Style: &style.Style{ForegroundColor: 227}})
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a stack panel to use as a command history
	historyPanel := panel.NewStackPanel(&panel.Config{
		Width:       utils.TerminalWith,
		Height:      5,
		Title:       "History:",
		RenderTitle: true,
		TitleStyle: &style.Style{
			BackgroundColor: 238,
			ForegroundColor: 197,
			Blink:           true,
			Bold:            true,
		},
		ContentStyle: &style.Style{
			ForegroundColor: 240,
		},
	})

	// lets add this panel to top left corner (0,0)
	s.Add(historyPanel, 0, utils.TerminalHeight-7)

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

	s.Start()
}
