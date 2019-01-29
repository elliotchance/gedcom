package gedcom

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

	if left != right {
		if !left.Equals(right) {
			return false
		}
	}

	leftNodes := left.Nodes()
	rightNodes := right.Nodes()
	leftNodesLen := len(leftNodes)
	rightNodesLen := len(rightNodes)

	if leftNodesLen != rightNodesLen {
		return false
	}

	return DeepEqualNodes(leftNodes, rightNodes)
}

// DeepEqualNodes allows two slices of nodes to be compared.
//
// The slices must have the same length (including zero) or the result will
// always be false. If both slices contain zero elements then the result is
// always true.
//
// Every node in the left must DeepEqual a node on the right. The same node
// cannot be used twice in a comparison. The slices are allowed to have
// duplicate nodes (by reference or value) as long as the other side has an
// equal amount of duplicates.
//
// See DeepEqual for semantics on how nodes are compared.
func DeepEqualNodes(left, right Nodes) bool {
	leftLen := len(left)
	rightLen := len(right)

	if leftLen != rightLen {
		return false
	}

	matches := map[int]bool{}
	for _, leftChild := range left {
		foundMatch := false
		for i, rightChild := range right {
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
