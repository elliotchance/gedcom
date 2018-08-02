package gedcom_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/elliotchance/gedcom"
)

var stringTests = []struct {
	a, b string
	want float64
}{
	// Simple values
	{"", "", 0},
	{"Foo", "Foo", 1},
	{"foo bar", "foo bar", 1},
	{"Foo bar", "foo bar", 0.9333333333333333},
	{"foo bar", "Foo Bar", 0.8666666666666668},
	{"foo", "bar", 0},

	// Names
	{"Elliot Chance", "Elliot R. Chance", 0.9163461538461538},
	{"Elliot R Chance", "Elliot Chance", 0.9271794871794872},
	{"Elliot Rupert Chance", "Elliot R. Chance", 0.8975},
	{"Elliot Rupert Chance", "Elliot R d P Chance", 0.8678947368421053},
	{"Eliot Rupert Chance", "Elliot R. Chance", 0.7822055137844612},
	{"J Smith", "John Smith", 0.7814285714285714},
	{"John Smeeth", "John Smith", 0.9214141414141415},
	{"bob jones", "Bob Jones", 0.8962962962962963},

	// Places
	{", Connecticut", "Connecticut", 0.7972027972027972},
	{
		"28 Leinster Gardens, London, England, United Kingdom",
		"40 Augustus Rd, Birmingham, England, United Kingdom",
		0.6681987417281535,
	},
	{
		"Adams County, Adams County, Iowa, United States",
		"Adams, Umatilla County, Oregon, United States",
		0.8408101718740018,
	},
}

func TestJaroWinkler(t *testing.T) {
	for _, test := range stringTests {
		t.Run(test.a+"_"+test.b, func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.JaroWinkler(test.a, test.b))
		})
	}
}
