package gedcom

// BaptismNode is event of baptism (not LDS), performed in infancy or later.
//
// See also BAPL and CHR.
type BaptismNode struct {
	*SimpleNode
}

// NewBaptismNode creates a new BAPM node.
func NewBaptismNode(document *Document, value, pointer string, children []Node) *BaptismNode {
	return &BaptismNode{
		newSimpleNode(document, TagBaptism, value, pointer, children),
	}
}

// Dates returns zero or more dates associated with the Baptism.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
//
// If the node is nil the result will also be nil.
func (node *BaptismNode) Dates() []*DateNode {
	return Dates(node)
}

// Equal will always return true if both nodes are not nil.
//
// If either node is nil (including both) or if the right side is not a
// BaptismNode then false is always returned. Otherwise Equals will always
// return true.
//
// The reason Equals always returns true is because Equals is a shallow test and
// an individual can only ever have one baptism event. Therefore it is safe to
// assume that baptism events themselves are equal, even if the children they
// contain are not.
//
// This logic is especially important for CompareNodes.
func (node *BaptismNode) Equals(node2 Node) bool {
	if IsNil(node) || IsNil(node2) {
		return false
	}

	if _, ok := node2.(*BaptismNode); !ok {
		return false
	}

	return true
}
