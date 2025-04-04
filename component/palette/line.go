package palette

type Line struct {
	buffer []rune
	cap    int
	index  int
	width  int
	dirty  bool
}

func NewLine(width int) *Line {
	return &Line{
		width:  width,
		buffer: make([]rune, 0, 8),
		dirty:  true,
	}
}

func (l *Line) Backspace() {
	if l.index == 0 {
		return
	}
	if l.index == l.cap {
		l.buffer = l.buffer[:l.index-1]
	} else {
		start := l.buffer[:l.index-1]
		end := l.buffer[l.index:]
		l.buffer = make([]rune, 0, len(l.buffer)-1)
		l.buffer = append(l.buffer, start...)
		l.buffer = append(l.buffer, end...)
	}
	l.cap--
	l.index--
	l.dirty = true
}

func (l *Line) Append(char rune) {
	if char == 0 {
		return
	}
	if l.cap >= l.width {
		return
	}
	if l.index == l.cap {
		l.buffer = append(l.buffer, char)
	} else {
		start := l.buffer[:l.index]
		end := l.buffer[l.index:]
		l.buffer = make([]rune, 0, len(l.buffer)+1)
		l.buffer = append(l.buffer, start...)
		l.buffer = append(l.buffer, char)
		l.buffer = append(l.buffer, end...)
	}
	l.cap++
	l.index++
	l.dirty = true
}

func (l *Line) Left() bool {
	if l.index == 0 {
		return false
	}
	l.index--
	return true
}

func (l *Line) Right() bool {
	if l.index == l.cap {
		return false
	}
	l.index++
	return true
}

func (l *Line) Set(text string) {
	rText := []rune(text)
	l.buffer = rText
	l.cap = len(rText)
	l.index = len(rText)
	l.dirty = true
}

func (l *Line) Clear() {
	l.buffer = make([]rune, 0, 8)
	l.cap = 0
	l.index = 0
	l.dirty = true
}

func (l *Line) GotoStart() {
	l.index = 0
}

func (l *Line) GotoEnd() {
	l.index = len(l.buffer)
}

func (l *Line) GetCursorIndex() int {
	return l.index
}

func (l *Line) String() string {
	return string(l.buffer)
}

func (l *Line) IsDirty() bool {
	return l.dirty
}

func (l *Line) Rendered() {
	l.dirty = false
}
