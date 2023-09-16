package main

import (
	"fmt"

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
		Width:                20,
		Height:               10,
		Title:                "Panel_1:",
		RenderTitle:          true,
		TitleBackgroundColor: 22,
		ForegroundColor:      197,
	})

	sp1 := panel.NewStackPanel(&panel.Config{
		Width:                20,
		Height:               10,
		Title:                "Panel_2:",
		RenderTitle:          true,
		TitleBackgroundColor: 95,
		ForegroundColor:      227,
	})

	sp2 := panel.NewStackPanel(&panel.Config{
		Width:                20,
		Height:               10,
		Title:                "Panel_3:",
		RenderTitle:          true,
		TitleBackgroundColor: 147,
		ForegroundColor:      14,
	})

	sp3 := panel.NewStackPanel(&panel.Config{
		Width:                20,
		Height:               10,
		Title:                "Panel_4:",
		RenderTitle:          true,
		TitleBackgroundColor: 88,
		ForegroundColor:      201,
	})

	s.Add(sp, 0, 0)
	s.Add(sp1, 22, 0)
	s.Add(sp2, 22, 11)
	s.Add(sp3, 0, 11)

	s.CommandPalette.ListenKeyEventEnter(func(input string) {
		switch input {
		case "clear":
			sp.Clear()
			sp1.Clear()
		default:
			sp.Push(input)
			sp1.Push(input)
			sp2.Push(input)
			sp3.Push(input)
			s.CommandPalette.AddToHistory(input)
		}
	})

	s.Start()
}
