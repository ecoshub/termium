package panel

type Stack struct {
	*Basic
	index int
}

func NewStackPanel(conf *Config) *Stack {
	b := NewBasicPanel(conf)
	return &Stack{
		Basic: b,
	}
}

func (sp *Stack) Push(line string) {
	if sp.index >= sp.height {
		sp.lines = sp.lines[1:]
		sp.lines = append(sp.lines, line)
	} else {
		sp.lines[sp.index] = line
		sp.index++
	}
	sp.render()
}

func (sp *Stack) Clear() {
	sp.Basic.Clear()
	sp.index = 0
}
