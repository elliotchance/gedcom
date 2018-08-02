package gedcom

import (
	"strconv"
	"strings"
)

func valueToPointer(val string) string {
	if len(val) > 2 && val[0] == '@' && val[len(val)-1] == '@' {
		return val[1 : len(val)-1]
	}

	return ""
}

// Atoi is a fault tolerant way to convert a string to an integer.
//
// Atoi ultimately uses strconv.Atoi to do the conversion, but will clean the
// string by removing any surrounding spaces or processing "0" characters.
//
// It is usually the logic behind any function in this package that expects to
// receive an integer from a string value.
//
// If the string cannot be parsed to an integer then 0 is returned.
func Atoi(s string) int {
	// Trim off leading zeros and surrounding spaces as they affect how the
	// integer may be parsed.
	s = strings.TrimLeft(s, "0")
	s = strings.TrimSpace(s)

	i, _ := strconv.Atoi(s)

	return i
}

// CleanSpace works similar to strings.TrimSpace except that it also replaces
// consecutive spaces anywhere in the string with a single space.
//
//   "  Foo   bar BAZ" -> "Foo bar BAZ"
//
// CleanSpace is used in many places throughout the library to clean values that
// are known to not place any significance on their spaces. Such as individual
// and place names.
func CleanSpace(s string) string {
	// Replace twice if there is an odd number of spaces in a row.
	s = strings.Replace(s, "  ", " ", -1)
	s = strings.Replace(s, "  ", " ", -1)

	// Trim whatever spaces are left on either side.
	s = strings.TrimSpace(s)

	return s
}

// Coalesce returns the first non-nil item. If there are no items provided or
// all items are nil then nil will be returned, otherwise the first non-nil
// item.
func Coalesce(items ...interface{}) interface{} {
	for _, item := range items {
		if item != nil {
			return item
		}
	}

	return nil
}
