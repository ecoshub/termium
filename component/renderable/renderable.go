package renderable

import (
	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/line"
)

type Renderable interface {
	Buffer() []*line.Line
	Configuration() *config.Config
	ListenChangeHandler(h func())
}
