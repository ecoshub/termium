package panel

import "term/internal/models/dimension"

type Config struct {
	Size *dimension.D2

	// FIXME delete later
	AutoDummyInput bool
}
