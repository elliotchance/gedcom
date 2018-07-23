package gedcom

// PlaceNode represents a place (location).
type PlaceNode struct {
	*SimpleNode
}

func NewPlaceNode(value, pointer string, children []Node) *PlaceNode {
	return &PlaceNode{
		&SimpleNode{
			tag:      TagPlace,
			value:    value,
			pointer:  pointer,
			children: children,
		},
	}
}
