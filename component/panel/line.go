package panel

import "github.com/ecoshub/termium/component/style"

type Line struct {
	Line  string
	Style *style.Style
}

func NewLines(size int) []*Line {
	lines := make([]*Line, size)
	for i := range lines {
		lines[i] = &Line{Line: "", Style: &style.Style{}}
	}
	return lines
}
