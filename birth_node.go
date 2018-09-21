package gedcom

// BirthNode is the event of entering into life.
type BirthNode struct {
	*SimpleNode
}

// NewBirthNode creates a new BIRT node.
func NewBirthNode(document *Document, value, pointer string, children []Node) *BirthNode {
	return &BirthNode{
		newSimpleNode(document, TagBirth, value, pointer, children),
	}
}

// Dates returns zero or more dates associated with the Birth.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
//
// If the node is nil the result will also be nil.
func (node *BirthNode) Dates() (result []*DateNode) {
	if node == nil {
		return nil
	}

	for _, n := range NodesWithTag(node, TagDate) {
		result = append(result, n.(*DateNode))
	}

	return
}

// Equal will always return true if both nodes are not nil.
//
// If either node is nil (including both) or if the right side is not a
// BirthNode then false is always returned. Otherwise Equals will always return
// true.
//
// The reason Equals always returns true is because Equals is a shallow test and
// an individual can only ever have one birth event. Therefore it is safe to
// assume that birth events themselves are equal, even if the children they
// contain are not.
//
// This logic is especially important for CompareNodes.
func (node *BirthNode) Equals(node2 Node) bool {
	if IsNil(node) || IsNil(node2) {
		return false
	}

	if _, ok := node2.(*BirthNode); !ok {
		return false
	}

	return true
}
