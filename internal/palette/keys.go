package palette

import (
	"os"
	"term/internal/ansi"
)

const (
	CommandPaletteCommandQuit1        = "q"
	CommandPaletteCommandQuit2        = "quit"
	CommandPaletteCommandQuit3        = "exit"
	CommandPaletteCommandClearHistory = "clr"
)

func (p *Command) keyEnter() {
	if len(p.buffer) == 0 {
		return
	}
	cmd := string(p.buffer)
	p.buffer = make([]rune, 0, p.Config.Width)
	p.cursorIndex = len(p.Config.PromptString)

	switch cmd {
	case CommandPaletteCommandQuit1,
		CommandPaletteCommandQuit2,
		CommandPaletteCommandQuit3:
		os.Exit(0)
	case CommandPaletteCommandClearHistory:
		p.ClearHistory()
	default:
		p.runActionFunction(ActionEnter, cmd)
	}
}

func (p *Command) keySpace() {
	if p.cursorIndex >= p.Config.Width {
		return
	}
	print(" ")
	p.buffer = append(p.buffer, ' ')
	p.cursorIndex++
}

func (p *Command) keyBackspace() {
	if p.cursorIndex <= len(p.Config.PromptString) {
		return
	}
	ansi.GoLeft(1)
	print(" ")
	ansi.GoLeft(1)

	p.buffer = p.buffer[:len(p.buffer)-1]

	p.cursorIndex--
}

func (p *Command) keyEsc() {
	p.runActionFunction(ActionEsc, "")
}

func (p *Command) keyDefault(char rune) {
	if p.cursorIndex >= p.Config.Width {
		return
	}
	print(string(char))
	p.buffer = append(p.buffer, char)
	p.cursorIndex++
}

func (p *Command) keyArrowUp() {
	up := p.history.Up()
	print(up)
	p.runActionFunction(ActionInnerEvent, up)
	p.cursorIndex = len(p.Config.PromptString) + len(up)
	p.buffer = []rune(up)
}

func (p *Command) keyArrowDown() {
	down := p.history.Down()
	print(down)
	p.runActionFunction(ActionInnerEvent, down)
	p.cursorIndex = len(p.Config.PromptString) + len(down)
	p.buffer = []rune(down)
}

func (p *Command) keyArrowLeft()  {}
func (p *Command) keyArrowRight() {}
