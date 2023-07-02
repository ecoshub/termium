package panel

import "termium/internal/models/dimension"

type Panel interface {
	GetSize() *dimension.Vector
	GetBuffer() [][]rune
}
