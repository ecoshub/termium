package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/example/tcp/core"
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
	commPanel := panel.NewStackPanel(&config.Config{
		Width:  utils.TerminalWith,
		Height: utils.TerminalHeight - 1,
		TitleStyle: &style.Style{
			BackgroundColor: 60,
		},
		ContentStyle: &style.Style{
			ForegroundColor: 195,
		},
	})

	infoPanel := panel.NewStackPanel(&config.Config{
		Width:        utils.TerminalWith,
		Height:       1,
		ContentStyle: style.DefaultStyleError,
	})

	// lets add this panel to top left corner (0,0)
	s.Add(commPanel, 0, 0)
	s.Add(infoPanel, 0, utils.TerminalHeight-2)

	go func() {
		for range time.NewTicker(time.Second * 10).C {
			infoPanel.Clear()
		}
	}()

	go core.StartListen(":9090", infoPanel)

	var client net.Conn
	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {
		s.CommandPalette.AddToHistory(input)

		args := strings.Split(input, " ")
		command := args[0]
		if len(args) > 1 {
			args = args[1:]
		}
		switch command {
		case "connect":
			client, err = net.Dial("tcp", args[0])
			if err != nil {
				infoPanel.Push(err.Error())
				return
			}
			if client != nil {
				go core.ReadClient(client, commPanel, infoPanel)
			}
			return
		}

		if client == nil {
			infoPanel.Push("no connection establish")
			return
		}

		commPanel.Push(fmt.Sprintf("<< %s", input), style.DefaultStyleEvent)
		_, err = client.Write([]byte(input))
		if err != nil {
			infoPanel.Push(err.Error())
			return
		}
	})

	s.Start()
}
