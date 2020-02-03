package gedcom

import "github.com/elliotchance/gedcom/tag"

// LatitudeNode represents a value indicating a coordinate position on a line,
// plane, or space.
//
// New in Gedcom 5.5.1.
type LatitudeNode struct {
	*SimpleNode
}

// NewLatitudeNode creates a new LATI node.
func NewLatitudeNode(value string, children ...Node) *LatitudeNode {
	return &LatitudeNode{
		newSimpleNode(tag.TagLatitude, value, "", children...),
	}
}
