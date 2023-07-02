package panel

import "github.com/ecoshub/termium/models/dimension"

type Panel interface {
	GetSize() *dimension.Vector
	GetBuffer() [][]rune
}
