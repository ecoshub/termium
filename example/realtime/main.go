package realtime

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/screen"
)

var (
	mainScreen *screen.Screen
	results    *panel.Basic
)

func main() {

	length := 100
	list := make([]int, 0, length)

	// lets say we have an array that can constantly change and
	go func() {
		for {
			index := rand.Intn(length)
			r := rand.Intn(1000)
			list[index] = r
		}
	}()

}

func initScreen() {
	var err error
	mainScreen, err = screen.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// lets create a stack panel to use as a command history
	results = panel.NewBasicPanel(screen.TerminalWith, 5)

	// creating  command pallet
	mainScreen.CreateCommandPallet(&screen.CommandPaletteConfig{
		Prompt: "ecoshub$ ",
	})

	// lets add this panel to top left corner (0,0)
	mainScreen.Add(results, &screen.ComponentConfig{
		Title:       "History:",
		RenderTitle: true,
		PosX:        0,
		// 7 is panel height (5) + terminal height(1) + history panel title(1)
		PosY: screen.TerminalHeight - 7,
	})
}

func someFunction(s *screen.Screen) {}
