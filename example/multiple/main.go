package main

import (
	"fmt"
	"strings"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
)

func main() {
	s, err := screen.New(&palette.CommandPaletteConfig{Prompt: ">> ", ForegroundColor: 47})
	if err != nil {
		fmt.Println(err)
		return
	}

	sp := panel.NewStackPanel(&panel.Config{
		Width:                10,
		Height:               10,
		Title:                "Panel_1:",
		RenderTitle:          true,
		TitleBackgroundColor: 22,
		ForegroundColor:      197,
	})
	s.Add(sp, 0, 0)

	sp1 := panel.NewStackPanel(&panel.Config{
		Width:                10,
		Height:               10,
		Title:                "Panel_2:",
		RenderTitle:          true,
		TitleBackgroundColor: 95,
		ForegroundColor:      227,
	})

	s.Add(sp1, 11, 0)

	s.CommandPalette.ListenKeyEventEnter(func(input string) {
		switch input {
		case "clear":
			sp.Clear()
			sp1.Clear()
		default:
			tokens := strings.Split(input, " ")
			for i, word := range tokens {
				if i%2 == 0 {
					sp.Push(word)
					continue
				}
				sp1.Push(word)
			}
			s.CommandPalette.AddToHistory(input)
		}
	})

	s.Start()
}
