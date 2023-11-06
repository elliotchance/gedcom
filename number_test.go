package gedcom_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var numberTests = map[gedcom.Number]struct {
	upperRoman string
	lowerRoman string
	err        error
}{
	// 0 is a special case.
	0: {"N", "n", nil},

	// We check all of the numbers up to 20 as that will cover all the
	// combinations with several different roman numerals.
	1:  {"I", "i", nil},
	2:  {"II", "ii", nil},
	3:  {"III", "iii", nil},
	4:  {"IV", "iv", nil},
	5:  {"V", "v", nil},
	6:  {"VI", "vi", nil},
	7:  {"VII", "vii", nil},
	8:  {"VIII", "viii", nil},
	9:  {"IX", "ix", nil},
	10: {"X", "x", nil},
	11: {"XI", "xi", nil},
	12: {"XII", "xii", nil},
	13: {"XIII", "xiii", nil},
	14: {"XIV", "xiv", nil},
	15: {"XV", "xv", nil},
	16: {"XVI", "xvi", nil},
	17: {"XVII", "xvii", nil},
	18: {"XVIII", "xviii", nil},
	19: {"XIX", "xix", nil},
	20: {"XX", "xx", nil},

	// These are just some random prime numbers we see up to 1000.
	29:  {"XXIX", "xxix", nil},
	71:  {"LXXI", "lxxi", nil},
	113: {"CXIII", "cxiii", nil},
	229: {"CCXXIX", "ccxxix", nil},
	349: {"CCCXLIX", "cccxlix", nil},
	409: {"CDIX", "cdix", nil},
	541: {"DXLI", "dxli", nil},
	601: {"DCI", "dci", nil},
	733: {"DCCXXXIII", "dccxxxiii", nil},
	863: {"DCCCLXIII", "dccclxiii", nil},
	941: {"CMXLI", "cmxli", nil},

	// Larger numbers.
	1000:  {"M", "m", nil},
	1234:  {"MCCXXXIV", "mccxxxiv", nil},
	2000:  {"MM", "mm", nil},
	2001:  {"MMI", "mmi", nil},
	2235:  {"MMCCXXXV", "mmccxxxv", nil},
	3000:  {"MMM", "mmm", nil},
	3001:  {"MMMI", "mmmi", nil},
	4000:  {"MMMM", "mmmm", nil},
	4001:  {"MMMMI", "mmmmi", nil},
	5000:  {"MMMMM", "mmmmm", nil},
	5001:  {"MMMMMI", "mmmmmi", nil},
	9999:  {"MMMMMMMMMCMXCIX", "mmmmmmmmmcmxcix", nil},
	10000: {"", "", errors.New("number is greater than 9999: 10000")},
	10078: {"", "", errors.New("number is greater than 9999: 10078")},

	// Invalid numbers.
	-1:  {"", "", errors.New("negative number: -1")},
	-45: {"", "", errors.New("negative number: -45")},
}

func TestNumber_UpperRoman(t *testing.T) {
	for number, tt := range numberTests {
		t.Run(fmt.Sprintf("%d", number), func(t *testing.T) {
			got, err := number.UpperRoman()

			assertError(t, tt.err, err)
			assert.Equal(t, tt.upperRoman, got)
		})
	}
}

func TestNumber_LowerRoman(t *testing.T) {
	for number, tt := range numberTests {
		t.Run(fmt.Sprintf("%d", number), func(t *testing.T) {
			got, err := number.LowerRoman()

			assertError(t, tt.err, err)
			assert.Equal(t, tt.lowerRoman, got)
		})
	}
}

func TestNumber_Int(t *testing.T) {
	for number := range numberTests {
		t.Run(fmt.Sprintf("%d", number), func(t *testing.T) {
			got := number.Int()
			assert.Equal(t, int(number), got)
		})
	}
}

func TestNewNumberWithInt(t *testing.T) {
	NewNumberWithInt := tf.Function(t, gedcom.NewNumberWithInt)

	NewNumberWithInt(0).Returns(gedcom.Number(0))
	NewNumberWithInt(123).Returns(gedcom.Number(123))
	NewNumberWithInt(-23).Returns(gedcom.Number(-23))
	NewNumberWithInt(123456789).Returns(gedcom.Number(123456789))
}
