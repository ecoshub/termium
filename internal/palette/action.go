package palette

type ActionCode int

const (
	ActionEsc          ActionCode = 0xf0
	ActionEnter        ActionCode = 0xf1
	ActionInnerEvent   ActionCode = 0xf2
	ActionCursorUpdate ActionCode = 0xf3
)

type Action struct {
	Action ActionCode
	Input  string
}

func NewAction(action ActionCode, input string) *Action {
	return &Action{Action: action, Input: input}
}
