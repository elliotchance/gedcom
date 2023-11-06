package gedcom_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/elliotchance/gedcom/v39"
	"github.com/stretchr/testify/assert"
)

//                              4 Sep    19 Sep
//                            3 Sep |    | 20 Sep
//                          2 Sep | |    | | 21 Sep
//                        1 Sep | | |    | | | 30 Sep
//                            v v v v    v v v v
//   Left:                    .   |========|   .
//   Right                    .                .
//     Equal:                 .   |========|   .
//     Inside:                .   | <====> |   .
//     InsideStart:           .   |======> |   .
//     InsideEnd:             .   | <======|   .
//     Outside:               <===+========+===>
//     OutsideStart:          .   |========+===>
//     OutsideEnd:            <===+========|   .
//     PartiallyBefore:       <===+=>      |   .
//     PartiallyAfter:        .   |      <=+===>
//     Before:                <===|        |   .
//     After:                 .   |        |===>
//     EntirelyBefore:        <=> |        |   .
//     EntirelyAfter:         .   |        | <=>
//
var dateRangeCompareTest = map[string]struct {
	comparison    gedcom.DateRangeComparison
	before, after bool
}{
	"1_1": {
		gedcom.DateRangeComparisonEntirelyBefore,
		true,
		false,
	},
	"1_2": {
		gedcom.DateRangeComparisonEntirelyBefore,
		true,
		false,
	},
	"1_3": {
		gedcom.DateRangeComparisonBefore,
		true,
		false,
	},
	"1_4": {
		gedcom.DateRangeComparisonPartiallyBefore,
		true,
		false,
	},
	"1_19": {
		gedcom.DateRangeComparisonPartiallyBefore,
		true,
		false,
	},
	"1_20": {
		gedcom.DateRangeComparisonOutsideEnd,
		true,
		false,
	},
	"1_21": {
		gedcom.DateRangeComparisonOutside,
		true,
		true,
	},
	"1_30": {
		gedcom.DateRangeComparisonOutside,
		true,
		true,
	},

	"2_2": {
		gedcom.DateRangeComparisonEntirelyBefore,
		true,
		false,
	},
	"2_3": {
		gedcom.DateRangeComparisonBefore,
		true,
		false,
	},
	"2_4": {
		gedcom.DateRangeComparisonPartiallyBefore,
		true,
		false,
	},
	"2_19": {
		gedcom.DateRangeComparisonPartiallyBefore,
		true,
		false,
	},
	"2_20": {
		gedcom.DateRangeComparisonOutsideEnd,
		true,
		false,
	},
	"2_21": {
		gedcom.DateRangeComparisonOutside,
		true,
		true,
	},
	"2_30": {
		gedcom.DateRangeComparisonOutside,
		true,
		true,
	},

	"3_3": {
		gedcom.DateRangeComparisonInsideStart,
		false,
		false,
	},
	"3_4": {
		gedcom.DateRangeComparisonInsideStart,
		false,
		false,
	},
	"3_19": {
		gedcom.DateRangeComparisonInsideStart,
		false,
		false,
	},
	"3_20": {
		gedcom.DateRangeComparisonEqual,
		false,
		false,
	},
	"3_21": {
		gedcom.DateRangeComparisonOutsideStart,
		false,
		true,
	},
	"3_30": {
		gedcom.DateRangeComparisonOutsideStart,
		false,
		true,
	},

	"4_4": {
		gedcom.DateRangeComparisonInside,
		false,
		false,
	},
	"4_19": {
		gedcom.DateRangeComparisonInside,
		false,
		false,
	},
	"4_20": {
		gedcom.DateRangeComparisonInsideEnd,
		false,
		false,
	},
	"4_21": {
		gedcom.DateRangeComparisonPartiallyAfter,
		false,
		true,
	},
	"4_30": {
		gedcom.DateRangeComparisonPartiallyAfter,
		false,
		true,
	},

	"19_19": {
		gedcom.DateRangeComparisonInside,
		false,
		false,
	},
	"19_20": {
		gedcom.DateRangeComparisonInsideEnd,
		false,
		false,
	},
	"19_21": {
		gedcom.DateRangeComparisonPartiallyAfter,
		false,
		true,
	},
	"19_30": {
		gedcom.DateRangeComparisonPartiallyAfter,
		false,
		true,
	},

	"20_20": {
		gedcom.DateRangeComparisonInsideEnd,
		false,
		false,
	},
	"20_21": {
		gedcom.DateRangeComparisonAfter,
		false,
		true,
	},
	"20_30": {
		gedcom.DateRangeComparisonAfter,
		false,
		true,
	},

	"21_21": {
		gedcom.DateRangeComparisonEntirelyAfter,
		false,
		true,
	},
	"21_30": {
		gedcom.DateRangeComparisonEntirelyAfter,
		false,
		true,
	},

	"30_30": {
		gedcom.DateRangeComparisonEntirelyAfter,
		false,
		true,
	},
}

func dateRangeCompareRange(s string) gedcom.DateRange {
	var startDay, endDay int

	fmt.Sscanf(s, "%d_%d", &startDay, &endDay)

	start := gedcom.Date{
		Day:   startDay,
		Month: time.September,
		Year:  1943,
	}

	end := gedcom.Date{
		Day:   endDay,
		Month: time.September,
		Year:  1943,
	}

	return gedcom.NewDateRange(start, end)
}

func TestDateRange_Compare(t *testing.T) {
	base := dateRangeCompareRange("3_20")

	for testName, test := range dateRangeCompareTest {
		t.Run(testName, func(t *testing.T) {
			dr := dateRangeCompareRange(testName)
			expected := test.comparison.String()
			actual := dr.Compare(base).String()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestDateRange_IsBefore(t *testing.T) {
	base := dateRangeCompareRange("3_20")

	for testName, test := range dateRangeCompareTest {
		t.Run(testName, func(t *testing.T) {
			dr := dateRangeCompareRange(testName)
			assert.Equal(t, test.before, dr.IsBefore(base))
		})
	}
}

func TestDateRange_IsAfter(t *testing.T) {
	base := dateRangeCompareRange("3_20")

	for testName, test := range dateRangeCompareTest {
		t.Run(testName, func(t *testing.T) {
			dr := dateRangeCompareRange(testName)
			assert.Equal(t, test.after, dr.IsAfter(base))
		})
	}
}
