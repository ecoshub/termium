package panel

func ConstantText(line string) *Basic {
	bp := NewBasic(len(line), 1)
	bp.Write(0, line)
	return bp
}
