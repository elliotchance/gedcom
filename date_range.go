package gedcom

import (
	"fmt"
	"math"
	"regexp"
	"time"
)

// A DateRange represents a period of time.
//
// The minimum possible period is 1 day and ranges only have a resolution of a
// single day.
//
// DateRanges should be considered immutable and are passed by value because of
// this. You should create a new DateRange to represent a new range rather than
// mutating an existing DateRange.
type DateRange struct {
	start, end     Date
	originalString string
}

func NewZeroDateRange() DateRange {
	return DateRange{}
}

func NewDateRangeWithString(s string) (dr DateRange) {
	defer func(originalString string) {
		dr.originalString = originalString
	}(s)

	dateString := CleanSpace(s)

	// Try to match a range first.
	parts := dateRangeRegexp.FindStringSubmatch(dateString)
	if len(parts) > 0 {
		datePart1 := parseDateParts(parts[2], false)
		datePart2 := parseDateParts(parts[4], true)
		dateRange := NewDateRange(
			datePart1,
			datePart2,
		)

		return dateRange
	}

	// Single date.
	datePart1 := parseDateParts(dateString, false)
	datePart2 := parseDateParts(dateString, true)

	return NewDateRange(
		datePart1,
		datePart2,
	)
}

// NewDateRange creates a new date range between two provided dates. It is
// expected that the start date be less than or equal to the end date.
func NewDateRange(start, end Date) DateRange {
	start.IsEndOfRange = false
	end.IsEndOfRange = true

	return DateRange{
		start: start,
		end:   end,
	}
}

// Describes the matrix of possible ranges where each letter represents Before,
// Equal or After. A lower-case letter refers to the lower boundary. Conversely
// an upper-case letter refers to the upper boundary.
var dateRangeCompareMatrix = map[string]DateRangeComparison{
	"bb": DateRangeComparisonEntirelyBefore,
	"be": DateRangeComparisonBefore,
	"ba": DateRangeComparisonPartiallyBefore,
	"bB": DateRangeComparisonPartiallyBefore,
	"bE": DateRangeComparisonOutsideEnd,
	"bA": DateRangeComparisonOutside,
	"eb": DateRangeComparisonInvalid,
	"ee": DateRangeComparisonInsideStart,
	"ea": DateRangeComparisonInsideStart,
	"eB": DateRangeComparisonInsideStart,
	"eE": DateRangeComparisonEqual,
	"eA": DateRangeComparisonOutsideStart,
	"ab": DateRangeComparisonInvalid,
	"ae": DateRangeComparisonInvalid,
	"aa": DateRangeComparisonInside,
	"aB": DateRangeComparisonInside,
	"aE": DateRangeComparisonInsideEnd,
	"aA": DateRangeComparisonPartiallyAfter,
	// Bx is the same as ax.
	"Eb": DateRangeComparisonInvalid,
	"Ee": DateRangeComparisonInvalid,
	"Ea": DateRangeComparisonInvalid,
	"EB": DateRangeComparisonInvalid,
	"EE": DateRangeComparisonInsideEnd,
	"EA": DateRangeComparisonAfter,
	"Ab": DateRangeComparisonInvalid,
	"Ae": DateRangeComparisonInvalid,
	"Aa": DateRangeComparisonInvalid,
	"AB": DateRangeComparisonInvalid,
	"AE": DateRangeComparisonInvalid,
	"AA": DateRangeComparisonEntirelyAfter,
}

func compareDatesForLetter(value, start, end Date) string {
	// We only deal with whole days. This is needed for dates that are ending
	// dates so we don't get the 23:59:59.999 part.
	valueTime := value.Time().Truncate(24 * time.Hour)
	startTime := start.Time().Truncate(24 * time.Hour)
	endTime := end.Time().Truncate(24 * time.Hour)

	switch {
	case valueTime.Equal(startTime):
		return "e"

	case valueTime.Equal(endTime):
		return "E"

	case valueTime.Before(startTime):
		return "b"

	case valueTime.After(endTime):
		return "A"
	}

	// a and B would be the same thing.
	return "a"
}

func (dr DateRange) Compare(dr2 DateRange) DateRangeComparison {
	start := compareDatesForLetter(dr.start, dr2.start, dr2.end)
	end := compareDatesForLetter(dr.end, dr2.start, dr2.end)

	return dateRangeCompareMatrix[start+end]
}

// Start is the lower boundary of the date range.
func (dr DateRange) StartDate() Date {
	return dr.start
}

// End is the upper boundary of the date range.
func (dr DateRange) EndDate() Date {
	return dr.end
}

// Before returns true if the start date is before the other start date.
//
// The idea of "before" in the context of overlapping date ranges is ambiguous.
// The simplest way to think treat all these situations is to only look at the
// start date for each range. No matter when the end dates are or how much of
// each other then end up overlapping.
func (dr DateRange) IsBefore(dr2 DateRange) bool {
	return dr.start.IsBefore(dr2.start)
}

// After returns true if the end date is after the other end date.
//
// See Before for a more detailed explanation.
func (dr DateRange) IsAfter(dr2 DateRange) bool {
	return dr.end.IsAfter(dr2.end)
}

var dateRangeRegexp = regexp.MustCompile(
	fmt.Sprintf(`(?i)^(%s) (.+) (%s) (.+)$`, DateWordsBetween, DateWordsAnd))

// Years works in a similar way to Date.Years() but also takes into
// consideration the StartDate() and EndDate() values of a whole date range,
// like "Bet. 1943 and 1945". It does this by averaging out the Years() value of
// the StartDate() and EndDate() values.
//
// If the DateNode has a single date, like "Mar 1937" then Years will return the
// same value as the Years on the start or end date (no average will be used.)
//
// You can read the specific conversion rules in Date.Years() but be aware that
// the returned value is an approximation and should not be used in date
// calculations.
func (dr DateRange) Years() float64 {
	return (dr.StartDate().Years() + dr.EndDate().Years()) / 2.0
}

// Similarity returns a value from 0.0 to 1.0 to identify how similar two dates
// (or date ranges) are to each other. 1.0 would mean that the dates are exactly
// the same, whereas 0.0 would mean that they are not similar at all.
//
// Similarity is safe to use when either date is nil. If either side is nil then
// 0.5 is returned. Not because they are similar but because there is not enough
// information to make the distinction either way. This is important when using
// date comparisons in combination or part of larger calculations where missing
// data on both sides does not lead to very low scores unnecessarily.
//
// The returned value is calculated on a parabola that awards higher values to
// dates that are proportionally closer to each other. That is, dates that are
// twice as close will have more than twice the score. This attempts to satisfy
// a usable comparison values for close specific dates as well as more relaxed
// values (such as those that one provide an approximate year).
//
// Only the difference between dates is used in the calculation so it is not
// affected by time lines. That is to say that the difference between the years
// 500 and 502 would return the same similarity as the years 2000 to 2002.
//
// The maxYears allows the error margin to be adjusted. Dates that are further
// apart (in any direction) than maxYears will always return 0.0.
//
// A greater maxYears can be used when dates are less exact (such as ancient
// dates that could be commonly off by 10 years or more) or a smaller value when
// dealing with recent dates that are provided in a more exact form.
//
// A sensible default value for maxYears is provided with
// DefaultMaxYearsForSimilarity. You should use this if you are unsure. There is
// also more explanation on the constant.
func (dr DateRange) Similarity(dr2 DateRange, maxYears float64) float64 {
	leftYears := dr.Years()
	rightYears := dr2.Years()
	yearsApart := leftYears - rightYears
	similarity := math.Pow(yearsApart/maxYears, 2)

	// When one date is invalid the similarity will go asymptotic.
	if similarity > 1 {
		return 0
	}

	return 1 - similarity
}

// Equals compares the values of two dates taking into consideration the date
// constraint.
//
// If either date is nil then false is always returned. Even if both dates are
// nil.
//
// A DateNode is considered to be equal only when its StartDate() and EndDate()
// both equal their respective values in the other DateNode.
//
// The comparisons of dates is quite complicated. See the documentation for
// Date.Equals for a full explanation.
func (dr DateRange) Equals(dr2 DateRange) bool {
	// Phrases can only be compared to themselves and they must be the exact
	// same value to be considered equal.
	if dr.IsPhrase() && dr2.IsPhrase() && dr.originalString == dr2.originalString {
		return true
	}

	// Invalid dates follow the same rules as phrases.
	if !dr.IsValid() && !dr2.IsValid() && dr.originalString == dr2.originalString {
		return true
	}

	// Compare dates by value range.
	matchStartDate := dr.StartDate().Equals(dr2.StartDate())
	matchEndDate := dr.EndDate().Equals(dr2.EndDate())

	return matchStartDate && matchEndDate
}

func (dr DateRange) StartAndEndDates() (Date, Date) {
	return dr.StartDate(), dr.EndDate()
}

// IsValid returns true only when the node is not nil and the start and end date
// are non-zero.
//
// A "zero date" (Date.IsZero) is a date that is missing the year, month and
// day. Even if there is other associated information this date is considered to
// be useless for most purposes.
//
// It is safe and completely valid to use IsValid on a nil node.
func (dr DateRange) IsValid() bool {
	start, end := dr.StartAndEndDates()

	return !start.IsZero() && !end.IsZero()
}

// IsExact will return true if the date range represents a single day with an
// exact constraint.
//
// See Date.IsExact for more information.
func (dr DateRange) IsExact() bool {
	start, end := dr.StartAndEndDates()
	startIsExact := start.IsExact()
	endIsExact := end.IsExact()

	return startIsExact && endIsExact
}

// IsPhrase returns true if the date value is a phrase.
//
// A phrase is any statement offered as a date when the year is not
// recognizable to a date parser, but which gives information about when an
// event occurred. The date phrase is enclosed in matching parentheses.
//
// IsPhrase is safe to use on a nil DateNode, and will return false.
func (dr DateRange) IsPhrase() bool {
	if len(dr.originalString) == 0 {
		return false
	}

	firstLetter := dr.originalString[0]

	// ghost:ignore
	lastLetter := dr.originalString[len(dr.originalString)-1]

	return firstLetter == '(' && lastLetter == ')'
}

func (dr DateRange) ParseError() error {
	if err := dr.StartDate().ParseError; err != nil {
		return err
	}

	if err := dr.EndDate().ParseError; err != nil {
		return err
	}

	return nil
}

func (dr DateRange) String() string {
	start, end := dr.StartAndEndDates()
	if start.Equals(end) {
		return start.String()
	}

	return fmt.Sprintf("Bet. %s and %s", start, end)
}
