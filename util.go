package gedcom

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func valueToPointer(val string) string {
	valLen := len(val)
	firstCharIsAt := val[0] == '@'
	lastCharIsAt := val[valLen-1] == '@'
	if valLen > 2 && firstCharIsAt && lastCharIsAt {
		return val[1 : valLen-1]
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
func Atoi(mutS string) int {
	// Trim off leading zeros and surrounding spaces as they affect how the
	// integer may be parsed.
	mutS = strings.TrimLeft(mutS, "0")
	mutS = strings.TrimSpace(mutS)

	i, _ := strconv.Atoi(mutS)

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
func CleanSpace(s0 string) string {
	// Replace twice if there is an odd number of spaces in a row.
	s1 := strings.Replace(s0, "  ", " ", -1)
	s2 := strings.Replace(s1, "  ", " ", -1)

	// Trim whatever spaces are left on either side.
	return strings.TrimSpace(s2)
}

// First returns the first non-nil node of nodes. If the length of nodes is zero
// or a non-nil value is not found then nil is returned.
//
// First is useful in combination with other functions like:
//
//   birth := First(individual.Births())
//
func First(nodes interface{}) Node {
	n := Compound(nodes)
	if len(n) == 0 {
		return nil
	}

	return n[0]
}

// Last returns the last non-nil node of nodes. If the length of nodes is zero
// or a non-nil value is not found then nil is returned.
//
// Last is useful in combination with other functions like:
//
//   death := Last(individual.Deaths())
//
func Last(nodes interface{}) Node {
	n := Compound(nodes)
	for i := len(n) - 1; i >= 0; i-- {
		if n[i] != nil {
			return n[i]
		}
	}

	return nil
}

// Value is a safe way to fetch the Value() from a node. If the node is nil then
// an empty string will be returned.
func Value(node Node) string {
	if IsNil(node) {
		return ""
	}

	return node.Value()
}

// Compound is a easier way to join a collection of nodes. The input type is
// flexible to allow the following types:
//
//   nil
//   Node
//   Nodes
//
// If any of the inputs are not one of the above types then a panic is raised.
//
// Using nil as a Node or including nil as one of the elements for Nodes will
// be ignored, so you should not receive any nil values in the output.
func Compound(nodes ...interface{}) Nodes {
	mutResult := Nodes{}

	for _, n := range nodes {
		v := reflect.ValueOf(n)

		switch v.Kind() {
		case reflect.Invalid:
			// Ignore

		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if j := v.Index(i).Interface(); j != nil {
					mutResult = append(mutResult, j.(Node))
				}
			}

		default:
			mutResult = append(mutResult, v.Interface().(Node))
		}
	}

	return mutResult
}

// NodeCondition is a convenience method for inline conditionals.
func NodeCondition(condition bool, node1, node2 Node) Node {
	if condition {
		return node1
	}

	return node2
}

// Pointer is a safe way to fetch the Pointer() from a node. If the node is nil
// then an empty string will be returned.
func Pointer(node Node) string {
	if IsNil(node) {
		return ""
	}

	return node.Pointer()
}

// String is a safe way to fetch the String() from a node. If the node is nil
// then an empty string will be returned.
func String(node Node) string {
	if IsNil(node) {
		return ""
	}

	return node.String()
}

func maxInt64(values ...int64) int64 {
	valuesLen := len(values)

	if valuesLen == 0 {
		return 0
	}

	mutMaximumValue := values[0]
	for i := 1; i < valuesLen; i++ {
		if values[i] > mutMaximumValue {
			mutMaximumValue = values[i]
		}
	}

	return mutMaximumValue
}

func maxInt(values ...int) int {
	valuesLen := len(values)

	if valuesLen == 0 {
		return 0
	}

	mutMaximumValue := values[0]
	for i := 1; i < valuesLen; i++ {
		if values[i] > mutMaximumValue {
			mutMaximumValue = values[i]
		}
	}

	return mutMaximumValue
}

// DateAndPlace is a convenience method for fetching a date and place from a
// list of nodes.
//
// If multiple dates and places exist it will choose the first respective one.
func DateAndPlace(nodes ...Node) (date *DateNode, place *PlaceNode) {
	for _, node := range nodes {
		dates := Dates(node.(Node))
		places := Places(node.(Node))

		if date == nil && len(dates) > 0 {
			date = dates[0]
		}

		if place == nil && len(places) > 0 {
			place = places[0]
		}
	}

	return
}

func positiveDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}

	return d
}
