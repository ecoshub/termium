package panel

import "term/internal/models/dimension"

type Panel interface {
	GetSize() *dimension.Vector
	GetBuffer() [][]rune
}
