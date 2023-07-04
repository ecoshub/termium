package screen

import (
	"errors"
	"time"

	"github.com/ecoshub/termium/palette"
	"github.com/ecoshub/termium/utils"
)

const (
	DefaultCommandPaletteHeight    int    = 1
	DefaultCommandPalettePositionX int    = 0
	DefaultCommandPalettePrompt    string = "  > "

	DefaultRefreshDelay time.Duration = time.Millisecond * 50
)

type Screen struct {
	ShowFPS bool

	defaultCursorPosX int
	defaultCursorPosY int
	sizeX             int
	sizeY             int
	commandPalette    *palette.Command
	components        []*Component
	buffer            [][]rune
	lastRender        time.Time
	started           bool
}

func NewScreen() (*Screen, error) {
	return &Screen{
		sizeY:             TerminalHeight,
		sizeX:             TerminalWith,
		defaultCursorPosX: 0,
		defaultCursorPosY: TerminalHeight,
		components:        make([]*Component, 0, 4),
		buffer:            utils.InitRuneMatrix(TerminalWith, TerminalHeight, ' '),
	}, nil
}

func NewDefaultScreen() (*Screen, error) {
	pc := &palette.CommandPaletteConfig{
		Width:           TerminalWith,
		PromptString:    DefaultCommandPalettePrompt,
		HistoryCapacity: palette.DefaultHistoryCapacity,
	}
	cp, err := palette.New(pc)
	if err != nil {
		return nil, err
	}
	return &Screen{
		sizeX:             TerminalWith,
		sizeY:             TerminalHeight - DefaultCommandPaletteHeight,
		components:        make([]*Component, 0, 4),
		commandPalette:    cp,
		defaultCursorPosX: DefaultCommandPalettePositionX,
		defaultCursorPosY: TerminalHeight - DefaultCommandPaletteHeight + len(pc.PromptString) + 1,
		buffer:            utils.InitRuneMatrix(TerminalWith, TerminalHeight-DefaultCommandPaletteHeight, ' '),
	}, nil
}

func (s *Screen) AttachCommandPallet(cp *palette.Command) error {
	if s.commandPalette != nil {
		return errors.New("already has command palette")
	}
	s.commandPalette = cp
	s.defaultCursorPosX = DefaultCommandPalettePositionX
	s.defaultCursorPosY = TerminalHeight - DefaultCommandPaletteHeight + len(cp.Config.PromptString) + 1
	s.buffer = utils.InitRuneMatrix(TerminalWith, TerminalHeight-DefaultCommandPaletteHeight, ' ')
	return nil
}

func (s *Screen) AttachCommandPalletFunc(f func(a *palette.Action)) {
	if f != nil {
		s.commandPalette.ListenActions(f)
	}
}

func (s *Screen) AddToCommandPalletHistory(line string) {
	s.commandPalette.AddToHistory(line)
}

func (s *Screen) ClearCommandPalletHistory() {
	s.commandPalette.ClearHistory()
}
