package gedcom

// ResidenceNode is the act of dwelling at an address for a period of time.
type ResidenceNode struct {
	*SimpleNode
}

// NewResidenceNode creates a new RESI node.
func NewResidenceNode(document *Document, value, pointer string, children []Node) *ResidenceNode {
	return &ResidenceNode{
		NewSimpleNode(document, TagResidence, value, pointer, children),
	}
}

// Dates returns zero or more dates associated with the residence.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
func (node *ResidenceNode) Dates() (result []*DateNode) {
	for _, n := range NodesWithTag(node, TagDate) {
		result = append(result, n.(*DateNode))
	}

	return
}

// Equal tests if two residence events are the same.
//
// Two residence events are considered equal if they both contain the same date.
// If either or both contain more than one date only a single date must match
// both sides.
//
// If either node is nil (including both) or if the right side is not a
// ResidenceNode then false is always returned.
func (node *ResidenceNode) Equals(node2 Node) bool {
	if IsNil(node) || IsNil(node2) {
		return false
	}

	if n2, ok := node2.(*ResidenceNode); ok {
		leftDates := node.Dates()
		rightDates := n2.Dates()

		for _, left := range leftDates {
			for _, right := range rightDates {
				if left.Equals(right) {
					return true
				}
			}
		}
	}

	return false
}

// Years returns the Years value of the minimum date node in the node.
func (node *ResidenceNode) Years() float64 {
	if min := MinimumDateNode(node.Dates()); min != nil {
		return Years(min)
	}

	return 0
}
