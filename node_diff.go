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
// Child nodes belonging to two parent nodes that are considered equal can be
// merged. For example, all of the following examples are considered to be equal
// because they share the same parent value:
//
//   BIRT               |  BIRT               |  BIRT
//     DATE 3 SEP 1943  |    DATE 3 SEP 1943  |    PLAC England
//   BIRT               |    PLAC England     |    DATE 3 SEP 1943
//     PLAC England     |                     |  BIRT
//
// Finally, and perhaps the most important take away is that nodes are absolute
// in value. They can only be absolutely equal or they must be considered
// independent. No matter how similar the events or values they are never placed
// against each other to describe a "difference".
//
// For example, two individuals that have a different birth dates would not
// return a single DiffNode with the left and right side different values,
// instead two NodeDiffs would be returned that describe each side having an
// event that is not present in the other. Only when the nodes are absolutely
// equal are they said to be the same.
type NodeDiff struct {
	// Left or Right may be nil, but never both. Only in the case of the root
	// node may the Left and Right represent nodes of different values.
	// Otherwise nodes with different values will be added as separate children.
	//
	// Since nodes may be compared from different documents it's important to
	// retain the left and right nodes when they are both equal so they stay
	// connected to their original document.
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
// even if they different values. This is the only case where this is possible.
// It was decided to do it this way to allow comparing of nodes that often have
// a different root node value (like individuals with different pointers).
//
// If you need to be sure the root node are the same or different you can check
// before-hand, or afterwards with (this is also nil safe):
//
//   GedcomLine(-1, d.Left) == GedcomLine(-1, d.Right)
//
// The process of comparison is easiest to explain by following an example. Keep
// in mind that if any of the rules seem confusing you should refer to the
// documentation of NodeDiff.
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
// The first step is to flatten the nodes. Each node is represented as an slice
// containing all of its parent elements:
//
//   -- Left
//   ["INDI @P3@"]
//   ["INDI @P3@", "NAME John /Smith/"]
//   ["INDI @P3@", "BIRT"]
//   ["INDI @P3@", "BIRT", "DATE 3 SEP 1943"]
//   ["INDI @P3@", "DEAT"]
//   ["INDI @P3@", "DEAT", "PLAC England"]
//   ["INDI @P3@", "BIRT"]
//   ["INDI @P3@", "BIRT", "DATE Abt. Oct 1943"]
//
//   -- Right
//   ["INDI @P4@"]
//   ["INDI @P4@", "NAME J. /Smith/"]
//   ["INDI @P4@", "BIRT"]
//   ["INDI @P4@", "BIRT", "DATE Abt. Sep 1943"]
//   ["INDI @P4@", "BIRT"]
//   ["INDI @P4@", "BIRT", "DATE 3 SEP 1943"]
//   ["INDI @P4@", "DEAT"]
//   ["INDI @P4@", "DEAT", "DATE Aft. 2001"]
//   ["INDI @P4@", "DEAT", "PLAC Surry, England"]
//
// Now we begin unflattening the data back into the nested structure. The "L"
// and "R" represent which side is not nil in the DiffNode.
//
// Beginning with the Left items:
//
//   L  0 INDI @P3@
//   L  1 NAME John /Smith/
//   L  1 BIRT
//   L  2 DATE Abt. Oct 1943
//   L  2 DATE 3 SEP 1943
//   L  1 DEAT
//   L  2 PLAC England
//
// Then the Right items and applied on top. Equal nodes will be converted from
// "L" -> "LR" (assign right):
//
//   LR 0 INDI @P3@            | = Assign right (notice the @P4@ is not shown)
//   L  1 NAME John /Smith/    |
//   LR 1 BIRT                 | = Assign right
//   L  2 DATE Abt. Oct 1943   |
//   LR 2 DATE 3 SEP 1943      | = Assign right
//    R 2 DATE Abt. Sep 1943   | + Added
//   LR 1 DEAT                 | = Assign right
//   L  2 PLAC England         |
//    R 2 DATE Aft. 2001       | + Added
//    R 2 PLAC Surry, England  | + Added
//    R 1 NAME J. /Smith/      | + Added
//
// The output format is the same as NodeDiff.String().
func CompareNodes(left, right Node) *NodeDiff {
	// Flatten the nodes. We use the left prefix when flattening the right nodes
	// as to not interfere with the a root node on the right that may be
	// different. At the end we will assign the right root node.
	//
	// The extra conditional check is to make sure that a nil left does not
	// break the right only traversing.
	prefix := []string{GedcomLine(-1, NodeCondition(IsNil(left), right, left))}
	flatLeft := flattenNode(left, prefix)
	flatRight := flattenNode(right, prefix)

	// Unflatten the nodes back to the diff.
	result := &NodeDiff{}
	result.unflatten(flatLeft, false)
	result.unflatten(flatRight, true)

	// Fix up the right root node.
	result.Right = right

	return result
}

func (nd *NodeDiff) unflatten(flatNodes [][]string, assignRight bool) {
	for _, item := range flatNodes {
		nd.unflattenSingle(item, assignRight)
	}
}

func (nd *NodeDiff) unflattenSingle(flatNode []string, assignRight bool) {
	i := nd
	for _, line := range flatNode {
		// We can ignore the error here because we encoded the lines.
		parsedLine, _, _ := parseLine(nil, "0 "+line)

		switch {
		case nd.Left == nil && nd.Right == nil:
			// This is the first (root) node.

			nd.Left = NodeCondition(!assignRight, parsedLine, nd.Left)
			nd.Right = NodeCondition(assignRight, parsedLine, nd.Right)
			i = nd

		case GedcomLine(-1, i.Left) == line || GedcomLine(-1, i.Right) == line:
			// We have found a match with an existing line, traverse down. This
			// is not the same situation below were each of the children are
			// checked.

			if assignRight {
				i.Right = parsedLine
			}

		default:
			// We will be adding a new node. Search the children to see if there
			// is a match, otherwise add it to the end.

			child := &NodeDiff{
				Left:  NodeCondition(!assignRight, parsedLine, nil),
				Right: NodeCondition(assignRight, parsedLine, nil),
			}

			found := false
			for _, n := range i.Children {
				if GedcomLine(-1, n.Left) == line || GedcomLine(-1, n.Right) == line {
					if assignRight {
						n.Right = parsedLine
					}
					i = n
					found = true
					break
				}
			}

			if !found {
				i.Children = append(i.Children, child)
				i = child
			}
		}
	}
}

func flattenNode(node Node, prefix []string) [][]string {
	if IsNil(node) {
		return [][]string{}
	}

	children := node.Nodes()
	if len(children) == 0 {
		return [][]string{prefix}
	}

	r := [][]string{prefix}
	for _, child := range children {
		line := GedcomLine(-1, child)
		r = append(r, flattenNode(child, append(prefix, line))...)
	}

	return r
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
