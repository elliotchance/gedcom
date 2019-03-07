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
	"regexp"
	"strconv"
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

// The constants are used in regular expressions and documented on DateNode.
//
// Pipes are used to separate the values to make the options easier to use in
// regular expressions. The first value of each constant is important as it is
// the default when converting back to a string.
const (
	DateWordsBetween = "Bet.|bet|between|from"
	DateWordsAnd     = "and|to|-"
	DateWordsAbout   = "Abt.|abt|about|c.|ca|ca.|cca|cca.|circa"
	DateWordsAfter   = "Aft.|aft|after"
	DateWordsBefore  = "Bef.|bef|before"
)

// Date is a single point in time.
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
//
// Values represented by a Date instance must be compatible with Go's time
// package. This only allows for date ranges of the year between 0 and 9999. So
// Date would not allow for BC/BCE dates.
//
// You should be careful about directly creating dates from the defined instance
// variables because they may contain 0 to signify that a date component was not
// provided. Unless you have a very special case you should use Time() to
// convert to a usable date.
type Date struct {
	// Day of the month. When the day is not provided (like "Feb 1990") this
	// will be 0.
	Day int

	// Month of the year. When the month is not provided (like "1999") this will
	// be 0.
	Month time.Month

	// Year number. Go only allows for date ranges of the year between 0 and
	// 9999. If this year is outside of that date you will not be able to use
	// the Time() function and you will probably run into all sort of other
	// trouble.
	Year int

	// IsEndOfRange signifies is this date is the start or end of the range
	// (provided by DateNode). This is important for Time() to create a
	// timestamp that is constrained to the lower or upper bound.
	//
	// For example if the date was "Feb 1822" the Time() function would return:
	//
	//    1 Feb 1822 00:00:00.000000000 // IsEndOfRange = false
	//   29 Feb 1822 23:59:59.999999999 // IsEndOfRange = true
	//
	IsEndOfRange bool

	// Constraint indicates if the date is the exact value specified,
	// approximate or bound to before or after its value. See the documentation
	// for DateConstraint for a full explanation.
	//
	// This value does not affect the value calculated by Time() but it may be
	// important in some cases to know the date may not be the value returned by
	// Time().
	Constraint DateConstraint

	// If the date cannot be parsed this will contain the error.
	ParseError error
}

// NewDateWithTime creates a new Date with the provided time.Time.
//
// It is important to note that a Date only has a resolution of a single day and
// does not take into account timezone information.
//
// The isEndOfRange must be provided to signal if the Date returned represents
// the start or end of the day since the minimum resolution is one day.
//
// The returned Date will have an Exact constraint.
//
// If t IsZero then a zero Date will be returned (see Date.IsZero).
func NewDateWithTime(t time.Time, isEndOfRange bool) Date {
	if t.IsZero() {
		return NewZeroDate()
	}

	return Date{
		Day:          t.Day(),
		Month:        t.Month(),
		Year:         t.Year(),
		IsEndOfRange: isEndOfRange,
		Constraint:   DateConstraintExact,
	}
}

// NewDateWithNow creates a two Dates that represents the the start and end of
// the current day. See NewDateWithTime for implementation details.
func NewDateRangeWithNow() DateRange {
	now := time.Now()
	start := NewDateWithTime(now, false)
	end := NewDateWithTime(now, true)

	return NewDateRange(start, end)
}

func (date Date) safeParse(s string) time.Time {
	d, err := time.Parse("_2 1 2006", s)
	if err != nil {
		return time.Time{}
	}

	return d
}

// Time returns the minimum or maximum (depending on IsEndOfRange)
// representation of the Date as a Go Time instance.
func (date Date) Time() time.Time {
	var d string

	switch {
	case date.Day != 0 && date.Month != 0 && date.Year != 0:
		// Best case scenario, a full DMY.
		d = fmt.Sprintf("%d %d %04d", date.Day, date.Month, date.Year)

	case date.Month != 0 && date.Year != 0:
		// The month and year should return the first day of that month.
		d = fmt.Sprintf("1 %d %04d", date.Month, date.Year)

	case date.Year != 0:
		// Just the year should return the first day of that year.
		d = fmt.Sprintf("1 1 %04d", date.Year)

	default:
		// There is no valid time, settle for a zeroed timestamp which would
		// represent the start of the year 0.
	}

	result := date.safeParse(d)

	// If the safeParse could not parse the date it will return a zero date.
	// Make sure we don't try to adjust the zero date.
	if date.IsEndOfRange && !result.IsZero() {
		switch {
		case date.Day != 0:
			result = result.AddDate(0, 0, 1)
		case date.Month != 0:
			result = result.AddDate(0, 1, 0)
		case date.Year != 0:
			result = result.AddDate(1, 0, 0)
		}

		result = result.Add(-time.Nanosecond)
	}

	return result
}

// String returns the date in one of the three forms:
//
//   17 Jul 1890
//   Jul 1890
//   1890
//
// All forms of the date may also be proceeded with one of the constraints:
//
//   Abt.
//   Aft.
//   Bef.
//
func (date Date) String() string {
	day := ""
	if date.Day != 0 {
		day = strconv.Itoa(date.Day)
	}

	monthName := ""
	if date.Month != 0 {
		monthName = date.Month.String()[:3]
	}

	year := ""
	if date.Year != 0 {
		year = strconv.Itoa(date.Year)
	}

	rawDate := fmt.Sprintf("%s %s %s %s",
		date.Constraint.String(), day, monthName, year)

	return CleanSpace(rawDate)
}

// Is compares two dates. Dates are only considered to be the same if the day,
// month, year and constraint are all the same.
//
// The IsEndOfRange property is not used as part of the comparison because it
// only affects the behaviour of Time().
func (date Date) Is(date2 Date) bool {
	if date.Day != date2.Day {
		return false
	}

	if date.Month != date2.Month {
		return false
	}

	if date.Year != date2.Year {
		return false
	}

	return date.Constraint == date2.Constraint
}

// Years returns the number of years of a date as a floating-point. It can be
// used as an approximation to get a general idea of how far apart dates are but
// should not be treated as an accurate representation of time.
//
// The smallest date unit in a GEDCOM is a day. For specific dates it is
// calculated as the number of days that have past, divided by the number of
// days in that year (to correct for leap years). For example "10 Oct 2009"
// would return 2009.860274.
//
// Since some date components can be missing (like the day or month) Years
// compensates by returning the midpoint (average) of the maximum and minimum
// value in days. For example the date "Mar 1945" is the equivalent to the
// average Years value of "1 Mar 1945" and "31 Mar 1945", returning 1945.205479.
//
// When only a year is provided 0.5 will be added to the year. For example,
// "1845" returns 1845.5. This is not the exact midpoint of the year but will be
// close enough for general calculations. You should not rely on 0.5 being
// returned (as part of a check) as this may change in the future.
//
// The value returned from Years is not effected by any other component of the
// date. Such as if the date is approximate ("Abt.", etc) or directional
// ("Bef.", "Aft.", etc). If this property is important to you will need to take
// it into account in an appropriate way.
func (date Date) Years() float64 {
	hasDay := date.Day != 0
	hasMonth := date.Month != 0
	hasYear := date.Year != 0

	if hasDay && hasMonth && hasYear {
		// Calculate the total number of days in this year so we can take into
		// account leap years. The easiest way to do this is by going to the
		// first day of the next year then moving back one day.
		//
		// We must add one day to make sure the last day of the year is less
		// than 1.0.
		t := date.Time()
		daysInYear := time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC).
			AddDate(0, 0, -1).YearDay() + 1

		fractional := float64(t.YearDay()) / float64(daysInYear)

		return float64(t.Year()) + fractional
	}

	if hasMonth && hasYear {
		start := Date{
			Day:   1,
			Month: date.Month,
			Year:  date.Year,
		}.Years()

		// Find the last day of the month. Using the same method as above.
		t := date.Time()
		lastDay := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.UTC).
			AddDate(0, 0, -1).Day()

		end := Date{
			Day:   lastDay,
			Month: date.Month,
			Year:  date.Year,
		}.Years()

		return (start + end) / 2
	}

	if hasYear {
		return float64(date.Year) + 0.5
	}

	return 0
}

// IsZero returns true if the day, month and year are not provided. No other
// attributes are taken into consideration.
func (date Date) IsZero() bool {
	zeroDay := date.Day == 0
	zeroMonth := date.Month == 0
	zeroYear := date.Year == 0

	return zeroDay && zeroMonth && zeroYear
}

// Equals compares two dates.
//
// Unlike Is(), Equals() takes into what the date and its constraint represents,
// rather than just its raw value.
//
// For example, "3 Sep 1943" == "Bef. Oct 1943" returns true because 3 Sep 1943
// is before Oct 1943.
//
// If either date (including both) is IsZero then false is always returned.
//
// If Is() is true when comparing both dates then true is always returned.
//
// Otherwise the comparison used is selected from the following matrix:
//
//          ----------- Left ----------
//          Exact  About  Before  After
//   Exact    A      A      B       C
//   About    A      A      D       D
//  Before    C      D      C       D
//   After    B      D      D       B
//
// A. A match if the day, month and year are all equal.
//
// B. Match if left.Years() > right.Years().
//
// C. Match if left.Years() < right.Years().
//
// D. Never a match.
func (date Date) Equals(date2 Date) bool {
	if date.IsZero() {
		return false
	}

	if date2.IsZero() {
		return false
	}

	if date.Is(date2) {
		return true
	}

	matchers := [][]func(d1, d2 Date) bool{
		{Date.equalsA, Date.equalsA, Date.equalsB, Date.equalsC},
		{Date.equalsA, Date.equalsA, Date.equalsD, Date.equalsD},
		{Date.equalsC, Date.equalsD, Date.equalsC, Date.equalsD},
		{Date.equalsB, Date.equalsD, Date.equalsD, Date.equalsB},
	}

	return matchers[date2.Constraint][date.Constraint](date, date2)
}

// See Equals.
func (date Date) equalsA(date2 Date) bool {
	if date.Day != date2.Day {
		return false
	}

	if date.Month != date2.Month {
		return false
	}

	return date.Year == date2.Year
}

// See Equals.
func (date Date) equalsB(date2 Date) bool {
	leftYears := date.Years()
	rightYears := date2.Years()

	return leftYears > rightYears
}

// See Equals.
func (date Date) equalsC(date2 Date) bool {
	leftYears := date.Years()
	rightYears := date2.Years()

	return leftYears < rightYears
}

// See Equals.
func (date Date) equalsD(date2 Date) bool {
	return false
}

// IsExact will return true all parts of the date are complete and the date
// constraint is exact.
//
// This is to say that is points to a specific day.
func (date Date) IsExact() bool {
	return date.Day != 0 && date.Constraint == DateConstraintExact
}

func (date Date) IsBefore(date2 Date) bool {
	leftYears := date.Years()
	rightYears := date2.Years()

	return leftYears < rightYears
}

func (date Date) IsAfter(date2 Date) bool {
	leftYears := date.Years()
	rightYears := date2.Years()

	return leftYears > rightYears
}

func (date Date) Sub(date2 Date) Duration {
	a := date.Time()
	b := date2.Time()

	// The Time() above will set ParseError if the date is invalid.
	isKnown := date.ParseError == nil
	isEstimate := !date.IsExact()

	return NewDuration(a.Sub(b), isKnown, isEstimate)
}

func NewZeroDate() Date {
	return Date{}
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
	if len(parts) == 0 {
		return Date{
			IsEndOfRange: isEndOfRange,
			ParseError:   fmt.Errorf("unable to parse date: %s", dateString),
		}
	}

	// Place holders for the locations of each regexp group.
	constraintPos, dayPos, monthPos, yearPos := 1, 2, 3, 4

	monthName, err := parseMonthName(parts, monthPos)
	if err != nil {
		return Date{
			IsEndOfRange: isEndOfRange,
			ParseError:   errors.New("the month is unknown"),
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
			ParseError:   err,
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
