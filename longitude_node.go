package gedcom

// LongitudeNode represents a value indicating a coordinate position on a line,
// plane, or space.
//
// New in Gedcom 5.5.1.
type LongitudeNode struct {
	*SimpleNode
}

// NewLongitudeNode creates a new LONG node.
func NewLongitudeNode(value string, children ...Node) *LongitudeNode {
	return &LongitudeNode{
		newSimpleNode(TagLongitude, value, "", children...),
	}
}
