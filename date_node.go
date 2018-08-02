package gedcom

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

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
}

func NewDateNode(value, pointer string, children []Node) *DateNode {
	return &DateNode{
		&SimpleNode{
			tag:      TagName,
			value:    value,
			pointer:  pointer,
			children: children,
		},
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

	return CleanSpace(strings.ToLower(parts[monthPos])), nil
}

func parseDateParts(dateString string, isEndOfRange bool) Date {
	dateRegexp := fmt.Sprintf(`(?i)^(%s|%s|%s)? ?(\d+ )?(\w+ )?(\d+)$`,
		DateWordsAbout, DateWordsBefore, DateWordsAfter)
	parts := regexp.MustCompile(dateRegexp).FindStringSubmatch(dateString)

	// Place holders for the locations of each regexp group.
	constraintPos, dayPos, monthPos, yearPos := 1, 2, 3, 4

	// Could not match the regexp or month is unknown.
	monthName, err := parseMonthName(parts, monthPos)
	if len(parts) == 0 || err != nil {
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

func (node *DateNode) DateRange() (Date, Date) {
	dateString := CleanSpace(node.Value())

	// Try to match a range first.
	rangeRegexp := fmt.Sprintf(`(?i)^(%s) (.+) (%s) (.+)$`,
		DateWordsBetween, DateWordsAnd)
	parts := regexp.MustCompile(rangeRegexp).FindStringSubmatch(dateString)
	if len(parts) > 0 {
		return parseDateParts(parts[2], false), parseDateParts(parts[4], true)
	}

	// Single date.
	return parseDateParts(dateString, false), parseDateParts(dateString, true)
}

func (node *DateNode) StartDate() Date {
	start, _ := node.DateRange()

	return start
}

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

	return fmt.Sprintf("Bet. %s and %s", startDate.String(), endDate.String())
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
	start, end := node.DateRange()

	return (start.Years() + end.Years()) / 2.0
}
