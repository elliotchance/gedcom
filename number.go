package gedcom

import (
	"fmt"
	"strings"
)

// Number can be used to transform a number into another representation.
//
// Even though a Number is an alias for int you should use the provided function
// Int() to get its value in case the underlying type changes in the future.
//
// It should also be noted that Number only supports integer values and many of
// the conversions or representations will only work with positive values. Some
// functions may also have more specific ranges.
type Number int

// NewNumberWithInt is the preferred way to create a new Number.
//
// While it is possible to cast the value directly to a Number, using this
// constructor is safer if the underlying type of Number changes.
func NewNumberWithInt(n int) Number {
	return Number(n)
}

// Int returns the true value of the number with no loss of precision.
func (n Number) Int() int {
	return int(n)
}

// UpperRoman returns uppercase Roman numerals, like "VII".
//
// The acceptable range is between 0 and 10000. Zero is a special case that will
// return "N", an abbreviation for "nulla". While it is possible to represent
// numbers number larger than 10000 it would get very cumbersome since each
// thousand would need an extra "M".
//
// If the number is outside of the acceptable range then an empty string is
// returned with an error.
//
// Also see LowerRoman.
func (n Number) UpperRoman() (string, error) {
	// Catch edge cases first.
	switch {
	case n == 0:
		return "N", nil

	case n < 0:
		return "", fmt.Errorf("negative number: %d", n)

	case n > 9999:
		return "", fmt.Errorf("number is greater than 9999: %d", n)
	}

	// Some of the following logic was copied from:
	// https://github.com/StefanSchroeder/Golang-Roman/blob/master/roman.go
	figure := []int{1000, 100, 10, 1}

	romanDigitA := []string{
		1:    "I",
		10:   "X",
		100:  "C",
		1000: "M",
	}

	romanDigitB := []string{
		1:    "V",
		10:   "L",
		100:  "D",
		1000: "M",
	}

	arg := n.Int()
	ret := ""
	var romanSlice []string
	x := ""

	// Correct for values greater than 4000.
	if arg >= 4000 {
		ms := arg / 1000
		ret = strings.Repeat("M", ms)
		arg -= ms * 1000
	}

	for _, f := range figure {
		digit, i, v := int(arg/f), romanDigitA[f], romanDigitB[f]
		switch digit {
		case 1:
			romanSlice = append(romanSlice, string(i))
		case 2:
			romanSlice = append(romanSlice, string(i)+string(i))
		case 3:
			romanSlice = append(romanSlice, string(i)+string(i)+string(i))
		case 4:
			romanSlice = append(romanSlice, string(i)+string(v))
		case 5:
			romanSlice = append(romanSlice, string(v))
		case 6:
			romanSlice = append(romanSlice, string(v)+string(i))
		case 7:
			romanSlice = append(romanSlice, string(v)+string(i)+string(i))
		case 8:
			romanSlice = append(romanSlice, string(v)+string(i)+string(i)+string(i))
		case 9:
			romanSlice = append(romanSlice, string(i)+string(x))
		}

		arg -= digit * f
		x = i
	}

	for _, e := range romanSlice {
		ret += e
	}

	return ret, nil
}

// LowerRoman returns the lowercase Roman numerals, like "vii".
//
// LowerRoman works with exactly the same rules and constraints as UpperRoman.
func (n Number) LowerRoman() (string, error) {
	roman, err := n.UpperRoman()

	return strings.ToLower(roman), err
}
