package component

import "term/internal/models/dimension"

type Component interface {
	GetSize() *dimension.D2
	GetBuffer() [][]rune
}
