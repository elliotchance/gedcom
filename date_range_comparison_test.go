package gedcom_test

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dateRangeComparisonTests = []struct {
	comparison                            gedcom.DateRangeComparison
	isEqual, isPartiallyEqual, isNotEqual bool
	string                                string
}{
	{gedcom.DateRangeComparisonInvalid, false, false, false, "DateRangeComparisonInvalid"},
	{gedcom.DateRangeComparisonEqual, true, false, false, "DateRangeComparisonEqual"},
	{gedcom.DateRangeComparisonInside, false, true, false, "DateRangeComparisonInside"},
	{gedcom.DateRangeComparisonInsideStart, false, true, false, "DateRangeComparisonInsideStart"},
	{gedcom.DateRangeComparisonInsideEnd, false, true, false, "DateRangeComparisonInsideEnd"},
	{gedcom.DateRangeComparisonOutside, false, true, false, "DateRangeComparisonOutside"},
	{gedcom.DateRangeComparisonOutsideStart, false, true, false, "DateRangeComparisonOutsideStart"},
	{gedcom.DateRangeComparisonOutsideEnd, false, true, false, "DateRangeComparisonOutsideEnd"},
	{gedcom.DateRangeComparisonPartiallyBefore, false, true, false, "DateRangeComparisonPartiallyBefore"},
	{gedcom.DateRangeComparisonPartiallyAfter, false, true, false, "DateRangeComparisonPartiallyAfter"},
	{gedcom.DateRangeComparisonBefore, false, false, true, "DateRangeComparisonBefore"},
	{gedcom.DateRangeComparisonAfter, false, false, true, "DateRangeComparisonAfter"},
	{gedcom.DateRangeComparisonEntirelyBefore, false, false, true, "DateRangeComparisonEntirelyBefore"},
	{gedcom.DateRangeComparisonEntirelyAfter, false, false, true, "DateRangeComparisonEntirelyAfter"},
}

func TestDateRangeComparison_IsEqual(t *testing.T) {
	for _, test := range dateRangeComparisonTests {
		t.Run(fmt.Sprintf("%d", test.comparison), func(t *testing.T) {
			assert.Equal(t, test.isEqual, test.comparison.IsEqual())
		})
	}
}

func TestDateRangeComparison_IsPartiallyEqual(t *testing.T) {
	for _, test := range dateRangeComparisonTests {
		t.Run(fmt.Sprintf("%d", test.comparison), func(t *testing.T) {
			assert.Equal(t, test.isPartiallyEqual, test.comparison.IsPartiallyEqual())
		})
	}
}

func TestDateRangeComparison_IsNotEqual(t *testing.T) {
	for _, test := range dateRangeComparisonTests {
		t.Run(fmt.Sprintf("%d", test.comparison), func(t *testing.T) {
			assert.Equal(t, test.isNotEqual, test.comparison.IsNotEqual())
		})
	}
}

func TestDateRangeComparison_String(t *testing.T) {
	for _, test := range dateRangeComparisonTests {
		t.Run(fmt.Sprintf("%d", test.comparison), func(t *testing.T) {
			assert.Equal(t, test.string, test.comparison.String())
		})
	}
}
