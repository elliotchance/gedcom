package gedcom

import (
	"reflect"
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
//   []Node
//
// If any of the inputs are not one of the above types then a panic is raised.
//
// Using nil as a Node or including nil as one of the elements for []Node will
// be ignored, so you should not receive any nil values in the output.
func Compound(nodes ...interface{}) []Node {
	result := []Node{}

	for _, n := range nodes {
		v := reflect.ValueOf(n)

		switch v.Kind() {
		case reflect.Invalid:
			// Ignore

		case reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if j := v.Index(i).Interface(); j != nil {
					result = append(result, j.(Node))
				}
			}

		default:
			result = append(result, v.Interface().(Node))
		}
	}

	return result
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
	if node == nil {
		return ""
	}

	return node.String()
}

// Places returns the shallow PlaceNodes for each of the provided nodes.
func Places(nodes ...Node) (places []*PlaceNode) {
	for _, node := range nodes {
		for _, n := range NodesWithTag(node, TagPlace) {
			places = append(places, n.(*PlaceNode))
		}
	}

	return
}

func maxInt64(values ...int64) (r int64) {
	if len(values) == 0 {
		return
	}

	r = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > r {
			r = values[i]
		}
	}

	return
}

func maxInt(values ...int) (r int) {
	if len(values) == 0 {
		return
	}

	r = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > r {
			r = values[i]
		}
	}

	return
}
