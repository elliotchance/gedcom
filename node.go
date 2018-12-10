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

// DeepEqual tests if left and right are recursively equal.
//
// If either left or right is nil (including both) then false is always
// returned.
//
// If left does not equal right (see Node.Equals) or both sides do not contain
// exactly the same amount of child nodes then false is returned.
//
// The GEDCOM standard allows nodes to appear in any order. So the children are
// compared in this way as well. For example the following root nodes are equal:
//
//   0 INDI @P1@        |  0 INDI @P1@
//   1 BIRT             |  1 BIRT
//   2 DATE 3 SEP 1943  |  2 PLAC England
//   2 PLAC England     |  2 DATE 3 SEP 1943
//
// DeepEqual heavily depends on the logic of the Equals method for each kind of
// node. Equals may or may not take into consideration child nodes to determine
// if the parent itself is equal. You should see the specific documentation for
// Equals on each node type.
//
// If Equals is not implemented it will fall back to SimpleNode.Equals.
//
// If an equal node appears multiple times on either side it will also have to
// appear the same number of times on the opposite side for the DeepEqual to be
// true.
func DeepEqual(left, right Node) bool {
	if IsNil(left) {
		return false
	}

	if IsNil(right) {
		return false
	}

	if !left.Equals(right) {
		return false
	}

	leftNodes := left.Nodes()
	rightNodes := right.Nodes()
	leftNodesLen := len(leftNodes)
	rightNodesLen := len(rightNodes)

	if leftNodesLen != rightNodesLen {
		return false
	}

	matches := map[int]bool{}
	for _, leftChild := range leftNodes {
		foundMatch := false
		for i, rightChild := range rightNodes {
			if !matches[i] && DeepEqual(leftChild, rightChild) {
				matches[i] = true
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			return false
		}
	}

	return true
}
