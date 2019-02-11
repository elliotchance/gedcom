package gedcom

import "fmt"

// DateRangeComparison describe how two date ranges (each can be described with
// a DateNode) relate to each other.
//
// Since we are comparing ranges of time, rather than absolute points there are
// several variations of results to be considered. It's easiest to explain
// visually where Left and Right are the date range operands respectively:
//
//                                                       Method returns true:
//   Left:                        |========|
//   Right:
//     Equal:                     |========|             IsEqual()
//     Inside:                    | <====> |             IsPartiallyEqual()
//     InsideStart:               |======> |             IsPartiallyEqual()
//     InsideEnd:                 | <======|             IsPartiallyEqual()
//     Outside:               <===+========+===>         IsPartiallyEqual()
//     OutsideStart:              |========+===>         IsPartiallyEqual()
//     OutsideEnd:            <===+========|             IsPartiallyEqual()
//     PartiallyBefore:       <===+===>    |             IsPartiallyEqual()
//     PartiallyAfter:            |     <==+===>         IsPartiallyEqual()
//     Before:                <===|        |             IsNotEqual()
//     After:                     |        |===>         IsNotEqual()
//     EntirelyBefore:        <=> |        |             IsNotEqual()
//     EntirelyAfter:             |        | <=>         IsNotEqual()
//
// The Simplified value is the DateRangeComparisonSimplified value which is
// derived from DateRangeComparison.Simplified().
type DateRangeComparison int

// Each of the constants below represent how the ranges cross over. See
// DateRangeComparison for a visual explanation.
const (
	// DateRangeComparisonInvalid is only used in cases of an error.
	DateRangeComparisonInvalid = DateRangeComparison(iota)

	// Equal means that both date ranges are exactly the same in start and end
	// date, or they have an equivalent other value:
	//
	//   Yes:
	//     3 Sep 1943            3 Sep 1943
	//     Sep 1943 - Mar 1944   Sep 1943 - Mar 1944
	//     (world war 2)         (world war 2)
	//
	//   No:
	//     3 Sep 1943            Sep 1943
	//     Sep 1943 - Mar 1944   Sep 1943 - 3 Mar 1944
	//     (world war 2)         (world war 1)
	//
	DateRangeComparisonEqual = iota

	// Inside means that the right operand is both smaller and encapsulated by
	// the greater range in both directions of the left operand. The Start and
	// End derivatives represent if certain boundaries are also equal.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   5 Sep 1943 - 17 Sep 1943
	//     Sep 1943                   3 Sep 1943 - 20 Sep 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 5 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 10 Oct 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonInside      = iota
	DateRangeComparisonInsideStart = iota
	DateRangeComparisonInsideEnd   = iota

	// Outside means that the right operand is larger in both directions. The
	// opposite of Inside. The Start and End derivatives represent if certain
	// boundaries are also equal.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   14 Jul 1943 - 10 Oct 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 5 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 10 Oct 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonOutside      = iota
	DateRangeComparisonOutsideStart = iota
	DateRangeComparisonOutsideEnd   = iota

	// PartiallyBefore means the right operand range surrounds the start of the
	// left operand.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 5 Sep 1943
	//     Sep 1943                   14 Jul 1943 - 10 Sep 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 10 Oct 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonPartiallyBefore = iota

	// PartiallyAfter means the right operand range surrounds the end of the
	// left operand.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   17 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   10 Sep 1943 - 10 Oct 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   10 Sep 1943 - 15 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 25 Sep 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonPartiallyAfter = iota

	// Before means the right operand range starts before the left operand's
	// before, but ends at the same value as the start of the left operand.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 3 Sep 1943
	//     Sep 1943                   14 Jul 1943 - 1 Sep 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 4 Sep 1943
	//     Sep 1943                   14 Jul 1943 - 10 Sep 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonBefore = iota

	// After means the right operand starts at the end of the left operand and
	// ends somewhere thereafter.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   20 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   30 Sep 1943 - 10 Oct 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   19 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 25 Sep 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonAfter = iota

	// EntirelyBefore means that the full range of the right operand is before
	// any of the range of the left operand.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   1 Sep 1943 - 2 Sep 1943
	//     Sep 1943                   3 Jul 1943 - 10 Jul 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   19 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 25 Sep 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonEntirelyBefore = iota

	// EntirelyAfter means that the full range of the right operand is after
	// any of the range of the left operand.
	//
	//   Yes:
	//     3 Sep 1943 - 20 Sep 1943   25 Sep 1943 - 30 Sep 1943
	//     Sep 1943                   3 Oct 1943 - 10 Oct 1943
	//   No:
	//     3 Sep 1943 - 20 Sep 1943   19 Sep 1943 - 25 Sep 1943
	//     Sep 1943                   20 Sep 1943 - 25 Sep 1943
	//     (world war 2)              (world war 2)
	//
	DateRangeComparisonEntirelyAfter = iota
)

// IsEqual returns true only if both date ranges are exactly the same.
//
// It's important to note that a DateRangeComparison can be in one of three
// simplified states: Equal, PartiallyEqual or NotEqual. So !IsEqual() is not
// the same as IsNotEqual().
func (c DateRangeComparison) IsEqual() bool {
	return c == DateRangeComparisonEqual
}

// IsPartiallyEqual returns true if any part of the date range touches or
// intersects another.
func (c DateRangeComparison) IsPartiallyEqual() bool {
	switch c {
	case DateRangeComparisonInside,
		DateRangeComparisonInsideStart,
		DateRangeComparisonInsideEnd,
		DateRangeComparisonOutside,
		DateRangeComparisonOutsideStart,
		DateRangeComparisonOutsideEnd,
		DateRangeComparisonPartiallyBefore,
		DateRangeComparisonPartiallyAfter:
		return true
	}

	return false
}

// IsNotEqual returns true if the two date ranges do not intersect at any point.
//
// It's important to note that a DateRangeComparison can be in one of three
// simplified states: Equal, PartiallyEqual or NotEqual. So !IsNotEqual() is not
// the same as IsEqual().
func (c DateRangeComparison) IsNotEqual() bool {
	switch c {
	case DateRangeComparisonBefore,
		DateRangeComparisonAfter,
		DateRangeComparisonEntirelyBefore,
		DateRangeComparisonEntirelyAfter:
		return true
	}

	return false
}

// String returns the name of the constant like
// "DateRangeComparisonInsideStart".
func (c DateRangeComparison) String() string {
	switch c {
	case DateRangeComparisonInvalid:
		return "DateRangeComparisonInvalid"

	case DateRangeComparisonEqual:
		return "DateRangeComparisonEqual"

	case DateRangeComparisonInside:
		return "DateRangeComparisonInside"

	case DateRangeComparisonInsideStart:
		return "DateRangeComparisonInsideStart"

	case DateRangeComparisonInsideEnd:
		return "DateRangeComparisonInsideEnd"

	case DateRangeComparisonOutside:
		return "DateRangeComparisonOutside"

	case DateRangeComparisonOutsideStart:
		return "DateRangeComparisonOutsideStart"

	case DateRangeComparisonOutsideEnd:
		return "DateRangeComparisonOutsideEnd"

	case DateRangeComparisonPartiallyBefore:
		return "DateRangeComparisonPartiallyBefore"

	case DateRangeComparisonPartiallyAfter:
		return "DateRangeComparisonPartiallyAfter"

	case DateRangeComparisonBefore:
		return "DateRangeComparisonBefore"

	case DateRangeComparisonAfter:
		return "DateRangeComparisonAfter"

	case DateRangeComparisonEntirelyBefore:
		return "DateRangeComparisonEntirelyBefore"

	case DateRangeComparisonEntirelyAfter:
		return "DateRangeComparisonEntirelyAfter"

	default:
		panic(fmt.Sprintf("invalid constant: %d", c))
	}
}
