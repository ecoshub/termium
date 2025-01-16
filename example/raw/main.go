package main

import (
	"fmt"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
	"github.com/eiannone/keyboard"
)

func main() {
	s, err := screen.New(&screen.Config{
		DisableCommentPallet: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	raw := panel.NewRawPanel(&config.Config{
		Width:  80,
		Height: 24,
		ContentStyle: &style.Style{
			ForegroundColor: 195,
			BackgroundColor: 232,
		},
	})

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

	posX := (utils.TerminalWith - raw.Config.Width) / 2

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
	s.CommandPalette.AttachKeyEventHandler(func(event keyboard.KeyEvent) {
		input := uint8(event.Key)
		if event.Key == 0 {
			input = uint8(event.Rune)
		}
		raw.Write(index, input)
		index++
	})

	s.Start()

}
