package gedcom

// DeathNode is the event when mortal life terminates.
type DeathNode struct {
	*SimpleNode
}

// NewDeathNode creates a new DEAT node.
func NewDeathNode(value string, children ...Node) *DeathNode {
	return &DeathNode{
		newSimpleNode(TagDeath, value, "", children...),
	}
}

// Dates returns zero or more dates associated with the death.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
//
// If the node is nil the result will also be nil.
func (node *DeathNode) Dates() DateNodes {
	return Dates(node)
}

// Equal will always return true if both nodes are not nil.
//
// If either node is nil (including both) or if the right side is not a
// DeathNode then false is always returned. Otherwise Equals will always
// return true.
//
// The reason Equals always returns true is because Equals is a shallow test and
// an individual can only ever have one death event. Therefore it is safe to
// assume that death events themselves are equal, even if the children they
// contain are not.
//
// This logic is especially important for CompareNodes.
func (node *DeathNode) Equals(node2 Node) bool {
	if IsNil(node) {
		return false
	}

	if IsNil(node2) {
		return false
	}

	if _, ok := node2.(*DeathNode); !ok {
		return false
	}

	return true
}
