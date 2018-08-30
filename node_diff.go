package gedcom

import (
	"fmt"
	"strings"
)

// NodeDiff is used to describe the difference when two nodes are compared.
//
// It is important to understand the semantics of what a "difference" (and
// therefore "equality") means for GEDCOM data as this heavily factors into
// influencing the way the algorithms and returned values represent.
//
// GEDCOM files are quite unbounded when it comes to events and facts. For
// example, it's common to have multiple birth events (BIRT tag) for the same
// individual. This is not necessarily a bug in the data but rather a way to
// describe two possible known birth dates or locations.
//
// The order of nodes in the GEDCOM file is also insignificant. That is to say
// that a birth event that appears before another birth event is no more
// important than any other tag, including other birth events.
//
// Child nodes belonging to two parent nodes that are considered equal (by using
// Equals) can be merged. For example, all of the following examples are
// considered to be equal because the BirthNode.Equals regards all BirthNodes as
// equal (see specific documentation for a complete explanation):
//
//   BIRT               |  BIRT               |  BIRT
//     DATE 3 SEP 1943  |    DATE 3 SEP 1943  |    PLAC England
//   BIRT               |    PLAC England     |    DATE 3 SEP 1943
//     PLAC England     |                     |  BIRT
//
// However, the semantics of Equals is quite different for other types of nodes.
// For example ResidenceNodes are considered equal only if they have the same
// date, as it wouldn't make sense (or just be plain wrong) to merge children
// from separate Residence events.
type NodeDiff struct {
	// Left or Right may be nil, but never both.
	//
	// Since nodes may be compared from different documents and have different
	// raw values it's important to retain the left and right nodes. Make sure
	// when displaying or traversing your data you are showing the correct side.
	//
	// The Left and Right retain their original children as well so you an still
	// perform all the same operations on the nodes inside a NodeDiff.
	Left, Right Node

	// Children represents each of the compared child nodes from both sides. See
	// CompareNodes for a full explanation.
	Children []*NodeDiff
}

// CompareNodes returns the recursive comparison of two root nodes. To properly
// understand the motivations and expected result you should first understand
// NodeDiff before you continue reading.
//
// The returned NodeDiff will have root node assigned to the Left and Right,
// even if they are non-equal values. This is the only case where this is
// possible. It was decided to do it this way to allow comparing of nodes that
// often have a different root node value (like individuals with different
// pointers), and also keep the output from CompareNodes consistent.
//
// If you need to be sure the root node are the equal after the comparison, you
// can use (this is also nil safe):
//
//   d.Left.Equals(d.Right)
//
// The algorithm to perform the diff is actually very simple:
//
// 1. It creates an empty NodeDiff instance.
//
// 2. It traverses down the left creating all the respective child nodes as it
// goes. Before it adds a child node at any level it will always check
// previously created nodes at the same level for equality. If it finds a match
// it will redirect the traversal through this parent rather than creating a new
// child.
//
// 3. It traverses the right side with the same rules. The only real difference
// is that it will assign the node to the right side on a match/new child
// instead of the left.
//
// Here are two individuals that have slightly different data:
//
//   0 INDI @P3@           |  0 INDI @P4@
//   1 NAME John /Smith/   |  1 NAME J. /Smith/
//   1 BIRT                |  1 BIRT
//   2 DATE 3 SEP 1943     |  2 DATE Abt. Sep 1943
//   1 DEAT                |  1 BIRT
//   2 PLAC England        |  2 DATE 3 SEP 1943
//   1 BIRT                |  1 DEAT
//   2 DATE Abt. Oct 1943  |  2 DATE Aft. 2001
//                         |  2 PLAC Surry, England
//
// In this case both of the root nodes are different (because of the different
// pointer values). The returned left and right will have the respective root
// nodes.
//
// Here is the output, rendered with NodeDiff.String():
//
//   LR 0 INDI @P3@
//   L  1 NAME John /Smith/
//   LR 1 BIRT
//   L  2 DATE Abt. Oct 1943
//   LR 2 DATE 3 SEP 1943
//    R 2 DATE Abt. Sep 1943
//   LR 1 DEAT
//   L  2 PLAC England
//    R 2 DATE Aft. 2001
//    R 2 PLAC Surry, England
//    R 1 NAME J. /Smith/
//
func CompareNodes(left, right Node) *NodeDiff {
	result := &NodeDiff{}

	result.traverse(left, true)
	result.traverse(right, false)

	return result
}

func (nd *NodeDiff) traverse(n Node, isLeft bool) {
	if n == nil {
		return
	}

	if isLeft && nd.Left == nil {
		nd.Left = n
	}

	if !isLeft && nd.Right == nil {
		nd.Right = n
	}

	for _, child := range n.Nodes() {
		found := false
		for _, diffChild := range nd.Children {
			if diffChild.Left != nil && diffChild.Left.Equals(child) {
				diffChild.traverse(child, isLeft)
				found = true
				break
			}

			if diffChild.Right != nil && diffChild.Right.Equals(child) {
				diffChild.traverse(child, isLeft)
				found = true
				break
			}
		}

		if !found {
			newNd := &NodeDiff{}
			newNd.traverse(child, isLeft)

			nd.Children = append(nd.Children, newNd)
		}
	}
}

func (nd *NodeDiff) lrLine(indent int) string {
	left := GedcomLine(indent, nd.Left)
	right := GedcomLine(indent, nd.Right)

	if IsNil(nd.Left) {
		return fmt.Sprintf(" R %s", right)
	}

	if IsNil(nd.Right) {
		return fmt.Sprintf("L  %s", left)
	}

	// Only the root can have different values for the left and right node.
	// We want to display this so we show it as two different LR root nodes.
	if left != right {
		return fmt.Sprintf("LR %s\nLR %s", left, right)
	}

	return fmt.Sprintf("LR %s", GedcomLine(indent, nd.Left))
}

func (nd *NodeDiff) string(indent int) string {
	s := nd.lrLine(indent)

	for _, child := range nd.Children {
		s += "\n" + child.string(indent+1)
	}

	return s
}

// String returns a readable comparison of nodes, like:
//
//   LR 0 INDI @P3@
//   L  1 NAME John /Smith/
//   LR 1 BIRT
//   L  2 DATE Abt. Oct 1943
//   LR 2 DATE 3 SEP 1943
//    R 2 DATE Abt. Sep 1943
//   LR 1 DEAT
//   L  2 PLAC England
//    R 2 DATE Aft. 2001
//    R 2 PLAC Surry, England
//    R 1 NAME J. /Smith/
//
// The L/R/LR represent which side has the node, followed by the GEDCOM indent
// and node line.
//
// There is a special case if both root nodes are different. They will be
// displayed as two separate lines even though they both belong to the same
// NodeDiff:
//
//   LR 0 INDI @P3@
//   LR 0 INDI @P4@
//
// You should not rely on this format to be machine readable as it may change in
// the future.
func (nd *NodeDiff) String() string {
	return strings.TrimSpace(nd.string(0))
}
