package screen

import (
	"time"

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

	CommandPalette *CommandPalette
}

func New(optionalCommandPaletteConfig ...*CommandPaletteConfig) (*Screen, error) {
	var cfg *CommandPaletteConfig
	if len(optionalCommandPaletteConfig) == 0 {
		cfg = &CommandPaletteConfig{
			Prompt: DefaultCommandPalettePrompt,
		}
	} else {
		cfg = optionalCommandPaletteConfig[0]
	}
	s := &Screen{
		sizeX:             TerminalWith,
		sizeY:             TerminalHeight - DefaultCommandPaletteHeight,
		components:        make([]*Component, 0, 4),
		defaultCursorPosX: DefaultCommandPalettePositionX,
		defaultCursorPosY: TerminalHeight - DefaultCommandPaletteHeight + len(cfg.Prompt) + 1,
		buffer:            utils.InitRuneMatrix(TerminalWith, TerminalHeight-DefaultCommandPaletteHeight, ' '),
	}
	cp, err := newCommandPalette(cfg, s)
	if err != nil {
		return nil, err
	}
	s.CommandPalette = cp
	return s, nil
}

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
