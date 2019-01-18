// Equality
//
// Node.Equals performs a shallow comparison between two nodes. The
// implementation is different depending on the types of nodes being compared.
// You should see the specific documentation for the Node.
//
// Equality is not to be confused with the Is function seen on some of the
// nodes, such as Date.Is. The Is function is used to compare exact raw values
// in nodes.
//
// DeepEqual tests if left and right are recursively equal.
package gedcom

import (
	"fmt"
	"reflect"
)

type Node interface {
	fmt.Stringer
	NodeCopier
	Noder
	GEDCOMLiner
	GEDCOMStringer

	// The node itself.
	Tag() Tag
	Value() string
	Pointer() string

	// Document.
	Document() *Document
	SetDocument(document *Document)

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
func IsNil(node interface{}) bool {
	if node == nil {
		return true
	}

	return reflect.ValueOf(node).IsNil()
}
