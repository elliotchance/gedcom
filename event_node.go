package gedcom

// EventNode is a noteworthy happening related to an individual, a group, or an
// organization.
type EventNode struct {
	*SimpleNode
}

// EventNode creates a new EVEN node.
func NewEventNode(value string, children ...Node) *EventNode {
	return &EventNode{
		newSimpleNode(TagEvent, value, "", children...),
	}
}

// Dates returns zero or more dates associated with the event.
//
// When more than one date is returned you should not assume that the order has
// any significance for the importance of the dates.
//
// If the node is nil the result will also be nil.
func (node *EventNode) Dates() DateNodes {
	return Dates(node)
}

// Equal tests if two events are the same.
//
// Two events are considered equal if:
//
// 1. They both contain the same date. If either or both contain more than one
// date only a single date must match both sides.
//
// 2. They both do not have any dates, but all other attributes are the same.
//
// If either node is nil (including both) or if the right side is not a
// EventNode then false is always returned.
func (node *EventNode) Equals(node2 Node) bool {
	if IsNil(node) {
		return false
	}

	if IsNil(node2) {
		return false
	}

	if n2, ok := node2.(*EventNode); ok {
		leftDates := node.Dates()
		rightDates := n2.Dates()

		for _, left := range leftDates {
			for _, right := range rightDates {
				if left.Equals(right) {
					return true
				}
			}
		}

		if len(leftDates) == 0 && len(rightDates) == 0 && node.Value() == node2.Value() {
			return DeepEqualNodes(node.Nodes(), node2.Nodes())
		}
	}

	return false
}

// Years returns the Years value of the minimum date node in the node.
func (node *EventNode) Years() float64 {
	if min := node.Dates().Minimum(); min != nil {
		return Years(min)
	}

	return 0
}
