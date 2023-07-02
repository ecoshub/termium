package palette

import (
	"os"

	"github.com/ecoshub/termium/history"
	"github.com/eiannone/keyboard"
)

var (
	DefaultHistoryCapacity = 5
)

type CommandPaletteConfig struct {
	Width           int
	PromptString    string
	HistoryCapacity int
}

type Command struct {
	Config *CommandPaletteConfig

	cursorIndex int
	buffer      []rune
	history     *history.History
	keyEvents   <-chan keyboard.KeyEvent
	actionFunc  func(action *Action)
}

func New(cpc *CommandPaletteConfig) (*Command, error) {
	keyEvents, err := keyboard.GetKeys(3)
	if err != nil {
		return nil, err
	}

	p := &Command{
		Config:      cpc,
		keyEvents:   keyEvents,
		history:     history.New(DefaultHistoryCapacity),
		buffer:      make([]rune, 0, cpc.Width),
		cursorIndex: len(cpc.PromptString),
	}

	go p.listenKeyEvents()
	return p, nil
}

func (p *Command) ListenActions(f func(action *Action)) {
	if f != nil {
		p.actionFunc = f
	}
}

func (p *Command) ClearHistory() {
	p.history.Clear()
}

func (p *Command) AddToHistory(line string) {
	p.history.Add(line)
}

func (p *Command) listenKeyEvents() {
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
				p.keyEsc()
			case keyboard.KeyCtrlC:
				os.Exit(0)
			default:
				p.keyDefault(event.Rune)
			}
		}
	}
}

func (p *Command) Buffer() string {
	return p.Config.PromptString + string(p.buffer)
}

func (p *Command) runActionFunction(action ActionCode, input string) {
	if p.actionFunc == nil {
		return
	}
	p.actionFunc(&Action{Action: action, Input: input})
}
