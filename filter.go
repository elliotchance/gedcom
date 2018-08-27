package gedcom

// FilterFunction is used with the Filter function.
//
// The node (as passed in through the parameter) will be replaced with newNode.
// You should return the same node argument if you do not want the node to be
// changed.
//
// The traverseChildren argument decides if the traversal should continue
// through the children of node. If the traversal continues (traverseChildren is
// true) the FilterFunction will always receive the children of the node, and
// not the children (if any) of newNode.
//
// If the newNode is nil then it will be removed and the children will not be
// traversed, regardless of the traverseChildren value.
type FilterFunction func(node Node) (newNode Node, traverseChildren bool)

// Filter returns a new nest node structure by recursively filtering all
// children based on a callback FilterFunction.
//
// See FilterFunction for the implementation details of fn.
//
// There are several other functions that can be used as filters including;
// WhitelistTagFilter, BlacklistTagFilter and OfficialTagFilter.
func Filter(root Node, fn FilterFunction) Node {
	newRoot, keepTraversing := fn(root)
	if newRoot == nil {
		return nil
	}

	// Shallow copy the node.
	// https://github.com/elliotchance/gedcom/issues/74
	result := NewNode(newRoot.Document(), newRoot.Tag(), newRoot.Value(),
		newRoot.Pointer())

	if keepTraversing {
		for _, child := range root.Nodes() {
			newNode := Filter(child, fn)
			if newNode != nil {
				result.AddNode(newNode)
			}
		}
	}

	return result
}

// WhitelistTagFilter returns any node that is one of the provided tags.
//
// This is the opposite of BlacklistTagFilter.
//
// See the Filter() function.
func WhitelistTagFilter(tags ...Tag) FilterFunction {
	return func(node Node) (Node, bool) {
		for _, tag := range tags {
			if tag.Is(node.Tag()) {
				return node, true
			}
		}

		return nil, false
	}
}

// BlacklistTagFilter returns any node that is not one of the provided tags.
//
// This is the opposite of WhitelistTagFilter.
//
// See the Filter function.
func BlacklistTagFilter(tags ...Tag) FilterFunction {
	return func(node Node) (Node, bool) {
		for _, tag := range tags {
			if tag.Is(node.Tag()) {
				return nil, false
			}
		}

		return node, true
	}
}

// OfficialTagFilter returns any node that is official tag. See Tag.IsOfficial
// for more information.
//
// See the Filter function.
func OfficialTagFilter() FilterFunction {
	return func(node Node) (Node, bool) {
		isOfficial := node.Tag().IsOfficial()

		return NodeCondition(isOfficial, node, nil), isOfficial
	}
}
