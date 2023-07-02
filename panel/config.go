package panel

import "github.com/ecoshub/termium/models/dimension"

type Config struct {
	Size *dimension.Vector

	// FIXME delete later
	AutoDummyInput bool
}
