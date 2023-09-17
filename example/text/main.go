package main

import (
	"fmt"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New(&screen.Config{
		CommandPaletteConfig: &palette.Config{
			Prompt: "prompt# ",
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
	tp := panel.NewTextPanel(&panel.Config{
		Width:       20,
		Height:      10,
		Title:       "TextBox_1:",
		RenderTitle: true,
		TitleStyle: &style.Style{
			BackgroundColor: 123,
		},
		ContentStyle: &style.Style{
			ForegroundColor: 195,
		},
	})

	// lets add this panel to top left corner (0,0)
	s.Add(tp, 0, 0)

	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {
		tp.Append(input)
		s.CommandPalette.AddToHistory(input)
	})

	s.Start()
}
