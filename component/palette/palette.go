package palette

import (
	"os"

	"github.com/ecoshub/termium/component/history"
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

var (
	DefaultHistoryCapacity = 10
)

type CommandPaletteConfig struct {
	Prompt          string
	ForegroundColor int
	BackgroundColor int
}

type CommandPalette struct {
	Config *CommandPaletteConfig

	cursorIndex int
	buffer      []rune
	bufferSize  int
	history     *history.History
	keyEvents   <-chan keyboard.KeyEvent
	actionFunc  func(action *KeyAction)
	hasChanged  func()
	plp         int
}

func New(cpc *CommandPaletteConfig) (*CommandPalette, error) {
	keyEvents, err := keyboard.GetKeys(3)
	if err != nil {
		return nil, err
	}

	cpc.Prompt = ansi.Strip(cpc.Prompt)
	p := &CommandPalette{
		Config:      cpc,
		keyEvents:   keyEvents,
		history:     history.New(DefaultHistoryCapacity),
		buffer:      utils.InitRuneArray(utils.TerminalWith, ' '),
		cursorIndex: len(cpc.Prompt),
		plp:         len(cpc.Prompt),
	}

	go p.listenKeyEvents()
	return p, nil
}

func (p *CommandPalette) ClearHistory() {
	p.history.Clear()
}

func (p *CommandPalette) AddToHistory(line string) {
	p.history.Add(line)
}

func (p *CommandPalette) ChangeEvent(f func()) {
	if f == nil {
		return
	}
	p.hasChanged = f
}

func (p *CommandPalette) ListenKeyEventEnter(f func(input string)) {
	if f == nil {
		return
	}
	p.actionFunc = func(action *KeyAction) {
		if action.Action == KeyActionEnter {
			f(action.Input)
		}
	}
}

func (p *CommandPalette) listenActions(f func(action *KeyAction)) {
	if f == nil {
		return
	}
	p.actionFunc = f
}

func (p *CommandPalette) listenKeyEvents() {
	for {
		select {
		case event := <-p.keyEvents:
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
				os.Exit(0)
			case keyboard.KeyCtrlC:
				os.Exit(0)
			default:
				p.keyDefault(event.Rune)
			}
		}
	}
}

func (p *CommandPalette) GetCursorIndex() int {
	return p.cursorIndex
}

func (p *CommandPalette) String() string {
	s := ""
	prompt := ansi.SetColor(p.Config.Prompt, p.Config.ForegroundColor, p.Config.BackgroundColor)
	s += prompt
	if len(p.buffer) != 0 {
		s += string(p.buffer[p.plp : p.plp+p.bufferSize])
	}
	return s
}

func (p *CommandPalette) getBuffer() string {
	plen := utils.PrintableLen(p.Config.Prompt)
	if len(p.buffer) == 0 {
		return p.Config.Prompt
	}
	return p.Config.Prompt + string(p.buffer[plen:plen+p.bufferSize])
}

func (p *CommandPalette) runActionFunction(action ActionCode, input string) {
	if p.actionFunc == nil {
		return
	}
	p.actionFunc(&KeyAction{Action: action, Input: input})
}
