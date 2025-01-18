package screen

import (
	"sync"
	"time"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

const (
	DefaultFPSLimit int = 60
)

type Config struct {
	DisableCommentPallet bool
	CommandPaletteConfig *palette.Config
	FPSLimit             int
}

type Renderer struct {
	sync.Mutex
	terminalWidth          int
	terminalHeight         int
	components             []*Component
	componentRendered      map[int]bool
	componentTitleRenderer map[int]bool
	commandPalette         *palette.Palette
	renderCommandPallet    bool
	lastRender             time.Time
	maxRenderTimeGap       time.Duration
}

type Screen struct {
	Config         *Config
	CommandPalette *palette.Palette
	TerminalWidth  int
	TerminalHeight int
	lineBuffer     string
	renderer       *Renderer
	started        bool
}

func New(optionalConfig ...*Config) (*Screen, error) {
	width, height, err := utils.GetTerminalSize()
	if err != nil {
		return nil, err
	}
	cfg, err := resolveConfig(optionalConfig)
	if err != nil {
		return nil, err
	}
	cp, err := palette.New(cfg.CommandPaletteConfig)
	if err != nil {
		return nil, err
	}
	s := &Screen{
		Config:         cfg,
		CommandPalette: cp,
		TerminalWidth:  width,
		TerminalHeight: height,
		renderer: &Renderer{
			terminalWidth:          width,
			terminalHeight:         height,
			components:             make([]*Component, 0, 2),
			commandPalette:         cp,
			renderCommandPallet:    !cfg.DisableCommentPallet,
			maxRenderTimeGap:       time.Second / time.Duration(cfg.FPSLimit),
			componentRendered:      make(map[int]bool),
			componentTitleRenderer: make(map[int]bool),
		},
	}
	s.CommandPalette.AttachKeyEventHandler(func(event keyboard.KeyEvent) { s.renderer.RenderCommandPalette() })
	return s, nil
}

func resolveConfig(optionalConfig []*Config) (*Config, error) {
	if len(optionalConfig) == 0 {
		// return default config
		return &Config{
			CommandPaletteConfig: &palette.Config{
				Prompt: DefaultCommandPalettePrompt,
				Style:  &style.Style{},
			},
		}, nil
	}
	// validate and modify selected config (first config)
	selectedConfig := optionalConfig[0]
	if selectedConfig.CommandPaletteConfig == nil {
		selectedConfig.CommandPaletteConfig = &palette.Config{
			Prompt: DefaultCommandPalettePrompt,
			Style:  &style.Style{},
		}
	}
	if selectedConfig.CommandPaletteConfig.Style == nil {
		selectedConfig.CommandPaletteConfig.Style = &style.Style{}
	}
	if selectedConfig.FPSLimit == 0 {
		selectedConfig.FPSLimit = DefaultFPSLimit
	}
	return selectedConfig, nil
}

func (s *Screen) ResetScreen() {
	print(ansi.SaveCursorPos)
	defer print(ansi.RestoreCursorPos)

	print(ansi.GoToFirstBlock)
	s.renderer.Render()
	s.renderer.RenderCommandPalette()
}
