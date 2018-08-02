package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDate_Time(t *testing.T) {
	t.Run("StartDate", func(t *testing.T) {
		for date, test := range dateTests {
			t.Run(date, func(t *testing.T) {
				node := gedcom.NewDateNode(date, "", nil)

				assert.Equal(t, node.StartDate().Time(), test.startTime)
			})
		}
	})

	t.Run("EndDate", func(t *testing.T) {
		for date, test := range dateTests {
			t.Run(date, func(t *testing.T) {
				node := gedcom.NewDateNode(date, "", nil)

				assert.Equal(t, node.EndDate().Time(), test.endTime)
			})
		}
	})
}

func TestDate_String(t *testing.T) {
	tests := map[gedcom.Date]string{
		// Exact
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}:     "",
		gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact}:  "1932",
		gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact}:  "Mar 1987",
		gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact}: "24 Apr 1774",

		// Non-exact
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintBefore}:     "Bef.",
		gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintAbout}:   "Abt. 1932",
		gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintAfter}:   "Aft. Mar 1987",
		gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintBefore}: "Bef. 24 Apr 1774",
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
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact},
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact},
			true,
		},
		{
			gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact},
			gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact},
			true,
		},
		{
			gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact},
			gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact},
			true,
		},
		{
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact},
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact},
			true,
		},
		{
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact},
			gedcom.Date{24, 4, 1774, true, gedcom.DateConstraintExact},
			true,
		},

		// Non-matches.
		{
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact},
			gedcom.Date{0, 0, 0, false, gedcom.DateConstraintAbout},
			false,
		},
		{
			gedcom.Date{0, 0, 1933, false, gedcom.DateConstraintExact},
			gedcom.Date{0, 0, 1932, false, gedcom.DateConstraintExact},
			false,
		},
		{
			gedcom.Date{0, 2, 1987, false, gedcom.DateConstraintExact},
			gedcom.Date{0, 3, 1987, false, gedcom.DateConstraintExact},
			false,
		},
		{
			gedcom.Date{25, 4, 1774, false, gedcom.DateConstraintExact},
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintExact},
			false,
		},
		{
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintAbout},
			gedcom.Date{24, 4, 1774, false, gedcom.DateConstraintBefore},
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
		{gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, 0.0},

		// Year
		{gedcom.Date{0, 0, 750, false, gedcom.DateConstraintExact}, 750.5},
		{gedcom.Date{0, 0, 1845, false, gedcom.DateConstraintExact}, 1845.5},

		// Months
		{gedcom.Date{0, 1, 1845, false, gedcom.DateConstraintExact}, 1845.0437158469945},
		{gedcom.Date{0, 3, 1999, false, gedcom.DateConstraintExact}, 1999.204918032787},
		{gedcom.Date{0, 12, 1832, false, gedcom.DateConstraintExact}, 1832.956403269755},

		// Days
		{gedcom.Date{1, 1, 1789, false, gedcom.DateConstraintExact}, 1789.0027322404371},
		{gedcom.Date{31, 1, 1435, false, gedcom.DateConstraintExact}, 1435.0846994535518},
		{gedcom.Date{1, 2, 1601, false, gedcom.DateConstraintExact}, 1601.0874316939892},
		{gedcom.Date{1, 3, 845, false, gedcom.DateConstraintExact}, 845.1639344262295},
		{gedcom.Date{31, 12, 2010, false, gedcom.DateConstraintExact}, 2010.9972677595629},
	}

	for _, test := range tests {
		t.Run(test.date.String(), func(t *testing.T) {
			assert.Equal(t, test.expected, test.date.Years())
		})
	}
}
