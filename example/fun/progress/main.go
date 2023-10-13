package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/progress"
	"github.com/ecoshub/termium/component/screen"
	"github.com/ecoshub/termium/component/style"
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

	bar1, err := progress.New(&progress.Config{Width: 50, BarStyle: progress.DefaultBarStyle})
	if err != nil {
		panic(err)
	}

	bar2, err := progress.New(&progress.Config{Width: 50, BarStyle: []rune{'◉', '◉', '◉', '○', '◉'}})
	if err != nil {
		panic(err)
	}

	s.Add(bar1, 0, 0)
	s.Add(bar2, 0, 1)

	updatePeriodically(bar1, 0.1, time.Millisecond*200)
	updatePeriodically(bar2, 0.05, time.Millisecond*100)

	// command handler
	s.CommandPalette.ListenKeyEventEnter(func(input string) {

		// also add command pallets own history module to select with up/down arrow keys later
		s.CommandPalette.AddToHistory(input)

		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			bar1.Update(num)
		}
	})

	s.Start()
}

func updatePeriodically(bar *progress.ProgressBar, delta float64, delay time.Duration) {
	tot := 0.0
	go func() {
		for range time.NewTicker(delay).C {
			tot += delta
			if tot >= 1 {
				tot = 0.0
			}
			bar.Update(tot)
		}
	}()
}
