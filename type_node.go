package gedcom

// TypeNode represents a further qualification to the meaning of the associated
// superior tag.
//
// The value does not have any computer processing reliability. It is more in
// the form of a short one or two word note that should be displayed any time
// the associated data is displayed.
type TypeNode struct {
	*SimpleNode
}

// NewTypeNode creates a new TYPE node.
func NewTypeNode(value string, children ...Node) *TypeNode {
	return &TypeNode{
		newSimpleNode(TagType, value, "", children...),
	}
}
