package gedcom

import (
	"fmt"
	"reflect"
)

type Node interface {
	fmt.Stringer

	// The node itself.
	Tag() Tag
	Value() string
	Pointer() string
	Document() *Document
	SetDocument(document *Document)

	// Child nodes.
	Nodes() []Node
	AddNode(node Node)

	// Comparison.
	Equals(node2 Node) bool
}

// IsNil is the safe and reliable way to check if a node is nil. You should not
// compare Node values with an untyped nil as it will lead to unexpected
// results.
//
// As a side node IsNil cannot be part of the Node interface because more
// specific node types (such as DateNode) use SimpleNode as an instance variable
// and that would cause a nil pointer panic.
func IsNil(node Node) bool {
	return node == nil || reflect.ValueOf(node).IsNil()
}
