package maybe

import (
	"encoding/json"
	"time"
)

// Time is an option type: it is either Some (a time.Time) or None
type Time struct {
	Some time.Time
	None bool
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Time) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}
