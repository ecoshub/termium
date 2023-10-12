package palette

import (
	"github.com/ecoshub/termium/component/history"
	"github.com/ecoshub/termium/component/style"
	"github.com/ecoshub/termium/utils"
	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

type Config struct {
	// Prompt prompt
	Prompt string
	// Style prompt style
	Style *style.Style
}

type Palette struct {
	// Config base palette configuration
	Config *Config
	// PromptLine is a line implementation that can control input and cursor
	PromptLine *Line
	// history history for old inputs
	history *history.History
	// keyEvents key event channel
	keyEvents <-chan keyboard.KeyEvent
	// eventHandler event handler function
	eventHandler func(eventCode EventCode, input string)
	// changeHandler it triggers when there is a change occur
	changeHandler func()
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
		PromptLine: NewLine(utils.TerminalWith - len(cpc.Prompt)),
	}

	go p.listenKeyEvents()
	return p, nil
}

// ClearHistory clear command palette history
func (p *Palette) ClearHistory() {
	p.history.Clear()
}

// AddToHistory add to command line history to access it with up down arrow keys later
func (p *Palette) AddToHistory(line string) {
	p.history.Add(line)
}

// AttachChangeHandler attach change handler that will trigger when command pallet has any change
// including 'typing' and other key interactions
func (p *Palette) AttachChangeHandler(f func()) {
	if f == nil {
		return
	}
	p.changeHandler = f
}

// ListenKeyEventEnter listen for key event 'Enter'
// it is the main handler for command input
// 'input' can not be null string
func (p *Palette) ListenKeyEventEnter(f func(input string)) {
	if f == nil {
		return
	}
	p.eventHandler = func(eventCode EventCode, input string) {
		if eventCode == EnterKeyPressed {
			f(input)
		}
	}
}

// Prompt get prompt string with style
func (p *Palette) Prompt() string {
	return style.SetStyle(p.Config.Prompt, p.Config.Style)
}

// Input get input line
func (p *Palette) Input() string {
	return p.PromptLine.String()
}
