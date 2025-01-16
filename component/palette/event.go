package palette

import (
	"os"

	"github.com/ecoshub/termium/utils/ansi"
	"github.com/eiannone/keyboard"
)

// EventCode is a event indicator for special keys presses.
type EventCode int

const (
	EnterKeyPressed    EventCode = 0
	InnerKeyPressEvent EventCode = 1
)

func (p *Palette) keyPressHandlerEnter() {
	cmd := p.PromptLine.String()
	if cmd == "" {
		// enter key press event can only sent if input is not null
		return
	}
	p.PromptLine.Clear()
	p.runEvent(EnterKeyPressed, cmd)
}

func (p *Palette) keyPressHandlerSpace() {
	p.PromptLine.Append(' ')
}

func (p *Palette) keyPressHandlerBackspace() {
	p.PromptLine.Backspace()
}

func (p *Palette) keyPressHandlerDefaultKeys(char rune) {
	p.PromptLine.Append(char)
}

func (p *Palette) keyPressHandlerArrowUp() {
	if p.history.Len() <= 0 {
		return
	}
	up := p.history.Up()
	p.PromptLine.Set(up)
	p.runEvent(InnerKeyPressEvent, up)
}

func (p *Palette) keyPressHandlerArrowDown() {
	if p.history.Len() <= 0 {
		return
	}
	down := p.history.Down()
	p.PromptLine.Set(down)
	p.runEvent(InnerKeyPressEvent, down)
}

func (p *Palette) keyPressHandlerArrowLeft() {
	if p.PromptLine.Left() {
		print(ansi.GoLeftOneChar)
	}
}

func (p *Palette) keyPressHandlerArrowRight() {
	if p.PromptLine.Right() {
		print(ansi.GoRightOneChar)
	}
}

func (p *Palette) keyPressHandlerCTRL_A() {
	p.PromptLine.GotoStart()
}

func (p *Palette) keyPressHandlerCTRL_E() {
	p.PromptLine.GotoEnd()
}

func (p *Palette) listenKeyEvents() {
	for event := range p.keyEvents {
		if p.baseListenerDisable {
			p.triggerEventHandlers(event)
			continue
		}
		switch event.Key {
		case keyboard.KeyEnter:
			p.keyPressHandlerEnter()
		case keyboard.KeyArrowUp:
			p.keyPressHandlerArrowUp()
		case keyboard.KeyArrowDown:
			p.keyPressHandlerArrowDown()
		case keyboard.KeyArrowLeft:
			p.keyPressHandlerArrowLeft()
		case keyboard.KeyArrowRight:
			p.keyPressHandlerArrowRight()
		case keyboard.KeySpace:
			p.keyPressHandlerSpace()
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			p.keyPressHandlerBackspace()
		case keyboard.KeyEsc:
			if p.Config.PressEscapeToExit {
				ResetTerm()
				os.Exit(0)
			}
		case keyboard.KeyCtrlC:
			ResetTerm()
			os.Exit(0)
		case keyboard.KeyCtrlA:
			p.keyPressHandlerCTRL_A()
		case keyboard.KeyCtrlE:
			p.keyPressHandlerCTRL_E()
		default:
			p.keyPressHandlerDefaultKeys(event.Rune)
		}
		p.triggerEventHandlers(event)
	}
}

func (p *Palette) runEvent(eventCode EventCode, input string) {
	if p.eventHandler == nil {
		return
	}
	p.eventHandler(eventCode, input)
}

func (p *Palette) triggerEventHandlers(event keyboard.KeyEvent) {
	for _, f := range p.eventHandlers {
		if f == nil {
			continue
		}
		f(event)
	}
}

func ResetTerm() {
	print(ansi.ResetAllModes)
	print(ansi.MakeCursorVisible)
}
