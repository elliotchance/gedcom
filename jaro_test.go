package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stringTests = []struct {
	a, b      string
	jaro, str float64
}{
	// Simple values
	{"", "", 0, 0},
	{"Foo", "Foo", 1, 1},
	{"foo bar", "foo bar", 1, 1},
	{"Foo bar", "foo bar", 0.9333333333333333, 1},
	{"foo bar", "Foo Bar", 0.8666666666666668, 1},
	{"foo", "bar", 0, 0},

	// Names
	{"Elliot Chance", "Elliot R. Chance", 0.9163461538461538, 0.9271794871794872},
	{"Elliot R Chance", "Elliot Chance", 0.9271794871794872, 0.9271794871794872},
	{"Elliot Rupert Chance", "Elliot R. Chance", 0.8975, 0.91},
	{"Elliot Rupert Chance", "Elliot R d P Chance", 0.8678947368421053, 0.8784210526315789},
	{"Eliot Rupert Chance", "Elliot R. Chance", 0.7822055137844612, 0.7977610693400166},
	{"J Smith", "John Smith", 0.7814285714285714, 0.7814285714285714},
	{"John Smeeth", "John Smith", 0.9214141414141415, 0.9214141414141415},
	{"bob jones", "Bob Jones", 0.8962962962962963, 1},
	{"Elliot   Chance", "Elliot Chance", 0.9271794871794872, 1},
	{"Elliot Rupert Chance", "Elliot 'Rupert' Chance", 0.9218181818181819, 1},

	// Places
	{", Connecticut", "Connecticut", 0.7972027972027972, 1},
	{
		"28 Leinster Gardens, London, England, United Kingdom",
		"40 Augustus Rd, Birmingham, England, United Kingdom",
		0.6681987417281535,
		0.6800831443688585,
	},
	{
		"Adams County, Adams County, Iowa, United States",
		"Adams, Umatilla County, Oregon, United States",
		0.8408101718740018,
		0.8632912132912133,
	},
}

func TestJaroWinkler(t *testing.T) {
	for _, test := range stringTests {
		t.Run(test.a+"_"+test.b, func(t *testing.T) {
			assert.Equal(t, test.jaro, gedcom.JaroWinkler(test.a, test.b))
		})
	}
}

func TestStringSimilarity(t *testing.T) {
	for _, test := range stringTests {
		t.Run(test.a+"_"+test.b, func(t *testing.T) {
			assert.Equal(t, test.str, gedcom.StringSimilarity(test.a, test.b))
		})
	}
}
