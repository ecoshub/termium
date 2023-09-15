package screen

import (
	"fmt"
	"os"

	"github.com/ecoshub/termium/palette"
	"github.com/ecoshub/termium/utils"
)

func (s *Screen) CreateCommandPallet(conf *palette.CommandPaletteConfig) {
	cp, err := palette.New(conf)
	if err != nil {
		fmt.Println("command creation failed. error", err)
		os.Exit(1)
	}
	s.AddCommandPallet(cp)
}

func (s *Screen) AddCommandPallet(cp *palette.Command) {
	if s.CommandPalette != nil {
		fmt.Println("already has command palette")
		os.Exit(1)
	}

	s.CommandPalette = cp
	s.defaultCursorPosX = DefaultCommandPalettePositionX
	s.defaultCursorPosY = TerminalHeight - DefaultCommandPaletteHeight + len(cp.Config.PromptString) + 1
	s.buffer = utils.InitRuneMatrix(TerminalWith, TerminalHeight-DefaultCommandPaletteHeight, ' ')
}
