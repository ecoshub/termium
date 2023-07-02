package panel

import "github.com/ecoshub/termium/internal/models/dimension"

type Panel interface {
	GetSize() *dimension.Vector
	GetBuffer() [][]rune
}
