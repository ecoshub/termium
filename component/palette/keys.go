package palette

import (
	"os"

	"github.com/ecoshub/termium/utils/ansi"
)

const (
	CommandPaletteCommandQuit1 = "q"
	CommandPaletteCommandQuit2 = "quit"
	CommandPaletteCommandQuit3 = "exit"
)

type ActionCode int

const (
	KeyActionEsc        ActionCode = 0xf0
	KeyActionEnter      ActionCode = 0xf1
	KeyActionInnerEvent ActionCode = 0xf2
)

type KeyAction struct {
	Action ActionCode
	Input  string
}

func newAction(action ActionCode, input string) *KeyAction {
	return &KeyAction{Action: action, Input: input}
}

func (p *CommandPalette) keyEnter() {
	cmd := p.PromptLine.String()
	if cmd == "" {
		return
	}

	p.PromptLine.Clear()

	switch cmd {
	case CommandPaletteCommandQuit1,
		CommandPaletteCommandQuit2,
		CommandPaletteCommandQuit3:
		os.Exit(0)
	default:
		p.runActionFunction(KeyActionEnter, cmd)
	}
}

func (p *CommandPalette) keySpace() {
	p.PromptLine.Append(' ')
}

func (p *CommandPalette) keyBackspace() {
	p.PromptLine.Backspace()
}

func (p *CommandPalette) keyEsc() {
	p.runActionFunction(KeyActionEsc, "")
}

func (p *CommandPalette) keyDefault(char rune) {
	p.PromptLine.Append(char)
}

func (p *CommandPalette) keyArrowUp() {
	if p.history.Len() <= 0 {
		return
	}
	up := p.history.Up()
	p.PromptLine.Set(up)
	p.runActionFunction(KeyActionInnerEvent, up)
}

func (p *CommandPalette) keyArrowDown() {
	if p.history.Len() <= 0 {
		return
	}
	down := p.history.Down()
	p.PromptLine.Set(down)
	p.runActionFunction(KeyActionInnerEvent, down)
}

func (p *CommandPalette) keyArrowLeft() {
	if p.PromptLine.Left() {
		print(ansi.GoLeftOneChar)
	}
}

func (p *CommandPalette) keyArrowRight() {
	if p.PromptLine.Right() {
		print(ansi.GoRightOneChar)
	}
}
