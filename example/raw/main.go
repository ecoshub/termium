package main

import (
	"fmt"

	"github.com/ecoshub/termium/component/config"
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
			Prompt: "> ",
			Style: &style.Style{
				ForegroundColor: 83,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a stack panel to use as a command history
	raw := panel.NewRawPanel(&config.Config{
		Width:  80,
		Height: 24,
		ContentStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 232,
		},
	})

	// lets create a stack panel to use as a command history
	logPanel := panel.NewStackPanel(&config.Config{
		Width:  utils.TerminalWith,
		Height: utils.TerminalHeight - 24,
		Title:  "Logs:",
		TitleStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 235,
		},
		RenderTitle: true,
		ContentStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 235,
		},
	})

	// lets add this panel to top left corner (0,0)
	posX := (utils.TerminalWith - raw.Config.Width) / 2

	// lets create a stack panel to use as a command history
	memoryPanel := panel.NewStackPanel(&config.Config{
		Width:  posX - 1,
		Height: 24,
		Title:  "Memory:",
		TitleStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 234,
		},
		RenderTitle: true,
		ContentStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 234,
		},
	})

	s.Add(raw, posX, 0)
	s.Add(memoryPanel, 0, 0)
	s.Add(logPanel, 0, 24)

	index := 0
	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {

		// also add command pallets own history module to select with up/down arrow keys later
		s.CommandPalette.AddToHistory(input)

		// lets add a command
		// if command is clear. clear the history pallet
		if input == ":clear" {
			raw.Clear()
			logPanel.Clear()
			s.CommandPalette.ClearHistory()
			s.ResetScreen()
			return
		}

		for i := range input {
			raw.Write(index, input[i])
			index++
		}

		logPanel.Push(input)
	})

	s.Start()

}
