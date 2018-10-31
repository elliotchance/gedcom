package gedcom

// DateNodes is a slice of zero or more DateNodes. It may also be nil as other
// slice.
type DateNodes []*DateNode

// Minimum returns the date node with the minimum Years value from the provided
// slice.
//
// For date ranges the start value is used. For example "Between 1923 and 1943"
// is considered less than "Between 1924 and 1934" because 1923 is less than
// 1924. Even though the Years value (which is an average) would place them in
// the opposite order.
//
// If the slice is nil or contains zero elements then nil will be returned.
func (dates DateNodes) Minimum() *DateNode {
	min := (*DateNode)(nil)

	for _, date := range dates {
		if min == nil || date.StartDate().Years() < min.StartDate().Years() {
			min = date
		}
	}

	return min
}

// Maximum returns the date node with the maximum Years value from the provided
// slice.
//
// For date ranges the end value is used. For example "Between 1924 and 1934" is
// considered greater than "Between 1923 and 1943" because 1924 is greater than
// 1923. Even though the Years value (which is an average) would place them in
// the opposite order.
//
// If the slice is nil or contains zero elements then nil will be returned.
func (dates DateNodes) Maximum() *DateNode {
	min := (*DateNode)(nil)

	for _, date := range dates {
		if min == nil || date.EndDate().Years() > min.EndDate().Years() {
			min = date
		}
	}

	return min
}

// StripZero returns a new slice that only contains dates that are not zero.
//
// A "zero date" (Date.IsZero) is a date that is missing the year, month and
// day. Even if there is other associated information this date is considered to
// be useless for most purposes.
func (dates DateNodes) StripZero() (validDates DateNodes) {
	for _, date := range dates {
		if date.IsValid() {
			validDates = append(validDates, date)
		}
	}

	return
}
