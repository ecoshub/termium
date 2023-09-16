package ansi

import "fmt"

func SetColor(line string, fg int, bg int) string {
	s := ""
	if fg != 0 {
		s += fmt.Sprintf(SetForegroundColor, fg)
	}
	if bg != 0 {
		s += fmt.Sprintf(SetBackgroundColor, bg)
	}
	s += line
	s += ResetAllModes
	return s
}
