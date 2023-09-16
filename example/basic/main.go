package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/screen"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a basic panel.
	testPanel1 := panel.NewBasicPanel(20, 5)

	// a dummy function to add values in to basic panel.
	// this "panel" variable is your handler to add remove values
	go func() {
		count := 1
		for range time.NewTicker(time.Millisecond * 250).C {
			i := rand.Intn(5)
			testPanel1.Write(i, fmt.Sprintf("this is %d. hello message and this is a long message", count))
			count++
		}
	}()

	// lets add this panel to top left corner (0,0)
	s.Add(testPanel1, &screen.ComponentConfig{PosX: 0, PosY: 0})

	// run the screen
	s.Run()

	// main thread blocker
	select {}

}
