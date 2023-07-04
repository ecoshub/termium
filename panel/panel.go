package panel

type Panel interface {
	GetSize() (int, int)
	GetBuffer() [][]rune
}
