package gedcom

// BurialNode is the event of the proper disposing of the mortal remains of a
// deceased person.
type BurialNode struct {
	*SimpleNode
}

// NewBurialNode creates a new BURI node.
func NewBurialNode(document *Document, value, pointer string, children []Node) *BurialNode {
	return &BurialNode{
		newSimpleNode(document, TagBurial, value, pointer, children),
	}
}

// Dates returns zero or more dates associated with the burial.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
//
// If the node is nil the result will also be nil.
func (node *BurialNode) Dates() []*DateNode {
	return Dates(node)
}

// Equal will always return true if both nodes are not nil.
//
// If either node is nil (including both) or if the right side is not a
// BurialNode then false is always returned. Otherwise Equals will always
// return true.
//
// The reason Equals always returns true is because Equals is a shallow test and
// an individual can only ever have one burial event. Therefore it is safe to
// assume that burial events themselves are equal, even if the children they
// contain are not.
//
// This logic is especially important for CompareNodes.
func (node *BurialNode) Equals(node2 Node) bool {
	if IsNil(node) || IsNil(node2) {
		return false
	}

	if _, ok := node2.(*BurialNode); !ok {
		return false
	}

	return true
}
