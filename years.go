package gedcom

// Yearer allows some kind of date node or value to return its date
// representation in a number of years. This is implemented in several ways, you
// should read the docs for each implementation for more details.
type Yearer interface {
	Years() float64
}

// Years is a safe way to fetch the Years() value from a value. If the value is
// nil or does not implement Yearer then 0.0 will be returned. Otherwise the
// value of Years() is returned.
func Years(v interface{}) float64 {
	if y, ok := v.(Yearer); ok {
		return y.Years()
	}

	if d, ok := v.(Node); ok {
		return Years(Dates(d).Minimum())
	}

	return 0
}
