package screen

import (
	"os"

	"github.com/ecoshub/termium/ansi"
	"github.com/ecoshub/termium/panel"
	"github.com/ecoshub/termium/utils"
)

const (
	CommandPaletteCommandQuit1        = "q"
	CommandPaletteCommandQuit2        = "quit"
	CommandPaletteCommandQuit3        = "exit"
	CommandPaletteCommandClearHistory = "clr"
)

func (p *CommandPalette) keyEnter() {
	defer p.renderer.Render()

	if p.bufferSize == 0 {
		return
	}

	cmd := string(p.buffer[p.plp : p.plp+p.bufferSize])
	p.buffer = panel.FixedSizeLine(cmd, TerminalWith)
	p.cursorIndex = p.plp
	p.bufferSize = 0

	switch cmd {
	case CommandPaletteCommandQuit1,
		CommandPaletteCommandQuit2,
		CommandPaletteCommandQuit3:
		os.Exit(0)
	case CommandPaletteCommandClearHistory:
		p.ClearHistory()
	default:
		p.runActionFunction(KeyActionEnter, cmd)
	}
}

func (p *CommandPalette) keySpace() {
	defer p.renderer.Render()
	if p.cursorIndex >= TerminalWith {
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
	ansi.GoLeft(1)
	p.renderer.Render()
}

func (p *CommandPalette) keyEsc() {
	defer p.renderer.Render()
	p.runActionFunction(KeyActionEsc, "")
}

func (p *CommandPalette) keyDefault(char rune) {
	defer p.renderer.Render()
	if p.cursorIndex >= TerminalWith {
		return
	}
	p.buffer[p.cursorIndex] = char
	p.cursorIndex++
	p.bufferSize++
}

func (p *CommandPalette) keyArrowUp() {
	defer p.renderer.Render()
	up := p.history.Up()
	plu := utils.PrintableLen(up)
	copy(p.buffer[p.plp:p.plp+plu], []rune(up)[:])
	p.cursorIndex = plu
	p.bufferSize = plu
	p.runActionFunction(KeyActionInnerEvent, up)
}

func (p *CommandPalette) keyArrowDown() {
	defer p.renderer.Render()
	down := p.history.Down()
	pld := utils.PrintableLen(down)
	copy(p.buffer[p.plp:p.plp+pld], []rune(down)[:])
	p.cursorIndex = pld
	p.bufferSize = pld
	p.runActionFunction(KeyActionInnerEvent, down)
}

func (p *CommandPalette) keyArrowLeft() {
	if p.cursorIndex <= p.plp {
		return
	}
	p.renderer.Render()
	p.cursorIndex--
	ansi.GoLeft(1)
}

func (p *CommandPalette) keyArrowRight() {
	if p.cursorIndex >= p.bufferSize+p.plp {
		return
	}
	ansi.GoRight(1)
	p.cursorIndex++
}
