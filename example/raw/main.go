package main

import (
	"fmt"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
	"github.com/eiannone/keyboard"
)

func main() {
	s, err := screen.New(&screen.Config{
		DisableCommentPallet: true,
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

	mainPanel := panel.NewRawPanel(&config.Config{
		Width:  80,
		Height: 24,
		ContentStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 232,
		},
	})

	s.Add(mainPanel, 0, 0)

	index := 0
	s.CommandPalette.SetBaseListener(false)
	s.CommandPalette.AttachKeyEventHandler(func(event keyboard.KeyEvent) {
		input := rune(event.Key)
		if event.Key == 0 {
			input = rune(event.Rune)
		}
		mainPanel.Write(index, input)
		index++
	})

	s.CommandPalette.ListenKeyEventEnter(func(input string) {
		s.CommandPalette.AddToHistory(input)
	})

	s.Start()

}
