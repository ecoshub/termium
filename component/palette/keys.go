package palette

import (
	"github.com/ecoshub/termium/utils/ansi"
)

type ActionCode int

const (
	KeyActionEsc        ActionCode = 0
	KeyActionEnter      ActionCode = 1
	KeyActionInnerEvent ActionCode = 2
)

type KeyAction struct {
	Action ActionCode
	Input  string
}

func (p *Palette) keyEnter() {
	cmd := p.PromptLine.String()
	if cmd == "" {
		return
	}
	p.PromptLine.Clear()
	p.runActionFunction(KeyActionEnter, cmd)
}

func (p *Palette) keySpace() {
	p.PromptLine.Append(' ')
}

func (p *Palette) keyBackspace() {
	p.PromptLine.Backspace()
}

func (p *Palette) keyEsc() {
	p.runActionFunction(KeyActionEsc, "")
}

func (p *Palette) keyDefault(char rune) {
	p.PromptLine.Append(char)
}

func (p *Palette) keyArrowUp() {
	if p.history.Len() <= 0 {
		return
	}
	up := p.history.Up()
	p.PromptLine.Set(up)
	p.runActionFunction(KeyActionInnerEvent, up)
}

func (p *Palette) keyArrowDown() {
	if p.history.Len() <= 0 {
		return
	}
	down := p.history.Down()
	p.PromptLine.Set(down)
	p.runActionFunction(KeyActionInnerEvent, down)
}

func (p *Palette) keyArrowLeft() {
	if p.PromptLine.Left() {
		print(ansi.GoLeftOneChar)
	}
}

func (p *Palette) keyArrowRight() {
	if p.PromptLine.Right() {
		print(ansi.GoRightOneChar)
	}
}
