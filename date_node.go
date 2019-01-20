// Dates
//
// Dates in GEDCOM files can be very complex as they can cater for many
// scenarios:
//
// 1. Incomplete, like "Dec 1943"
//
// 2. Anchored, like "Aft. 3 Sep 2003" or "Before 1923"
//
// 3. Ranges, like "Bet. 4 Apr 1823 and 8 Apr 1823"
//
// 4. Phrases, like "(Foo Bar)"
//
// This package provides a very rich API for dealing with all kind of dates in a
// meaningful and sensible way. Some notable features include:
//
// 1. All dates, even though that specify an specific day have a minimum and
// maximum value that are their true bounds. This is especially important for
// larger date ranges like the whole month of "Jun 1945".
//
// 2. Upper and lower bounds of dates can be converted to the native Go
// time.Time object.
//
// 3. There is a Years function that provides a convenient way to normalise a
// date range into a number for easier distance and comparison measurements.
//
// 4. Algorithms for calculating the similarity of dates on a configurable
// parabola.
package gedcom

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

// DefaultMaxYearsForSimilarity is a sensible default for the Similarity
// function (maxYears) when comparing dates. The importance of maxYears is
// explained in DateNode.Similarity.
//
// Unless you need to ensure similarity values are retained correctly through
// versions you should use this constant instead of specifying a raw value to
// DateNode.Similarity. This value may change in time if a more accurate default
// is found.
//
// The gedcomtune tool was used to find an ideal value for this. Generally
// speaking 2 - 3 years yielded much the same result. Any further in either
// direction led to a drop in accuracy for matching individuals.
const DefaultMaxYearsForSimilarity = float64(3)

// DateNode represents a DATE node.
//
// A date in GEDCOM always represents a range contained between the StartDate()
// and EndDate(), even when it represents a single day, like "23 Jan 1921".
//
// Before diving into the full specs below you should be aware of the known
// limitations:
//
// 1. Only the Gregorian calendar with the English language (for month names)
// is currently supported.
//
// 2. You should only expect dates that are valid and within the range of Go's
// supported libraries will work correctly. That is years between 0 and 9999. It
// is possible that dates outside of this range may be interpreted correctly but
// you should not rely on that remaining the same.
//
// 3. There are surly more keyword combinations used in GEDCOM files than are
// documented below. Interpreting these dates is not necessarily guaranteed to
// work, not work or retain the same behaviour between releases. If you believe
// there are other known cases please open an issue or pull request.
//
// Now into the specification. There are two basic forms of a DATE value:
//
//   between date and date
//   date
//
// The second case is actually equivalent to the first case the the same "date"
// substituted twice.
//
// The "between" keyword can be any of (non case sensitive):
//
//   between
//   bet
//   bet.
//   from
//
// The "and" keyword can be one of (non case sensitive):
//
//   -
//   and
//   to
//
// A "date" has three basic forms:
//
//   prefix? day month year
//   prefix? month year
//   prefix? year
//
// The "prefix" is optional and can be used to indicate if the date is
// approximate or not with one of the following keywords:
//
//   abt
//   abt.
//   about
//   c.
//   circa
//
// Or, the "prefix" can be used to signify unbounded dates with one of the
// following keywords:
//
//   after
//   aft
//   aft.
//   before
//   bef
//   bef.
//
// The "day" must be an integer between 1 and 31 and can have a single
// proceeding zero, like "03". The day should be valid against the month used.
// The behavior is unexpected when using invalid dates like "31 Feb 1999", but
// you will likely not receive a date at all if it's invalid.
//
// The "month" must be one of the following strings (case in-sensitive):
//
//   apr
//   april
//   aug
//   august
//   dec
//   december
//   feb
//   february
//   jan
//   january
//   jul
//   july
//   jun
//   june
//   mar
//   march
//   may
//   nov
//   november
//   oct
//   october
//   sep
//   september
//
// The "year" must be an integer with a value between 0 and 9999 (as to conform
// to the restrictions of the Go time package). It may be possible to parse
// dates outside of this range but they behaviour is not defined.
//
// The "year" may be 1 to 4 digits but it always treated as the absolute year.
// The year 89 is treated as the year 89, not 1989, for example.
type DateNode struct {
	*SimpleNode
	didParseDate    bool
	parsedStartDate Date
	parsedEndDate   Date
}

// NewDateNode creates a new DATE node.
func NewDateNode(document *Document, value, pointer string, children []Node) *DateNode {
	return &DateNode{
		newSimpleNode(document, TagDate, value, pointer, children),
		false, Date{}, Date{},
	}
}

func (node *DateNode) parse(dateToTest string, layoutsToTry []string) (Date, error) {
	for _, layout := range layoutsToTry {
		ts, err := time.Parse(layout, dateToTest)
		if err == nil {
			return Date{
				Day:   ts.Day(),
				Month: ts.Month(),
				Year:  ts.Year(),
			}, nil
		}
	}

	// Cannot parse date.
	return Date{}, errors.New("cannot parse date")
}

var months = map[string]time.Month{
	"apr":       time.April,
	"april":     time.April,
	"aug":       time.August,
	"august":    time.August,
	"dec":       time.December,
	"december":  time.December,
	"feb":       time.February,
	"february":  time.February,
	"jan":       time.January,
	"january":   time.January,
	"jul":       time.July,
	"july":      time.July,
	"jun":       time.June,
	"june":      time.June,
	"mar":       time.March,
	"march":     time.March,
	"may":       time.May,
	"nov":       time.November,
	"november":  time.November,
	"oct":       time.October,
	"october":   time.October,
	"sep":       time.September,
	"september": time.September,
}

func parseMonthName(parts []string, monthPos int) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("cannot parse month")
	}

	monthName := strings.ToLower(parts[monthPos])

	return CleanSpace(monthName), nil
}

var dateRegexp = regexp.MustCompile(
	fmt.Sprintf(`(?i)^(%s|%s|%s)? ?(\d+ )?(\w+ )?(\d+)$`,
		DateWordsAbout, DateWordsBefore, DateWordsAfter))

func parseDateParts(dateString string, isEndOfRange bool) Date {
	parts := dateRegexp.FindStringSubmatch(dateString)

	// Place holders for the locations of each regexp group.
	constraintPos, dayPos, monthPos, yearPos := 1, 2, 3, 4

	monthName, err := parseMonthName(parts, monthPos)

	switch {
	case len(parts) == 0, // Could not match the regexp.
		err != nil: // The month is unknown.
		return Date{
			IsEndOfRange: isEndOfRange,
		}
	}

	day := Atoi(parts[dayPos])
	month := time.Month(months[monthName])
	year := Atoi(parts[yearPos])

	// Check the date is valid.
	_, err = time.Parse("_2 1 2006",
		fmt.Sprintf("%d %d %04d", day, month, year))
	if parts[dayPos] != "" && err != nil {
		return Date{
			IsEndOfRange: isEndOfRange,
			Constraint:   DateConstraintFromString(parts[constraintPos]),
		}
	}

	return Date{
		Day:          day,
		Month:        month,
		Year:         year,
		IsEndOfRange: isEndOfRange,
		Constraint:   DateConstraintFromString(parts[constraintPos]),
	}
}

var rangeRegexp = regexp.MustCompile(fmt.Sprintf(`(?i)^(%s) (.+) (%s) (.+)$`,
	DateWordsBetween, DateWordsAnd))

// If the node is nil both results will be zero dates.
func (node *DateNode) DateRange() (startDate Date, endDate Date) {
	if node == nil {
		return Date{}, Date{}
	}

	// Parsing dates is very expensive. Cache them.
	if node.didParseDate {
		return node.parsedStartDate, node.parsedEndDate
	}

	defer func(node *DateNode) {
		node.parsedStartDate = startDate
		node.parsedEndDate = endDate
		node.didParseDate = true
	}(node)

	dateString := CleanSpace(node.Value())

	// Try to match a range first.
	parts := rangeRegexp.FindStringSubmatch(dateString)
	if len(parts) > 0 {
		return parseDateParts(parts[2], false), parseDateParts(parts[4], true)
	}

	// Single date.
	return parseDateParts(dateString, false), parseDateParts(dateString, true)
}

// StartDate returns the start component date of DateRange.
func (node *DateNode) StartDate() Date {
	start, _ := node.DateRange()

	return start
}

// EndDate returns the end component date of DateRange.
func (node *DateNode) EndDate() Date {
	_, end := node.DateRange()

	return end
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
	startDate, endDate := node.DateRange()

	if startDate.Is(endDate) {
		return startDate.String()
	}

	return fmt.Sprintf("Bet. %s and %s", startDate, endDate)
}

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
func (node *DateNode) Years() float64 {
	start, end := node.DateRangeYears()

	return (start + end) / 2.0
}

// DateRangeYears returns the Years values for the DateRange.
func (node *DateNode) DateRangeYears() (float64, float64) {
	start, end := node.DateRange()

	return start.Years(), end.Years()
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
func (node *DateNode) Similarity(node2 *DateNode, maxYears float64) float64 {
	if node == nil || node2 == nil {
		return 0.5
	}

	leftYears := node.Years()
	rightYears := node2.Years()
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
func (node *DateNode) Equals(node2 Node) bool {
	leftIsNil := IsNil(node)
	rightIsNil := IsNil(node2)

	if leftIsNil || rightIsNil {
		return false
	}

	if date2, ok := node2.(*DateNode); ok {
		// Phrases can only be compared to themselves and they must be the exact
		// same value to be considered equal.
		if node.IsPhrase() && date2.IsPhrase() && node.value == date2.value {
			return true
		}

		// Invalid dates follow the same rules as phrases.
		if !node.IsValid() && !date2.IsValid() && node.value == date2.value {
			return true
		}

		// Compare dates by value range.
		matchStartDate := node.StartDate().Equals(date2.StartDate())
		matchEndDate := node.EndDate().Equals(date2.EndDate())

		return matchStartDate && matchEndDate
	}

	return false
}

// IsValid returns true only when the node is not nil and the start and end date
// are non-zero.
//
// A "zero date" (Date.IsZero) is a date that is missing the year, month and
// day. Even if there is other associated information this date is considered to
// be useless for most purposes.
//
// It is safe and completely valid to use IsValid on a nil node.
func (node *DateNode) IsValid() bool {
	if node == nil {
		return false
	}

	start, end := node.DateRange()

	return !start.IsZero() && !end.IsZero()
}

// IsExact will return true if the date range represents a single day with an
// exact constraint.
//
// See Date.IsExact for more information.
func (node *DateNode) IsExact() bool {
	start, end := node.DateRange()
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
func (node *DateNode) IsPhrase() bool {
	if node == nil {
		return false
	}

	v := node.value

	if len(v) == 0 {
		return false
	}

	firstLetter := v[0]

	// ghost:ignore
	lastLetter := v[len(v)-1]

	return firstLetter == '(' && lastLetter == ')'
}
