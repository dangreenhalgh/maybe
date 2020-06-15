package maybe

import (
	"encoding/json"
)

// Int is an option type: it is either Some (an int) or None
type Int struct {
	Some int
	None bool
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Int) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}

// Int8 is an option type: it is either Some (an int8) or None
type Int8 struct {
	Some int8
	None bool
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Int8) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}

// Int16 is an option type: it is either Some (an int16) or None
type Int16 struct {
	Some int16
	None bool
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Int16) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}

// Int32 is an option type: it is either Some (an int32) or None
type Int32 struct {
	Some int32
	None bool
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Int32) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}

// Int64 is an option type: it is either Some (an int64) or None
type Int64 struct {
	Some int64
	None bool
}

// MarshalJSON implements the Marshaller interface. Decodes None values to `null`.
func (o Int64) MarshalJSON() ([]byte, error) {
	if o.None {
		return json.Marshal(nil)
	}
	return json.Marshal(o.Some)
}
