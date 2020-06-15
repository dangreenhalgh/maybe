package maybe

import (
	"encoding/json"
	"math"
)

// Float64 is an option type: it is either Some (a float64) or None
type Float64 struct {
	Some float64
	None bool
}

// NewFloat64 creates a Float64 from value. If the given value is NaN then it is
// assigned as None, which will be serialised as `null` in json.
func NewFloat64(value float64) Float64 {
	if math.IsNaN(value) {
		return Float64{None: true}
	}
	return Float64{Some: value, None: false}
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Float64) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}
