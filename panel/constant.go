package panel

func ConstantText(line string) *BasePanel {
	bp := NewBasePanel(&Config{SizeX: len(line), SizeY: 1})
	bp.Insert(0, line)
	return bp
}
