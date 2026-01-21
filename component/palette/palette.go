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
	// triggers os.Exit() if set true when pressed
	PressEscapeToExit bool
	// key press enter can send null string
	AllowEnterNullString bool
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
	// eventHandlers it triggers when there is a change occur
	eventHandlers []func(event keyboard.KeyEvent)
	// disable listener
	baseListenerDisable bool
}

func New(cpc *Config) (*Palette, error) {
	keyEvents, err := keyboard.GetKeys(3)
	if err != nil {
		return nil, err
	}

	cpc.Prompt = ansi.Strip(cpc.Prompt)

	width, _, err := utils.GetTerminalSize()
	if err != nil {
		return nil, err
	}

	p := &Palette{
		Config:        cpc,
		keyEvents:     keyEvents,
		history:       history.New(history.DefaultCapacity),
		PromptLine:    NewLine(width - len(cpc.Prompt)),
		eventHandlers: make([]func(event keyboard.KeyEvent), 0, 2),
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

// AttachKeyEventHandler attach change handler that will trigger when command pallet has any change
// including 'typing' and other key interactions
func (p *Palette) AttachKeyEventHandler(f func(event keyboard.KeyEvent)) {
	if f == nil {
		return
	}
	p.eventHandlers = append(p.eventHandlers, f)
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

// SetBaseListener disables or enables base listener that responsible for write
// and interpret all input in to command pallet system.
func (p *Palette) SetBaseListener(enable bool) {
	p.baseListenerDisable = enable
	if enable {
		print(ansi.MakeCursorVisible)
	} else {
		print(ansi.MakeCursorInvisible)
	}
}
