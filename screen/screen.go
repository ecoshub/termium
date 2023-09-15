package screen

import (
	"time"

	"github.com/ecoshub/termium/palette"
	"github.com/ecoshub/termium/utils"
)

const (
	DefaultCommandPaletteHeight    int    = 1
	DefaultCommandPalettePositionX int    = 0
	DefaultCommandPalettePrompt    string = "  > "

	DefaultRefreshDelay time.Duration = time.Millisecond * 100
)

type Screen struct {
	ShowFPS bool

	defaultCursorPosX int
	defaultCursorPosY int
	sizeX             int
	sizeY             int
	components        []*Component
	buffer            [][]rune
	lastRender        time.Time
	started           bool

	CommandPalette *palette.Command
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
		CommandPalette:    cp,
		defaultCursorPosX: DefaultCommandPalettePositionX,
		defaultCursorPosY: TerminalHeight - DefaultCommandPaletteHeight + len(pc.PromptString) + 1,
		buffer:            utils.InitRuneMatrix(TerminalWith, TerminalHeight-DefaultCommandPaletteHeight, ' '),
	}, nil
}
