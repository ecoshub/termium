package panel

import "term/internal/models/dimension"

type Panel interface {
	GetSize() *dimension.D2
	GetBuffer() [][]rune
}
