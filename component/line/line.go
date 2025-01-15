package line

import "github.com/ecoshub/termium/component/style"

type Line struct {
	Line  string
	Style *style.Style
}

func Empty() *Line {
	return &Line{Style: &style.Style{}}
}

func NewLines(size int, optionalStyle ...*style.Style) []*Line {
	sty := &style.Style{}
	if len(optionalStyle) > 0 {
		sty = optionalStyle[0]
	}
	lines := make([]*Line, size)
	for i := range lines {
		lines[i] = &Line{Line: "", Style: sty}
	}
	return lines
}
