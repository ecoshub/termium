package history

type History struct {
	capacity int
	index    int
	lines    []string
	cursor   int
}

func New(capacity int) *History {
	return &History{capacity: capacity, lines: make([]string, 0, capacity)}
}

func (h *History) Add(line string) {
	if h.index >= h.capacity {
		h.lines = h.lines[1:]
	} else {
		h.index++
	}
	h.lines = append(h.lines, line)
	h.cursor = h.index - 1
}

func (h *History) Clear() {
	h.lines = make([]string, 0, h.capacity)
	h.index = 0
	h.cursor = 0
}

func (h *History) Up() string {
	if len(h.lines) == 0 {
		return ""
	}
	temp := h.cursor
	h.cursor = (h.cursor + len(h.lines) - 1) % len(h.lines)
	return h.lines[temp]
}

func (h *History) Down() string {
	if len(h.lines) == 0 {
		return ""
	}
	temp := h.cursor
	h.cursor = (h.cursor + 1) % len(h.lines)
	return h.lines[temp]
}

func (h *History) Len() int {
	return len(h.lines)
}
