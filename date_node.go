package gedcom

import (
	"fmt"
	"math"
)

// DateNode represents a DATE node.
//
// See the full specification for dates in the documentation for Date.
type DateNode struct {
	*SimpleNode

	// Dates are expensive to parse so we should not attempt to parse the value
	// until it is needed. Also, if we have already parsed the value once it
	// should not be parsed again.
	alreadyParsed   bool
	parsedDateRange DateRange
}

// NewDateNode creates a new DATE node.
func NewDateNode(value string, children ...Node) *DateNode {
	return &DateNode{
		newSimpleNode(TagDate, value, "", children...),
		false, DateRange{},
	}
}

// If the node is nil both results will be zero dates.
func (node *DateNode) DateRange() (dateRange DateRange) {
	if node == nil {
		return NewZeroDateRange()
	}

	// Parsing dates is very expensive. Cache them.
	if node.alreadyParsed {
		return node.parsedDateRange
	}

	defer func(node *DateNode) {
		node.parsedDateRange = dateRange
		node.alreadyParsed = true
	}(node)

	return NewDateRangeWithString(node.Value())
}

// String returns the date range as defined in the specification of DateNode.
//
// There are too many combinations to document here, but the two main formats
// you will receive will look like:
//
//   Bet. Feb 1956 and Mar 1956
//   Abt. 13 Nov 1983
//
func (node *DateNode) String() string {
	startDate, endDate := node.StartAndEndDates()

	if startDate.Is(endDate) {
		return startDate.String()
	}

	return fmt.Sprintf("Bet. %s and %s", startDate, endDate)
}

// Years fulfils the Yearer interface and is a convenience for:
//
//   node.DateRange().Years()
//
func (node *DateNode) Years() float64 {
	return node.DateRange().Years()
}

func (node *DateNode) Similarity(node2 *DateNode, maxYears float64) float64 {
	if node == nil || node2 == nil {
		return 0.5
	}

	return node.DateRange().Similarity(node2.DateRange(), maxYears)
}

func (node *DateNode) Equals(node2 Node) bool {
	leftIsNil := IsNil(node)
	rightIsNil := IsNil(node2)

	if leftIsNil || rightIsNil {
		return false
	}

	if date2, ok := node2.(*DateNode); ok {
		return node.DateRange().Equals(date2.DateRange())
	}

	return false
}

func (node *DateNode) IsValid() bool {
	if node == nil {
		return false
	}

	return node.DateRange().IsValid()
}

func (node *DateNode) StartDate() Date {
	if node == nil {
		return Date{}
	}

	dateRange := node.DateRange()

	return dateRange.StartDate()
}

func (node *DateNode) EndDate() Date {
	if node == nil {
		return Date{}
	}

	dateRange := node.DateRange()

	return dateRange.EndDate()
}

func (node *DateNode) StartAndEndDates() (Date, Date) {
	if node == nil {
		return NewZeroDateRange().StartAndEndDates()
	}

	return node.DateRange().StartAndEndDates()
}

func (node *DateNode) IsExact() bool {
	if node == nil {
		return false
	}

	return node.DateRange().IsExact()
}

func (node *DateNode) IsPhrase() bool {
	if node == nil {
		return false
	}

	return node.DateRange().IsPhrase()
}

func (node *DateNode) Sub(node2 *DateNode) (min Duration, max Duration, errs error) {
	nodeStart, nodeEnd := node.DateRange().StartAndEndDates()
	node2Start, node2End := node2.DateRange().StartAndEndDates()

	errs = NewErrors(
		nodeStart.ParseError,
		nodeEnd.ParseError,
		node2Start.ParseError,
		node2End.ParseError,
	).Err()

	min = Duration(node.StartDate().Time().Sub(node2.StartDate().Time()))
	max = Duration(node.EndDate().Time().Sub(node2.EndDate().Time()))

	// Durations must always be positive.
	min = Duration(math.Abs(float64(min)))
	max = Duration(math.Abs(float64(max)))

	return
}

func (node *DateNode) Warnings() Warnings {
	if !node.IsValid() {
		return Warnings{
			NewUnparsableDateWarning(node),
		}
	}

	return nil
}
