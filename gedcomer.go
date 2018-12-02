package gedcom

// NoIndent can be used with GEDCOMLine and GEDCOMString so that the output does
// not contain the indent-levels.
const NoIndent = -1

// GEDCOMLiner allows an instance to return the single-line GEDCOM value as a
// string. This excludes any child nodes.
type GEDCOMLiner interface {
	// GEDCOMLine will return a single GEDCOM line, excluding children and with
	// an optional indent-level.
	//
	// The indent will only be included if it is at least 0. If you want to use
	// GEDCOMLine to compare the string values of nodes or exclude the indent
	// you should use the NoIndent constant.
	//
	// It is not safe to invoke GEDCOMLine on a nil value. Use GEDCOMLine() for
	// a safer alternative.
	GEDCOMLine(indent int) string
}

// GEDCOMStringer allows an instance to be rendered as multi-line GEDCOM.
//
// GEDCOMStringer can be thought of as a superset of GEDCOMLiner, often sharing
// the same line. However, for some entities (such as a Document) it does not
// make sense to implement GEDCOMLiner as well.
type GEDCOMStringer interface {
	// GEDCOMString returns the multi-line GEDCOM that includes an optional
	// indent for each line.
	//
	// The indent will only be included if it is at least 0. If you want to use
	// GEDCOMString to compare the string values of nodes or exclude the indent
	// you should use the NoIndent constant.
	//
	// It is not safe to invoke GEDCOMString on a nil value. Use GEDCOMString()
	// for a safer alternative.
	GEDCOMString(indent int) string
}

// GEDCOMLine is the safer alternative to GEDCOMLiner.GEDCOMLine. It will handle
// nils gracefully, returning an empty string.
func GEDCOMLine(value GEDCOMLiner, indent int) string {
	if IsNil(value) {
		return ""
	}

	return value.GEDCOMLine(indent)
}

// GEDCOMString is the safer alternative to GEDCOMStringer.GEDCOMString. It will
// handle nils gracefully, returning an empty string.
func GEDCOMString(value GEDCOMStringer, indent int) string {
	if IsNil(value) {
		return ""
	}

	return value.GEDCOMString(indent)
}
