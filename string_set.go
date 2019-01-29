package gedcom

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// StringSet represents a slice of strings where all elements are unique (a
// set).
type StringSet struct {
	elements sync.Map // map[string]struct{}
}

// NewStringSet creates a new StringSet that is initialized with zero or more
// elements.
//
// See Add() for semantics.
func NewStringSet(elements ...string) *StringSet {
	ss := &StringSet{}

	return ss.Add(elements...)
}

// Add will append zero more strings to the set. If any of the elements already
// exist in the set they will be ignored.
func (ss *StringSet) Add(elements ...string) *StringSet {
	for _, element := range elements {
		ss.elements.Store(element, nil)
	}

	return ss
}

// Has returns true if the string exists in the set.
func (ss *StringSet) Has(element string) bool {
	_, ok := ss.elements.Load(element)

	return ok
}

// Intersects returns true if the two sets contain at least one value that is
// the same.
func (ss *StringSet) Intersects(ss2 *StringSet) bool {
	found := false
	ss.elements.Range(func(key, _ interface{}) bool {
		if ss2.Has(key.(string)) {
			found = true
			return false
		}

		return true
	})

	return found
}

// Iterate will iterate all items, similar to a "for range".
//
// The order of the elements will change between successive iterations.
//
// The iteration will stop if the fn returns false.
func (ss *StringSet) Iterate(fn func(string) bool) {
	ss.elements.Range(func(key, _ interface{}) bool {
		if !fn(key.(string)) {
			return false
		}

		return true
	})
}

// Len returns the number of elements in the set.
func (ss *StringSet) Len() (len int) {
	ss.elements.Range(func(_, _ interface{}) bool {
		len++
		return true
	})

	return
}

// Strings returns all the values as a slice of strings. The result is always
// sorted.
func (ss *StringSet) Strings() (elements []string) {
	ss.Iterate(func(element string) bool {
		elements = append(elements, element)

		return true
	})

	sort.Strings(elements)

	return
}

// String returns all of the items sorted, in the form of:
//
//   (bar,foo)
//
func (ss *StringSet) String() string {
	return fmt.Sprintf("(%s)", strings.Join(ss.Strings(), ","))
}
