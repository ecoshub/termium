package history

var (
	DefaultCapacity = 50
)

type History struct {
	lines    []string
	capacity int
	index    int
	cursor   int
}

func New(capacity int) *History {
	h := &History{capacity: capacity, lines: make([]string, 0, capacity)}
	h.Add("")
	return h
}

func (h *History) Add(line string) {
	if h.index >= h.capacity {
		h.lines = h.lines[1:]
	} else {
		h.index++
	}
	h.lines = append(h.lines, line)
	h.cursor = h.index
}

func (h *History) Clear() {
	h.lines = make([]string, 0, h.capacity)
	h.index = 0
	h.cursor = 0
	h.Add("")
}

func (h *History) Up() string {
	if len(h.lines) == 0 {
		return ""
	}
	h.cursor = (len(h.lines) + h.cursor - 1) % len(h.lines)
	return h.lines[h.cursor]
}

func (h *History) Down() string {
	if len(h.lines) == 0 {
		return ""
	}
	h.cursor = (len(h.lines) + h.cursor + 1) % len(h.lines)
	return h.lines[h.cursor]
}

func (h *History) Len() int {
	return len(h.lines)
}
