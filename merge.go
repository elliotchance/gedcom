package gedcom

import (
	"fmt"
	"errors"
)

// MergeNodes returns a new node that merges children from both nodes.
//
// If either of the nodes are nil, or they are not the same tag an error will be
// returned and the result node will be nil.
//
// The node returned and all of the merged children will be created as new
// nodes as to not interfere with the original input.
func MergeNodes(left, right Node) (Node, error) {
	if IsNil(left) {
		return nil, errors.New("left is nil")
	}

	if IsNil(right) {
		return nil, errors.New("right is nil")
	}

	leftTag := left.Tag()
	rightTag := right.Tag()

	// We can only proceed if the nodes can be merged.
	if !leftTag.Is(rightTag) {
		return nil, fmt.Errorf("cannot merge %s and %s nodes",
			leftTag.Tag(), rightTag.Tag())
	}

	r := DeepCopy(left)

	for _, child := range right.Nodes() {
		for _, n := range r.Nodes() {
			if n.Equals(child) {
				for _, child2 := range child.Nodes() {
					n.AddNode(child2)
				}
				goto next
			}
		}

		r.AddNode(child)
	next:
	}

	return r, nil
}
