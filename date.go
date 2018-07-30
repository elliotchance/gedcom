package gedcom

import (
	"fmt"
	"strconv"
	"time"
)

// The constants are used in regular expressions and documented on DateNode.
//
// Pipes are used to separate the values to make the options easier to use in
// regular expressions. The first value of each constant is important as it is
// the default when converting back to a string.
const (
	DateWordsBetween = "Bet.|bet|between|from"
	DateWordsAnd     = "and|to|-"
	DateWordsAbout   = "Abt.|abt|about|c.|circa"
	DateWordsAfter   = "Aft.|aft|after"
	DateWordsBefore  = "Bef.|bef|before"
)

// Date is a single point in time.
//
// Values represented by a Date instance must be compatible with Go's time
// package. This only allows for date ranges of the year between 0 and 9999. So
// Date would not allow for BC/BCE dates.
//
// You should be careful about directly creating dates from the defined instance
// variables because they may contain 0 to signify that a date component was not
// provided. Unless you have a very special case you should use Time() to
// convert to a usable date.
//
// See the full specification for dates in the documentation for DateNode.
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

	return CleanSpace(fmt.Sprintf("%s %s %s %s",
		date.Constraint.String(), day, monthName, year))
}

// Is compares two dates. Dates are only considered to be the same if the day,
// month, year and constraint are all the same.
//
// The IsEndOfRange property is not used as part of the comparison because it
// only affects the behaviour of Time().
func (date Date) Is(date2 Date) bool {
	return date.Day == date2.Day && date.Month == date2.Month &&
		date.Year == date2.Year && date.Constraint == date2.Constraint
}
