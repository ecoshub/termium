package config

import "github.com/ecoshub/termium/component/style"

type Config struct {
	Width        int
	Height       int
	Title        string
	RenderTitle  bool
	TitleStyle   *style.Style
	ContentStyle *style.Style
}
