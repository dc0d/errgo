package errgo

import "fmt"

//-----------------------------------------------------------------------------

const markerFormat = "%s/%02d:%s()"

// Marker prepend err with markerFormat
type Marker struct {
	name string
	err  error
}

func (m *Marker) Error() string {
	var cause string
	if m.err != nil {
		cause = m.err.Error()
	} else {
		cause = "CAUSE:N/A"
	}
	return m.name + ": " + cause
}

// Cause implements Causer interface
func (m *Marker) Cause() error { return m.err }

// NewMarker whereever it be called, will prepend the other error with:
// markerFormat
func NewMarker(cause error) error {
	if cause == nil {
		return nil
	}
	name := makeMarker(markerFormat)
	return &Marker{
		name: name,
		err:  cause,
	}
}

//-----------------------------------------------------------------------------

type sentinelErr string

func (v sentinelErr) Error() string                { return string(v) }
func errorf(format string, a ...interface{}) error { return sentinelErr(fmt.Sprintf(format, a...)) }

//-----------------------------------------------------------------------------
