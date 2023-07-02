package panel

import "termium/internal/models/dimension"

type Config struct {
	Size *dimension.Vector

	// FIXME delete later
	AutoDummyInput bool
}
