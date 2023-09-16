package screen

import (
	"time"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
)

const (
	DefaultCommandPaletteHeight    int    = 1
	DefaultCommandPalettePositionX int    = 0
	DefaultCommandPalettePrompt    string = "eco$ "

	DefaultRefreshDelay time.Duration = time.Millisecond * 100
)

type Screen struct {
	ShowFPS bool

	defaultCursorPosX int
	defaultCursorPosY int
	sizeX             int
	sizeY             int
	components        []*Component
	lastRender        time.Time
	started           bool

	CommandPalette *palette.CommandPalette
}

func New(optionalCommandPaletteConfig ...*palette.CommandPaletteConfig) (*Screen, error) {
	var cfg *palette.CommandPaletteConfig
	if len(optionalCommandPaletteConfig) == 0 {
		cfg = &palette.CommandPaletteConfig{
			Prompt: DefaultCommandPalettePrompt,
			Style:  &style.Style{},
		}
	} else {
		cfg = optionalCommandPaletteConfig[0]
	}
	s := &Screen{
		sizeX:             utils.TerminalWith,
		sizeY:             utils.TerminalHeight - DefaultCommandPaletteHeight,
		components:        make([]*Component, 0, 4),
		defaultCursorPosX: DefaultCommandPalettePositionX,
		defaultCursorPosY: utils.TerminalHeight - DefaultCommandPaletteHeight + len(cfg.Prompt) + 1,
	}
	cp, err := palette.New(cfg)
	if err != nil {
		return nil, err
	}
	s.CommandPalette = cp
	s.CommandPalette.ChangeEvent(func() { s.RenderCommandPalette() })
	return s, nil
}
