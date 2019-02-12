package gedcom_test

import (
	"testing"
	"time"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestDate_Time(t *testing.T) {
	t.Run("StartDate", func(t *testing.T) {
		for date, test := range dateTests {
			t.Run(date, func(t *testing.T) {
				node := gedcom.NewDateNode(date)

				assert.Equal(t, node.StartDate().Time(), test.startTime)
			})
		}
	})

	t.Run("EndDate", func(t *testing.T) {
		for date, test := range dateTests {
			t.Run(date, func(t *testing.T) {
				node := gedcom.NewDateNode(date)

				assert.Equal(t, node.EndDate().Time(), test.endTime)
			})
		}
	})
}

func TestDate_String(t *testing.T) {
	tests := map[gedcom.Date]string{
		// Exact
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, nil}:     "",
		gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact, nil}:  "1932",
		gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact, nil}:  "Mar 1987",
		gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact, nil}: "24 Apr 1774",

		// Non-exact
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintBefore, nil}:     "Bef.",
		gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintAbout, nil}:   "Abt. 1932",
		gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintAfter, nil}:   "Aft. Mar 1987",
		gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintBefore, nil}: "Bef. 24 Apr 1774",
	}

	for date, expected := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, expected, date.String(), "%#+v", date)
		})
	}
}

func TestDate_Is(t *testing.T) {
	tests := []struct {
		date1, date2 gedcom.Date
		match        bool
	}{
		// Matches
		{
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, nil},
			true,
		},
		{
			gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact, nil},
			true,
		},
		{
			gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact, nil},
			true,
		},
		{
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact, nil},
			true,
		},
		{
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{24, 4, 1774, true, gedcom.DateConstraintExact, nil},
			true,
		},

		// Non-matches.
		{
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintAbout, nil},
			false,
		},
		{
			gedcom.Date{0, 0, 1933, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact, nil},
			false,
		},
		{
			gedcom.Date{0, 2, 1987, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact, nil},
			false,
		},
		{
			gedcom.Date{25, 4, 1774, false, gedcom.DateConstraintExact, nil},
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact, nil},
			false,
		},
		{
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintAbout, nil},
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintBefore, nil},
			false,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, test.match, test.date1.Is(test.date2),
				"%#+v %#+v", test.date1, test.date2)
		})
	}
}

func TestDate_Years(t *testing.T) {
	tests := []struct {
		date     gedcom.Date
		expected float64
	}{
		// Zero
		{gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, nil}, 0.0},

		// Year
		{gedcom.Date{0, 0, 750, false, gedcom.DateConstraintExact, nil}, 750.5},
		{gedcom.Date{0, 0, 1845, false, gedcom.DateConstraintExact, nil}, 1845.5},

		// Months
		{gedcom.Date{0, 1, 1845, false, gedcom.DateConstraintExact, nil}, 1845.0437158469945},
		{gedcom.Date{0, 3, 1999, false, gedcom.DateConstraintExact, nil}, 1999.204918032787},
		{gedcom.Date{0, 12, 1832, false, gedcom.DateConstraintExact, nil}, 1832.956403269755},

		// Days
		{gedcom.Date{1, 1, 1789, false, gedcom.DateConstraintExact, nil}, 1789.0027322404371},
		{gedcom.Date{31, 1, 1435, false, gedcom.DateConstraintExact, nil}, 1435.0846994535518},
		{gedcom.Date{1, 2, 1601, false, gedcom.DateConstraintExact, nil}, 1601.0874316939892},
		{gedcom.Date{1, 3, 845, false, gedcom.DateConstraintExact, nil}, 845.1639344262295},
		{gedcom.Date{31, 12, 2010, false, gedcom.DateConstraintExact, nil}, 2010.9972677595629},
	}

	for _, test := range tests {
		t.Run(test.date.String(), func(t *testing.T) {
			assert.Equal(t, test.expected, test.date.Years())
		})
	}
}

func TestDate_Equals(t *testing.T) {
	Equals := tf.Function(t, gedcom.Date.Equals)

	at14Jan1845 := gedcom.Date{14, 1, 1845, false, gedcom.DateConstraintExact, nil}
	abt14Jan1845 := gedcom.Date{14, 1, 1845, false, gedcom.DateConstraintAbout, nil}
	bef14Jan1845 := gedcom.Date{14, 1, 1845, false, gedcom.DateConstraintBefore, nil}
	aft14Jan1845 := gedcom.Date{14, 1, 1845, false, gedcom.DateConstraintAfter, nil}

	at15Jan1845 := gedcom.Date{15, 1, 1845, false, gedcom.DateConstraintExact, nil}
	abt15Jan1845 := gedcom.Date{15, 1, 1845, false, gedcom.DateConstraintAbout, nil}
	bef15Jan1845 := gedcom.Date{15, 1, 1845, false, gedcom.DateConstraintBefore, nil}
	aft15Jan1845 := gedcom.Date{15, 1, 1845, false, gedcom.DateConstraintAfter, nil}

	// Zero dates are equal.
	Equals(gedcom.Date{}, gedcom.Date{}).Returns(false)
	Equals(at14Jan1845, gedcom.Date{}).Returns(false)
	Equals(gedcom.Date{}, at14Jan1845).Returns(false)

	// Test matrix.
	Equals(at14Jan1845, at14Jan1845).Returns(true) // #4
	Equals(at14Jan1845, abt14Jan1845).Returns(true)
	Equals(at14Jan1845, bef14Jan1845).Returns(false)
	Equals(at14Jan1845, aft14Jan1845).Returns(false)
	Equals(at14Jan1845, at15Jan1845).Returns(false)
	Equals(at14Jan1845, abt15Jan1845).Returns(false)
	Equals(at14Jan1845, bef15Jan1845).Returns(true)
	Equals(at14Jan1845, aft15Jan1845).Returns(false)

	Equals(abt14Jan1845, at14Jan1845).Returns(true) // #12
	Equals(abt14Jan1845, abt14Jan1845).Returns(true)
	Equals(abt14Jan1845, bef14Jan1845).Returns(false)
	Equals(abt14Jan1845, aft14Jan1845).Returns(false)
	Equals(abt14Jan1845, at15Jan1845).Returns(false)
	Equals(abt14Jan1845, abt15Jan1845).Returns(false)
	Equals(abt14Jan1845, bef15Jan1845).Returns(false)
	Equals(abt14Jan1845, aft15Jan1845).Returns(false)

	Equals(bef14Jan1845, at14Jan1845).Returns(false) // #20
	Equals(bef14Jan1845, abt14Jan1845).Returns(false)
	Equals(bef14Jan1845, bef14Jan1845).Returns(true)
	Equals(bef14Jan1845, aft14Jan1845).Returns(false)
	Equals(bef14Jan1845, at15Jan1845).Returns(false)
	Equals(bef14Jan1845, abt15Jan1845).Returns(false)
	Equals(bef14Jan1845, bef15Jan1845).Returns(true)
	Equals(bef14Jan1845, aft15Jan1845).Returns(false)

	Equals(aft14Jan1845, at14Jan1845).Returns(false) // #28
	Equals(aft14Jan1845, abt14Jan1845).Returns(false)
	Equals(aft14Jan1845, bef14Jan1845).Returns(false)
	Equals(aft14Jan1845, aft14Jan1845).Returns(true)
	Equals(aft14Jan1845, at15Jan1845).Returns(true)
	Equals(aft14Jan1845, abt15Jan1845).Returns(false)
	Equals(aft14Jan1845, bef15Jan1845).Returns(false)
	Equals(aft14Jan1845, aft15Jan1845).Returns(false)

	Equals(at15Jan1845, at14Jan1845).Returns(false) // #36
	Equals(at15Jan1845, abt14Jan1845).Returns(false)
	Equals(at15Jan1845, bef14Jan1845).Returns(false)
	Equals(at15Jan1845, aft14Jan1845).Returns(true)
	Equals(at15Jan1845, at15Jan1845).Returns(true)
	Equals(at15Jan1845, abt15Jan1845).Returns(true)
	Equals(at15Jan1845, bef15Jan1845).Returns(false)
	Equals(at15Jan1845, aft15Jan1845).Returns(false)

	Equals(abt15Jan1845, at14Jan1845).Returns(false) // #44
	Equals(abt15Jan1845, abt14Jan1845).Returns(false)
	Equals(abt15Jan1845, bef14Jan1845).Returns(false)
	Equals(abt15Jan1845, aft14Jan1845).Returns(false)
	Equals(abt15Jan1845, at15Jan1845).Returns(true)
	Equals(abt15Jan1845, abt15Jan1845).Returns(true)
	Equals(abt15Jan1845, bef15Jan1845).Returns(false)
	Equals(abt15Jan1845, aft15Jan1845).Returns(false)

	Equals(bef15Jan1845, at14Jan1845).Returns(true) // #52
	Equals(bef15Jan1845, abt14Jan1845).Returns(false)
	Equals(bef15Jan1845, bef14Jan1845).Returns(false)
	Equals(bef15Jan1845, aft14Jan1845).Returns(false)
	Equals(bef15Jan1845, at15Jan1845).Returns(false)
	Equals(bef15Jan1845, abt15Jan1845).Returns(false)
	Equals(bef15Jan1845, bef15Jan1845).Returns(true)
	Equals(bef15Jan1845, aft15Jan1845).Returns(false)

	Equals(aft15Jan1845, at14Jan1845).Returns(false) // #60
	Equals(aft15Jan1845, abt14Jan1845).Returns(false)
	Equals(aft15Jan1845, bef14Jan1845).Returns(false)
	Equals(aft15Jan1845, aft14Jan1845).Returns(true)
	Equals(aft15Jan1845, at15Jan1845).Returns(false)
	Equals(aft15Jan1845, abt15Jan1845).Returns(false)
	Equals(aft15Jan1845, bef15Jan1845).Returns(false)
	Equals(aft15Jan1845, aft15Jan1845).Returns(true)
}

func TestDate_IsZero(t *testing.T) {
	IsZero := tf.Function(t, gedcom.Date.IsZero)

	IsZero(gedcom.Date{}).Returns(true)
	IsZero(gedcom.Date{14, 1, 1845, false, gedcom.DateConstraintExact, nil}).Returns(false)
	IsZero(gedcom.Date{0, 1, 1845, false, gedcom.DateConstraintExact, nil}).Returns(false)
	IsZero(gedcom.Date{0, 0, 1845, false, gedcom.DateConstraintExact, nil}).Returns(false)
	IsZero(gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, nil}).Returns(true)
	IsZero(gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, nil}).Returns(true)
	IsZero(gedcom.Date{0, 0, 0, false, gedcom.DateConstraintAfter, nil}).Returns(true)
}

func TestNewDateWithTime(t *testing.T) {
	NewDateWithTime := tf.Function(t, gedcom.NewDateWithTime)
	tm, err := time.Parse(time.UnixDate, "Mon Jan 2 15:04:05 MST 2006")
	assert.NoError(t, err)

	NewDateWithTime(time.Time{}, false).Returns(gedcom.Date{})
	NewDateWithTime(time.Time{}, true).Returns(gedcom.Date{})
	NewDateWithTime(tm, false).Returns(
		gedcom.Date{2, time.January, 2006, false, gedcom.DateConstraintExact, nil})
	NewDateWithTime(tm, true).Returns(
		gedcom.Date{2, time.January, 2006, true, gedcom.DateConstraintExact, nil})
}

func TestNewDateRangeWithNow(t *testing.T) {
	NewDateRangeWithNow := tf.Function(t, gedcom.NewDateRangeWithNow)
	now := time.Now()

	NewDateRangeWithNow().Returns(
		gedcom.Date{now.Day(), now.Month(), now.Year(), false, gedcom.DateConstraintExact, nil},
		gedcom.Date{now.Day(), now.Month(), now.Year(), true, gedcom.DateConstraintExact, nil},
	)
}
