package screen

import (
	"github.com/ecoshub/termium/utils"
)

func (s *Screen) CreateCommandPallet(conf *CommandPaletteConfig) {
	cp, err := newCommandPalette(conf, s)
	if err != nil {
		panic("command palette creation error. err:" + err.Error())
	}
	s.attachCommandPallet(cp)
}

func (s *Screen) attachCommandPallet(cp *CommandPalette) {
	if s.CommandPalette != nil {
		panic("already has command palette")
	}

	s.CommandPalette = cp
	s.defaultCursorPosX = DefaultCommandPalettePositionX
	s.defaultCursorPosY = TerminalHeight - DefaultCommandPaletteHeight + len(cp.Config.Prompt) + 1
	s.buffer = utils.InitRuneMatrix(TerminalWith, TerminalHeight-DefaultCommandPaletteHeight, ' ')
}
