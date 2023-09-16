package screen

import (
	"os"

	"github.com/ecoshub/termium/history"
	"github.com/ecoshub/termium/utils"
	"github.com/eiannone/keyboard"
)

var (
	DefaultHistoryCapacity = 10
)

type CommandPaletteConfig struct {
	Prompt string
	Enable bool
}

type CommandPalette struct {
	Config *CommandPaletteConfig

	cursorIndex int
	buffer      []rune
	bufferSize  int
	history     *history.History
	keyEvents   <-chan keyboard.KeyEvent
	actionFunc  func(action *KeyAction)
	renderer    *Screen
	plp         int
}

func newCommandPalette(cpc *CommandPaletteConfig, renderer *Screen) (*CommandPalette, error) {
	keyEvents, err := keyboard.GetKeys(3)
	if err != nil {
		return nil, err
	}

	plp := utils.PrintableLen(cpc.Prompt)

	p := &CommandPalette{
		Config:      cpc,
		keyEvents:   keyEvents,
		history:     history.New(DefaultHistoryCapacity),
		buffer:      utils.InitRuneArray(TerminalWith, ' '),
		cursorIndex: plp,
		renderer:    renderer,
		plp:         plp,
	}

	go p.listenKeyEvents()
	return p, nil
}

func (p *CommandPalette) ListenActions(f func(action *KeyAction)) {
	if f == nil {
		return
	}
	p.actionFunc = f
}

func (p *CommandPalette) ClearHistory() {
	p.history.Clear()
}

func (p *CommandPalette) AddToHistory(line string) {
	p.history.Add(line)
}

func (p *CommandPalette) SetEnable(enable bool) {
	p.Config.Enable = enable
}

func (p *CommandPalette) Enable() {
	p.Config.Enable = true
}

func (p *CommandPalette) Disable() {
	p.Config.Enable = false
}

func (p *CommandPalette) Buffer() string {
	plen := utils.PrintableLen(p.Config.Prompt)
	if len(p.buffer) == 0 {
		return p.Config.Prompt
	}
	return p.Config.Prompt + string(p.buffer[plen:plen+p.bufferSize])
}

func (p *CommandPalette) listenKeyEvents() {
	for {
		select {
		case event := <-p.keyEvents:
			if !p.Config.Enable {
				if event.Key == keyboard.KeyCtrlC {
					os.Exit(0)
				}
				continue
			}
			switch event.Key {
			case keyboard.KeyEnter:
				p.keyEnter()
			case keyboard.KeyArrowUp:
				p.keyArrowUp()
			case keyboard.KeyArrowDown:
				p.keyArrowDown()
			case keyboard.KeyArrowLeft:
				p.keyArrowLeft()
			case keyboard.KeyArrowRight:
				p.keyArrowRight()
			case keyboard.KeySpace:
				p.keySpace()
			case keyboard.KeyBackspace, keyboard.KeyBackspace2:
				p.keyBackspace()
			case keyboard.KeyEsc:
				p.keyEsc()
			case keyboard.KeyCtrlC:
				os.Exit(0)
			default:
				p.keyDefault(event.Rune)
			}
		}
	}
}

func (p *CommandPalette) runActionFunction(action ActionCode, input string) {
	if p.actionFunc == nil {
		return
	}
	p.actionFunc(&KeyAction{Action: action, Input: input})
}
