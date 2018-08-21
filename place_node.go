package gedcom

// PlaceNode represents a place (location).
type PlaceNode struct {
	*SimpleNode
}

func NewPlaceNode(document *Document, value, pointer string, children []Node) *PlaceNode {
	return &PlaceNode{
		NewSimpleNode(document, TagPlace, value, pointer, children),
	}
}
