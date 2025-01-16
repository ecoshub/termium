package screen

import (
	"sync"

	"github.com/ecoshub/termium/component/palette"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

type Config struct {
	DisableCommentPallet bool
	CommandPaletteConfig *palette.Config
}
type Renderer struct {
	sync.Mutex
	components          []*Component
	commandPalette      *palette.Palette
	renderCommandPallet bool
}

type Screen struct {
	Config         *Config
	CommandPalette *palette.Palette
	lineBuffer     string
	renderer       *Renderer
	started        bool
}

func New(optionalConfig ...*Config) (*Screen, error) {
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
		renderer: &Renderer{
			components:          make([]*Component, 0, 2),
			commandPalette:      cp,
			renderCommandPallet: !cfg.DisableCommentPallet,
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
	return selectedConfig, nil
}

func (s *Screen) ResetScreen() {
	print(ansi.SaveCursorPos)
	defer print(ansi.RestoreCursorPos)

	print(ansi.GoToFirstBlock)
	s.renderer.Render()
	s.renderer.RenderCommandPalette()
}
