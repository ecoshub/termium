package palette

import (
	"os"

	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/internal/history"
	"github.com/ecoshub/termium/internal/line"
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

type Config struct {
	Prompt string
	Style  *style.Style
}

type Palette struct {
	Config     *Config
	PromptLine *line.Line
	history    *history.History
	keyEvents  <-chan keyboard.KeyEvent
	actionFunc func(action *KeyAction)
	hasChanged func()
}

func New(cpc *Config) (*Palette, error) {
	keyEvents, err := keyboard.GetKeys(3)
	if err != nil {
		return nil, err
	}

	cpc.Prompt = ansi.Strip(cpc.Prompt)
	p := &Palette{
		Config:     cpc,
		keyEvents:  keyEvents,
		history:    history.New(history.DefaultCapacity),
		PromptLine: line.New(utils.TerminalWith - len(cpc.Prompt)),
	}

	go p.listenKeyEvents()
	return p, nil
}

func (p *Palette) ClearHistory() {
	p.history.Clear()
}

func (p *Palette) AddToHistory(line string) {
	p.history.Add(line)
}

func (p *Palette) ChangeEvent(f func()) {
	if f == nil {
		return
	}
	p.hasChanged = f
}

func (p *Palette) ListenKeyEventEnter(f func(input string)) {
	if f == nil {
		return
	}
	p.actionFunc = func(action *KeyAction) {
		if action.Action == KeyActionEnter {
			f(action.Input)
		}
	}
}

func (p *Palette) listenActions(f func(action *KeyAction)) {
	if f == nil {
		return
	}
	p.actionFunc = f
}

func (p *Palette) listenKeyEvents() {
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
			p.hasChanged()
		}
	}
}

func (p *Palette) PromptString() string {
	return style.SetStyle(p.Config.Prompt, p.Config.Style)
}

func (p *Palette) LineString() string {
	return p.PromptLine.String()
}

func (p *Palette) getBuffer() string {
	return p.Config.Prompt + p.PromptLine.String()
}

func (p *Palette) runActionFunction(action ActionCode, input string) {
	if p.actionFunc == nil {
		return
	}
	p.actionFunc(&KeyAction{Action: action, Input: input})
}
