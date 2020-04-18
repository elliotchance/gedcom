// Merging
//
// There are several functions available that handle different kinds of merging:
//
// - MergeNodes(left, right Node) Node: returns a new node that merges children
// from both nodes.
//
// - MergeNodeSlices(left, right Nodes, mergeFn MergeFunction) Nodes: merges
// two slices based on the mergeFn. This allows more advanced merging when
// dealing with slices of nodes.
//
// - MergeDocuments(left, right *Document, mergeFn MergeFunction) *Document:
// creates a new document with their respective nodes merged. You can use
// IndividualBySurroundingSimilarityMergeFunction with this to merge
// individuals, rather than just appending them all.
//
// The MergeFunction is a type that can be received in some of the merging
// functions. The closure determines if two nodes should be merged and what the
// result would be. Alternatively it can also describe when two nodes should not
// be merged.
//
// You may certainly create your own MergeFunction, but there are some that are
// already included:
//
// - IndividualBySurroundingSimilarityMergeFunction creates a MergeFunction that
// will merge individuals if their surrounding similarity is at least
// minimumSimilarity.
//
// - EqualityMergeFunction is a MergeFunction that will return a merged node if
// the node are considered equal (with Equals).
package gedcom

import (
	"errors"
	"fmt"
)

// MergeFunction will do one of two things:
//
// 1. If the nodes should be merged, it must return a new new node.
//
// 2. If the nodes should not or could not be merged, then nil is returned.
//
// A MergeFunction can be used with MergeNodeSlices.
//
// The document parameter is used as the destination for merged nodes that need
// to be attached to a document (such as individuals).
type MergeFunction func(left, right Node, document *Document) Node

// MergeNodes returns a new node that merges children from both nodes.
//
// If either of the nodes are nil, or they are not the same tag an error will be
// returned and the result node will be nil.
//
// The node returned and all of the merged children will be created as new
// nodes as to not interfere with the original input.
//
// The document must not be nil and will be used to attach the new nodes to
// (since some nodes require a document, such as individuals). You may supply
// the same document.
func MergeNodes(left, right Node, document *Document) (Node, error) {
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

	r := DeepCopy(left, document)

	for _, child := range right.Nodes() {
		for _, n := range r.Nodes() {
			if n.Equals(child) {
				newNodes := MergeNodeSlices(child.Nodes(), n.Nodes(), document,
					EqualityMergeFunction)
				n.SetNodes(newNodes)
				goto next
			}
		}

		r.AddNode(child)
	next:
	}

	return r, nil
}

// MergeNodeSlices merges two slices based on the mergeFn.
//
// The MergeFunction must not be nil, but may return nil. See MergeFunction for
// usage.
//
// The left and right may contain zero elements or be nil, these mean the same
// thing.
//
// MergeNodeSlices makes some guarantees about the result:
//
// 1. The result slice will contain at least the length of the greatest length
// of left and right. So if len(left) = 3 and len(right) = 5 then the result
// slice will have a minimum of 5 items. If both slices are empty or nil the
// minimum length will also be zero.
//
// 2. The result slice will not contain more elements than the sum of the
// lengths of the left and right. So if len(left) = 3 and len(right) = 5 then
// the largest possible slice returned is 8.
//
// 3. All of the nodes returned will be deep copies of the original nodes so it
// is safe to manipulate the result in any way without affecting the original
// input slices.
//
// 4. Any element from the left or right may only be merged once. That is to say
// that if a merge happens between a left and right node that the result node
// cannot be merged again.
//
// 5. Merges can only happen between a node on the left with a node on the
// right. Even if two nodes in the left could be merged they will not be. The
// same goes for all of the elements in the right slice.
//
// The document must not be nil and will be used to attach the new nodes to
// (since some nodes require a document, such as individuals). You may supply
// the same document.
func MergeNodeSlices(left, right Nodes, document *Document, mergeFn MergeFunction) Nodes {
	newSlice := Nodes{}

	// Be careful to duplicate the right slice. I'm not sure why this is
	// necessary but the unit tests that reuse mergeDocumentsTests will fail if
	// we do not have this.
	newRight := make(Nodes, len(right))
	copy(newRight, right)
	right = newRight

	// We start by adding all of the items on the left.
	for _, node := range left {
		newSlice = append(newSlice, DeepCopy(node, document))
	}

	// Each of the items on the right must be compared with all of the items in
	// the slice, which starts off with all the items from the left but will
	// grow.
	//
	// If the right item does not match anything previously seen then it is
	// appended to the end. Otherwise the left node is removed and the new
	// merged node is place on the end of the new slice.
	//
	// We can guarantee that all the items on the left will be inserted once.
	// However, one obvious problem is that the items on the right may be merged
	// multiple times into a single left element, or even be merged into an
	// element previously appended from the right.
	//
	// To get around this behavior we have to keep track of when a node has
	// already been replaced with a merged one. This gives us the same
	// guarantees for the items on the right.

	alreadyMerged := NodeSet{}

	for len(right) > 0 {
		found := false

		for i := 0; i < len(newSlice); i++ {
			node := newSlice[i]

			// Skip any nodes that were previously marked as merged.
			if alreadyMerged.Has(node) {
				goto next
			}

			for j, node2 := range right {
				merged := mergeFn(node, node2, document)
				if !IsNil(merged) {
					// Remove the current node, and append the new merged one.
					// This will change the order, but the order of nodes
					// doesn't matter in GEDCOM files.
					newSlice = append(newSlice[:i], newSlice[i+1:]...)
					newSlice = append(newSlice, merged)

					// We also have to remove the matching right node, otherwise
					// we may get stuck into an infinite loop.
					right = append(right[:j], right[j+1:]...)

					// Record the fact that this new node has been merged so it
					// will be avoided next time, even in the case of a merge
					// candidate.
					alreadyMerged.Add(merged)

					found = true
					break
				}
			}
		next:
		}

		if !found {
			// See the comment above about why we need to mark the new right
			// node as already merged.
			newNode := DeepCopy(right[0], document)
			alreadyMerged.Add(newNode)

			newSlice = append(newSlice, newNode)
			right = right[1:]
		}
	}

	return newSlice
}

// MergeDocuments creates a new document with their respective nodes merged.
//
// The MergeFunction must not be nil, but may return nil. See MergeFunction for
// usage.
//
// The left and right may be nil. This is treated as an empty document.
//
// The result document will have a deep copy of all nodes. So it's safe to
// manipulate the nodes without affecting the original nodes.
//
// Individuals will not be merged amongst each other, only appended to the final
// document. To merge similar individuals see MergeDocumentsAndIndividuals.
//
// The document (third parameter) must not be nil and will be used to attach the
// new nodes to (since some nodes require a document, such as individuals). You
// may supply the same document.
func MergeDocuments(left, right *Document, document *Document, mergeFn MergeFunction) *Document {
	var leftNodes, rightNodes Nodes

	if left != nil {
		leftNodes = left.Nodes()
	}

	if right != nil {
		rightNodes = right.Nodes()
	}

	newNodes := MergeNodeSlices(leftNodes, rightNodes, document, mergeFn)

	return NewDocumentWithNodes(newNodes)
}

// MergeDocumentsAndIndividuals merges two documents while also merging similar
// individuals. A new document will be returned.
//
// The MergeFunction must not be nil, but may return nil. It will only be used
// for nodes that are not individuals. See MergeFunction for usage.
//
// The options must be provided.
func MergeDocumentsAndIndividuals(left, right *Document, mergeFn MergeFunction, options *IndividualNodesCompareOptions) (*Document, error) {
	// Merge individuals.
	leftIndividuals := individuals(left)
	rightIndividuals := individuals(right)

	document := NewDocument()
	mergedIndividuals, err := leftIndividuals.Merge(rightIndividuals, document, options)
	if err != nil {
		return nil, err
	}

	// Merge non-individuals.
	leftOther := nonIndividuals(left)
	rightOther := nonIndividuals(right)

	mergedOther := MergeNodeSlices(leftOther, rightOther, document, mergeFn)

	allNodes := append(mergedIndividuals.Nodes(), mergedOther...)

	return NewDocumentWithNodes(allNodes), nil
}

// IndividualBySurroundingSimilarityMergeFunction creates a MergeFunction that
// will merge individuals if their surrounding similarity is at least
// minimumSimilarity.
//
// If either of the nodes are not IndividualNode instances then it will always
// return nil so they will not be merged.
//
// The minimumSimilarity should be value between 0.0 and 1.0. The options must
// not be nil, you should use NewSimilarityOptions() for sensible defaults.
//
// The document must not be nil and will determine where the merged nodes will
// end up. This is important for nodes that need to be attached to a document.
func IndividualBySurroundingSimilarityMergeFunction(minimumSimilarity float64, options SimilarityOptions) MergeFunction {
	return func(left, right Node, document *Document) Node {
		leftIndividual, leftOK := left.(*IndividualNode)
		rightIndividual, rightOK := right.(*IndividualNode)

		// Either side is not an individual.
		if !leftOK || !rightOK {
			return nil
		}

		similarity := leftIndividual.SurroundingSimilarity(rightIndividual, options, false)

		if similarity.WeightedSimilarity() > minimumSimilarity {
			// Ignore the error here because left and right must be the same
			// type.
			mergedIndividuals, _ := MergeNodes(left, right, document)

			return mergedIndividuals
		}

		// Do not merge.
		return nil
	}
}

// EqualityMergeFunction is a MergeFunction that will return a merged node if
// the node are considered equal (with Equals). This is probably the most
// generic case when doing general merges.
//
// The document must not be nil and will be used to attach the merged nodes as
// some nodes need a document to be created (such as Individual nodes).
func EqualityMergeFunction(left, right Node, document *Document) Node {
	if left.Equals(right) {
		merged, err := MergeNodes(left, right, document)
		if err != nil {
			return nil
		}

		return merged
	}

	return nil
}
