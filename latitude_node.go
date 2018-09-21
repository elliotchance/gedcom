package gedcom

// LatitudeNode represents a value indicating a coordinate position on a line,
// plane, or space.
//
// New in Gedcom 5.5.1.
type LatitudeNode struct {
	*SimpleNode
}

// NewLatitudeNode creates a new LATI node.
func NewLatitudeNode(document *Document, value, pointer string, children []Node) *LatitudeNode {
	return &LatitudeNode{
		newSimpleNode(document, TagLatitude, value, pointer, children),
	}
}
