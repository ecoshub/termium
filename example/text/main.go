package main

import (
	"fmt"
	"strings"
	"time"

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
	tp := panel.NewASCIIPanel(&panel.Config{
		Width:       100,
		Height:      10,
		Title:       "ASCII Panel",
		RenderTitle: true,
		TitleStyle: &style.Style{
			BackgroundColor: 60,
		},
		ContentStyle: &style.Style{
			ForegroundColor: 195,
		},
	})

	infoPanel := panel.NewStackPanel(&panel.Config{
		Width:        utils.TerminalWith,
		Height:       1,
		ContentStyle: style.DefaultStyleError,
	})

	// lets add this panel to top left corner (0,0)
	s.Add(tp, 0, 0)
	s.Add(infoPanel, 0, utils.TerminalHeight-2)

	go func() {
		for range time.NewTicker(time.Second * 10).C {
			infoPanel.Clear()
		}
	}()

	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {

		if strings.HasPrefix(input, ":dump") {
			tokens := strings.Split(input, " ")
			if len(tokens) != 2 {
				infoPanel.Push("unsupported number of argument for 'dump' command", style.DefaultStyleError)
				return
			}
			path := tokens[1]
			n, err := tp.Dump(path)
			if err != nil {
				infoPanel.Push("unsupported number of argument for 'dump' command", style.DefaultStyleError)
				return
			}
			infoPanel.Push(fmt.Sprintf("file dump success. %d bytes written. path: %s", n, path), style.DefaultStyleSuccess)
			return
		}

		if strings.HasPrefix(input, ":flush") {
			tokens := strings.Split(input, " ")
			if len(tokens) != 1 {
				infoPanel.Push("unsupported number of argument for 'dump' command", style.DefaultStyleError)
				return
			}
			tp.Flush()
			return
		}

		tp.Appendln(input)
		s.CommandPalette.AddToHistory(input)

	})

	s.Start()
}
