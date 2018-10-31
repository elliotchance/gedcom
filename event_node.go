package gedcom

// EventNode is a noteworthy happening related to an individual, a group, or an
// organization.
type EventNode struct {
	*SimpleNode
}

// EventNode creates a new EVEN node.
func NewEventNode(document *Document, value, pointer string, children []Node) *EventNode {
	return &EventNode{
		newSimpleNode(document, TagEvent, value, pointer, children),
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
// Two events are considered equal if they both contain the same date. If either
// or both contain more than one date only a single date must match both sides.
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
