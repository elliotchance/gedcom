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

	// Equals performs a shallow comparison between two nodes.
	//
	// The implementation is different depending on the types of nodes being
	// compared. You should see the specific documentation for the Node.
	//
	// That being said, as a general rule is fair to assume that in most cases
	// the following things take place:
	//
	// 1. The default logic, if not overridden by the node is to consider nodes
	// equal if they have the all of the same; tag, value and pointer.
	//
	// 2. If either side is nil (or both) the nodes are never considered equal.
	//
	// 3. If both nodes are different types they are not considered equal.
	//
	// 4. Despite being a shallow equality test some nodes need to check the
	// equality some children to make a determination. For example, a BirthNode
	// is always equal to another BirthNode, even if they contain different
	// children or dates. Whereas a ResidenceNode is only considered equal if
	// they both contain at least one DateNode with exactly the same value.
	//
	// The rules above play heavily into the logic of other things, such as
	// CompareNodes.
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

// NodeGedcom is the recursive version of GedcomLine. It will render a node and
// all of its children (if any) as a multi-line GEDCOM data.
//
// Unlike rendering a Document, the root node will have a depth of 0 (like the
// root nodes of a Document) which means that you will almost certainly not be
// able to use the returned GEDCOM data as a whole of part of GEDCOM file.
//
// That being said, you should be able to parse this data as a Document and
// retain the same nested node structure originally encoded.
//
// This function is mostly useful for debugging and displaying complex nodes in
// the understandable and consistent form of GEDCOM data.
func NodeGedcom(node Node) string {
	if IsNil(node) {
		return ""
	}

	document := &Document{
		Nodes: []Node{node},
	}

	return document.String()
}
