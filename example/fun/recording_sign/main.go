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

	recordingPanel := panel.NewBasicPanel(&config.Config{
		Width:  1,
		Height: 1,
		ContentStyle: &style.Style{
			ForegroundColor: style.DefaultStyleError.ForegroundColor,
		},
	})

	s.Add(recordingPanel, utils.TerminalWith-1, 0)

	recordingPanel.Write(0, " ")

	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {

		// also add command pallets own history module to select with up/down arrow keys later
		s.CommandPalette.AddToHistory(input)

		switch input {
		case ":clear":
			s.CommandPalette.ClearHistory()
			s.ResetScreen()
			recordingPanel.Write(0, " ")
			recordingPanel.Config.ContentStyle.Blink = false
		case ":record":
			recordingPanel.Config.ContentStyle.Blink = true
			recordingPanel.Write(0, "●")
		case ":stop":
			recordingPanel.Write(0, " ")
			recordingPanel.Config.ContentStyle.Blink = false
		}

	})

	s.Start()
}
