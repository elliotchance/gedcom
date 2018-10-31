package gedcom

// DateNodes is a slice of zero or more DateNodes. It may also be nil as other
// slice.
type DateNodes []*DateNode

// Minimum returns the date node with the minimum Years value from the provided
// slice.
//
// If the slice is nil or contains zero elements then nil will be returned.
func (dates DateNodes) Minimum() *DateNode {
	min := (*DateNode)(nil)

	for _, date := range dates {
		if min == nil || date.Years() < min.Years() {
			min = date
		}
	}

	return min
}

// Maximum returns the date node with the maximum Years value from the provided
// slice.
//
// If the slice is nil or contains zero elements then nil will be returned.
func (dates DateNodes) Maximum() *DateNode {
	min := (*DateNode)(nil)

	for _, date := range dates {
		if min == nil || date.Years() > min.Years() {
			min = date
		}
	}

	return min
}
