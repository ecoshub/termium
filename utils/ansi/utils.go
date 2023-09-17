package ansi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ecoshub/termium/utils"
)

const ansiRegex = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansiRegex)

func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}

func Match(str string) bool {
	return re.Match([]byte(str))
}

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

func SetBoldStyle(line string) string {
	return SetBold + line + ResetBold
}

func SetItalicStyle(line string) string {
	return SetItalic + line + ResetItalic
}

func SetUnderlineStyle(line string) string {
	return SetUnderline + line + ResetUnderline
}

func SetBlinkStyle(line string) string {
	return SetBlink + line + ResetBlink
}

func ClearLine(line string, limit int) string {
	line = Strip(line)
	line = strings.TrimSpace(line)
	if len(line) > limit {
		line = utils.CutUnicode(line, limit-4)
		line += "..."
	}
	return line
}
