package palette

import (
	"os"

	"github.com/ecoshub/termium/utils"
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
	defer p.hasChanged()

	if p.bufferSize == 0 {
		return
	}

	cmd := string(p.buffer[p.plp : p.plp+p.bufferSize])
	p.buffer = utils.FixedSizeLine(cmd, utils.TerminalWith)
	p.cursorIndex = p.plp
	p.bufferSize = 0

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
	defer p.hasChanged()
	if p.cursorIndex >= utils.TerminalWith {
		return
	}
	p.buffer[p.cursorIndex] = ' '
	p.cursorIndex++
	p.bufferSize++
}

func (p *CommandPalette) keyBackspace() {
	if p.bufferSize <= 0 {
		return
	}
	p.buffer[p.cursorIndex-1] = ' '
	p.cursorIndex--
	p.bufferSize--
	ansi.GoLeftNChar(1)
	p.hasChanged()
}

func (p *CommandPalette) keyEsc() {
	defer p.hasChanged()
	p.runActionFunction(KeyActionEsc, "")
}

func (p *CommandPalette) keyDefault(char rune) {
	defer p.hasChanged()
	if p.cursorIndex >= utils.TerminalWith {
		return
	}
	p.buffer[p.cursorIndex] = char
	p.cursorIndex++
	p.bufferSize++
}

func (p *CommandPalette) keyArrowUp() {
	if p.history.Len() <= 0 {
		return
	}
	defer p.hasChanged()
	up := p.history.Up()
	plu := utils.PrintableLen(up)
	copy(p.buffer[p.plp:p.plp+plu], []rune(up)[:])
	p.cursorIndex = p.plp + plu
	p.bufferSize = plu
	p.runActionFunction(KeyActionInnerEvent, up)
}

func (p *CommandPalette) keyArrowDown() {
	if p.history.Len() <= 0 {
		return
	}

	defer p.hasChanged()
	down := p.history.Down()
	pld := utils.PrintableLen(down)
	copy(p.buffer[p.plp:p.plp+pld], []rune(down)[:])
	p.cursorIndex = p.plp + pld
	p.bufferSize = pld
	p.runActionFunction(KeyActionInnerEvent, down)
}

func (p *CommandPalette) keyArrowLeft() {
	if p.cursorIndex <= p.plp {
		return
	}
	p.hasChanged()
	p.cursorIndex--
	ansi.GoLeftNChar(1)
}

func (p *CommandPalette) keyArrowRight() {
	if p.cursorIndex >= p.bufferSize+p.plp {
		return
	}
	ansi.GoRightNChar(1)
	p.cursorIndex++
}
