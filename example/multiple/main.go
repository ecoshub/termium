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

	sp := panel.NewStackPanel(10, 10)
	s.Add(sp, &screen.ComponentConfig{})

	sp1 := panel.NewStackPanel(10, 10)
	s.Add(sp1, &screen.ComponentConfig{PosX: 11})

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
		}
	})

	s.Start()
}
