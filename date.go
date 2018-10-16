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
