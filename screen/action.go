package screen

type ActionCode int

const (
	KeyActionEsc        ActionCode = 0xf0
	KeyActionEnter      ActionCode = 0xf1
	KeyActionInnerEvent ActionCode = 0xf2
)

type KeyAction struct {
	Action ActionCode
	Input  string
}

func newAction(action ActionCode, input string) *KeyAction {
	return &KeyAction{Action: action, Input: input}
}
