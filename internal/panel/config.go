package panel

import "term/internal/models/dimension"

type Config struct {
	Size *dimension.Vector

	// FIXME delete later
	AutoDummyInput bool
}
