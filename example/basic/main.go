package main

import (
	"fmt"
	"strings"

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
				ForegroundColor: 83,
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

		if strings.HasPrefix(input, "*") {
			input = strings.TrimPrefix(input, "*")
			historyPanel.Push(input, style.DefaultStyleInfo)
			return
		}

		if strings.HasPrefix(input, "?") {
			input = strings.TrimPrefix(input, "?")
			historyPanel.Push(input, style.DefaultStyleWarning)
			return
		}

		if strings.HasPrefix(input, "!") {
			input = strings.TrimPrefix(input, "!")
			historyPanel.Push(input, style.DefaultStyleError)
			return
		}

		if strings.HasPrefix(input, ":dump ") {
			tokens := strings.Split(input, " ")
			if len(tokens) != 2 {
				historyPanel.Push("unsupported number of argument for 'dump' command", style.DefaultStyleError)
				return
			}
			path := tokens[1]
			err := historyPanel.Dump(path)
			if err != nil {
				historyPanel.Push("'dump' command error. err: "+err.Error(), style.DefaultStyleError)
				return
			}
			return
		}

		if strings.HasPrefix(input, ":flush") {
			tokens := strings.Split(input, " ")
			if len(tokens) != 1 {
				historyPanel.Push("unsupported number of argument for 'dump' command", style.DefaultStyleError)
				return
			}
			historyPanel.Flush()
			return
		}

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
