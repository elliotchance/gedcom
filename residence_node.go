package gedcom

// ResidenceNode is the act of dwelling at an address for a period of time.
type ResidenceNode struct {
	*SimpleNode
}

// NewResidenceNode creates a new RESI node.
func NewResidenceNode(value string, children ...Node) *ResidenceNode {
	return &ResidenceNode{
		newSimpleNode(TagResidence, value, "", children...),
	}
}

// Dates returns zero or more dates associated with the residence.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
//
// If the node is nil the result will also be nil.
func (node *ResidenceNode) Dates() DateNodes {
	return Dates(node)
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
	if IsNil(node) {
		return false
	}

	if IsNil(node2) {
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

		// Residences are equal if they contain the same place but both do not
		// specify a date. NodesWithTag would be expensive to run all the time,
		// so only use it when we know both sides do not have a date.
		if len(leftDates)+len(rightDates) == 0 {
			leftPlaces := NodesWithTag(node, TagPlace)
			rightPlaces := NodesWithTag(node2, TagPlace)

			return DeepEqualNodes(leftPlaces, rightPlaces)
		}
	}

	return false
}

// Years returns the Years value of the minimum date node in the node.
func (node *ResidenceNode) Years() float64 {
	if min := node.Dates().Minimum(); min != nil {
		return Years(min)
	}

	return 0
}
