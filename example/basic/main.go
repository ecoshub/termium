package main

import (
	"fmt"
	"time"

	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/screen"
)

func main() {
	// create a screen. this is representation of terminal screen
	s, err := screen.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	// lets create a basic panel.
	testPanel1 := panel.NewBasicPanel(&panel.Config{
		Width:  100,
		Height: 10,
	})

	// a dummy function to add values in to basic panel.
	// this "panel" variable is your handler to add remove values
	go func() {
		count := 1
		for range time.NewTicker(time.Millisecond * 250).C {
			i := count % 10
			testPanel1.Write(i, fmt.Sprintf("this is %d. hello message.......................", count))
			count++
		}
	}()

	// lets add this panel to top left corner (0,0)
	s.Add(testPanel1, 0, 0)

	// run the screen
	s.Start()

}
