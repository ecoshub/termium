package ansi

import (
	"regexp"
)

const ansiRegex = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansiRegex)

func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}

func Match(str string) bool {
	return re.Match([]byte(str))
}
